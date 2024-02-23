package mongomgmt

import (
	"context"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"backend/domain/entities"
)

type FileManager struct {
	monogURI string
}

func NewFileManager() *FileManager {
	return &FileManager{
		monogURI: viper.GetString("Mongo.URI"),
	}
}

func (m *FileManager) connectToMongoDB() (*mongo.Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(m.monogURI))
	if err != nil {
		return nil, errors.Wrap(err, "could not connect to mongodb")
	}
	return client, nil
}

func (m *FileManager) GeneratePresignedURL(bucketName, objectKey string, expiration time.Duration) (string, error) {
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
	presignedURL, err := m.GeneratePresignedURL(file.S3Bucket, file.S3ObjectKey, 20*time.Second) // Adjust expiration time as needed
	if err != nil {
		return "", errors.Wrap(err, "failed to generate presigned URL")
	}

	// Return success response
	return presignedURL, nil
}
