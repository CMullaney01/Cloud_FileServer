package mongomgmt

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/domain/entities"
)

type FileManager struct {
	monogURI string
}

func NewFileManager() *FileManager {
	return &FileManager{
		monogURI: viper.GetString("Mongo_URI"),
	}
}

func (m *FileManager) connectToMongoDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.monogURI))
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to mongodb")
	}
	return client, nil
}

func (m *FileManager) GeneratePresignedURL(bucketName, objectKey string, method string, expiration time.Duration) (string, error) {
	// Retrieve AWS credentials from environment variables
	awsAccessKey := viper.GetString("AWS_ACCESS_KEY_ID")
	awsSecretKey := viper.GetString("AWS_SECRET_ACCESS_KEY")
	awsRegion := viper.GetString("AWS_REGION")

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
	req.HTTPRequest.Method = method
	presignedURL, err := req.Presign(expiration) // Use the expiration directly
	if err != nil {
		return "", err
	}

	return presignedURL, nil
}

func (m *FileManager) Upload(ctx context.Context, file *entities.File) (string, error) {
	client, err := m.connectToMongoDB()
	if err != nil {
		return "", errors.Wrap(err, "failed to connect to MongoDB")
	}

	coll := client.Database("user-data").Collection("files")

	// Use the provided context
	_, err = coll.InsertOne(ctx, file)
	if err != nil {
		return "", errors.Wrap(err, "failed to insert file into database")
	}

	// Generate presigned URL for the uploaded file
	presignedURL, err := m.GeneratePresignedURL(file.S3Bucket, file.S3ObjectKey, "POST", 20*time.Second) // Adjust expiration time as needed
	if err != nil {
		return "", errors.Wrap(err, "failed to generate presigned URL")
	}

	// Return success response
	return presignedURL, nil
}

func (m *FileManager) List(ctx context.Context, userID string) (map[string]string, error) {
	client, err := m.connectToMongoDB()
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB")
	}

	coll := client.Database("user-data").Collection("files")

	// Define filter to find files associated with the given user ID
	filter := bson.M{
		"$or": []bson.M{
			{"userId": userID}, // Filter for files specific to the user
			{"isPublic": true}, // Filter for public files
		},
	}

	// Find files based on the filter
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to list files")
	}
	defer cursor.Close(ctx)

	// Map to store file name and corresponding signed URLs
	fileURLs := make(map[string]string)

	// Iterate over the cursor and generate signed URLs
	for cursor.Next(ctx) {
		var file entities.File
		if err := cursor.Decode(&file); err != nil {
			return nil, errors.Wrap(err, "failed to decode file")
		}
		presignedURL, err := m.GeneratePresignedURL(file.S3Bucket, file.S3ObjectKey, "GET", 20*time.Second)
		if err != nil {
			return nil, errors.Wrap(err, "failed to generate presigned URL")
		}
		fileURLs[file.FileName] = presignedURL
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return fileURLs, nil
}

func (m *FileManager) Download(ctx context.Context, fileName string, userId string) (string, error) {
	client, err := m.connectToMongoDB()
	if err != nil {
		return "", err
	}

	coll := client.Database("user-data").Collection("files")

	// Define filter to find the file based on the file name and user ID
	filter := bson.M{"userId": userId, "name": fileName}

	// Find the file based on the filter
	var file entities.File
	if err := coll.FindOne(ctx, filter).Decode(&file); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", fmt.Errorf("file not found for userId: %s and fileName: %s", userId, fileName)
		}
		return "", err
	}

	// Assuming you have a function to generate a pre-signed URL for the file
	preSignedURL, err := m.GeneratePresignedURL(file.S3Bucket, file.S3ObjectKey, "GET", 20*time.Second)
	if err != nil {
		return "", err
	}

	return preSignedURL, nil
}