package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type File struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId"`
	FileName    string             `bson:"fileName"`
	S3Bucket    string             `bson:"s3Bucket"`
	S3ObjectKey string             `bson:"s3ObjectKey"`
	CreatedAt   time.Time          `bson:"createdAt"`
	IsPublic    bool               `bson:"isPublic"`
}

func GeneratePresignedURL(bucketName, objectKey string, expiration time.Duration) (string, error) {
	// Retrieve AWS credentials from environment variables
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	awsRegion := os.Getenv("AWS_REGION")

	// Create AWS credentials
	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecretKey, "")

	// Create a new AWS session with the provided credentials and region
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awsRegion),
		Credentials: creds,
	})
	if err != nil {
		return "", err
	}

	// Create a new S3 service client
	svc := s3.New(sess)

	// Generate the presigned URL
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	presignedURL, err := req.Presign(expiration) // Use the expiration directly
	if err != nil {
		return "", err
	}

	return presignedURL, nil
}

// create a type which we can utilise to implement different abilities
type BackendMongoClient struct {
	client *mongo.Client
}

func NewMongoClient(c *mongo.Client) *BackendMongoClient {
	return &BackendMongoClient{
		client: c,
	}
}

// given a MongoClient we want to get a specific
// this is a http handler
func (bmc *BackendMongoClient) handleUserFileList(w http.ResponseWriter, r *http.Request) {
	// Parse the user ID from the request or get it from the user's session
	userID := "user123" // Example user ID, replace with actual user ID

	// Access the "files" collection in the "user-data" database
	coll := bmc.client.Database("user-data").Collection("files")

	// Define the filter to find files for the specific user or labeled as public
	filter := bson.M{
		"$or": []bson.M{
			{"userId": userID}, // Filter for files specific to the user
			{"isPublic": true}, // Filter for public files
		},
	}

	// Query files for the specific user or labeled as public
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	// Define a slice to store files for the specific user or labeled as public
	var userFiles []File
	if err := cursor.All(context.Background(), &userFiles); err != nil {
		log.Fatal(err)
	}

	// Set HTTP response headers
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// Encode the userFiles slice to JSON and send the response
	if err := json.NewEncoder(w).Encode(userFiles); err != nil {
		log.Fatal(err)
	}
}
func (bmc *BackendMongoClient) handleFileUpload(w http.ResponseWriter, r *http.Request) {
	// Access the "files" collection in the "user-data" database
	coll := bmc.client.Database("user-data").Collection("files")

	// Decode the file data from the request body
	var fileData struct {
		FileName string `json:"fileName"`
		UserID   string `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&fileData); err != nil {
		http.Error(w, "Failed to decode file data", http.StatusBadRequest)
		return
	}

	// Generate other necessary fields
	fileID := primitive.NewObjectID()
	createdAt := time.Now()
	isPublic := false // You can set this as required

	// Construct S3 object key with userID and filename
	objectKey := fileData.UserID + "/" + fileData.FileName

	// Get the S3 bucket name from environment variable
	s3Bucket := os.Getenv("AWS_S3_BUCKET")
	if s3Bucket == "" {
		http.Error(w, "AWS_S3_BUCKET environment variable is not set", http.StatusInternalServerError)
		return
	}

	// Create the File object
	newFile := File{
		ID:          fileID,
		UserID:      fileData.UserID,
		FileName:    fileData.FileName,
		S3Bucket:    s3Bucket,
		S3ObjectKey: objectKey,
		CreatedAt:   createdAt,
		IsPublic:    isPublic,
	}

	// Insert the file data into the collection
	_, err := coll.InsertOne(context.TODO(), newFile)
	if err != nil {
		http.Error(w, "Failed to insert file into database", http.StatusInternalServerError)
		return
	}

	// Generate presigned URL for the uploaded file
	presignedURL, err := GeneratePresignedURL(s3Bucket, objectKey, 20*time.Second) // Adjust expiration time as needed
	if err != nil {
		http.Error(w, "Failed to generate presigned URL", http.StatusInternalServerError)
		return
	}

	// Include presigned URL in the response
	response := struct {
		File         File   `json:"file"`
		PresignedURL string `json:"presignedURL"`
	}{
		File:         newFile,
		PresignedURL: presignedURL,
	}

	// Respond with success and the presigned URL
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (bmc *BackendMongoClient) handleFileDownload(w http.ResponseWriter, r *http.Request) {
	// Access the "files" collection in the "user-data" database
	coll := bmc.client.Database("user-data").Collection("files")

	// Parse the file ID from the request or get it from the request parameters
	fileID := r.URL.Query().Get("fileID")
	if fileID == "" {
		http.Error(w, "File ID is required", http.StatusBadRequest)
		return
	}

	// Define a filter to find the file by its ID
	filter := bson.M{"_id": fileID}

	// Find the file in the collection
	var fileData File
	if err := coll.FindOne(context.Background(), filter).Decode(&fileData); err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Respond with the file data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileData)
}

func main() {
	err := godotenv.Load(".env.local")
	if err != nil {
		panic("Error loading .env.local file")
	}
	// Connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Create a new instance of BackendMongoClient
	bmc := NewMongoClient(client)

	// Register HTTP handlers for each endpoint
	http.HandleFunc("/listfiles", bmc.handleUserFileList) // Handle user file list
	http.HandleFunc("/upload", bmc.handleFileUpload)      // Handle file upload
	http.HandleFunc("/download", bmc.handleFileDownload)  // Handle file download

	// Start the HTTP server
	log.Println("Server started on :3000")
	http.ListenAndServe(":3000", nil)
}
