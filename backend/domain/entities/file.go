package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      string             `bson:"userId"`
	FileName    string             `bson:"fileName"`
	S3Bucket    string             `bson:"s3Bucket"`
	S3ObjectKey string             `bson:"s3ObjectKey"`
	CreatedAt   time.Time          `bson:"createdAt"`
	IsPublic    bool               `bson:"isPublic"`
	Size        int64              `bson:"size"`
	ContentType string             `bson:"contentType"`
}
