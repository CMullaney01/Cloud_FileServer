package mongomgmt

import (
	"context"
	"fmt"
	"log"
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

func (m *FileManager) connectToMongoDB(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(m.monogURI))
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
	client, err := m.connectToMongoDB(ctx)
	if err != nil {
		return "", errors.Wrap(err, "failed to connect to MongoDB")
	}
	defer client.Disconnect(ctx)

	coll := client.Database("user-data").Collection("files")

	filter := bson.M{"S3ObjectKey": file.S3ObjectKey}

	// Perform the replacement operation
	// stops there being multiple of the same file
	_, err = coll.ReplaceOne(ctx, filter, file, options.Replace().SetUpsert(true))
	if err != nil {
		return "", errors.Wrap(err, "failed to insert or update file")
	}

	// Generate presigned URL for the uploaded file
	presignedURL, err := m.GeneratePresignedURL(file.S3Bucket, file.S3ObjectKey, "PUT", 20*time.Second) // Adjust expiration time as needed
	if err != nil {
		return "", errors.Wrap(err, "failed to generate presigned URL")
	}

	// Return success response
	return presignedURL, nil
}

func (m *FileManager) List(ctx context.Context, userID string) ([]entities.File, error) {
	client, err := m.connectToMongoDB(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to MongoDB")
	}
	defer client.Disconnect(ctx)

	coll := client.Database("user-data").Collection("files")

	// Define filter to find files associated with the given user ID
	filter := bson.M{
		"$or": []bson.M{
			{"userId": userID}, // Filter for files specific to the user
			{"isPublic": true}, // Filter for public files
		},
	}

	// Query files for the specific user or labeled as public
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx) // Close cursor when function exits

	// Define a slice to store files for the specific user or labeled as public
	var userFiles []entities.File
	if err := cursor.All(ctx, &userFiles); err != nil {
		return nil, errors.Wrap(err, "failed to decode files")
	}

	return userFiles, nil
}

func (m *FileManager) Download(ctx context.Context, fileName string, userId string) (string, error) {
	client, err := m.connectToMongoDB(ctx)
	if err != nil {
		return "", err
	}
	defer client.Disconnect(ctx)

	coll := client.Database("user-data").Collection("files")

	// Define filter to find the file based on the file name and user ID
	filter := bson.M{"userId": userId, "fileName": fileName}

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
