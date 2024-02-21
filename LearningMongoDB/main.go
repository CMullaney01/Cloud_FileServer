package example

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

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
	var fileData File
	if err := json.NewDecoder(r.Body).Decode(&fileData); err != nil {
		http.Error(w, "Failed to decode file data", http.StatusBadRequest)
		return
	}

	// Insert the file data into the collection
	_, err := coll.InsertOne(context.TODO(), fileData)
	if err != nil {
		http.Error(w, "Failed to insert file into database", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fileData)
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
