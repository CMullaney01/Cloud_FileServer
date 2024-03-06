package entities

import (
	"time"
)

type File struct {
	UserID      string    `bson:"userId"`
	FileName    string    `bson:"fileName"`
	S3Bucket    string    `bson:"s3Bucket"`
	S3ObjectKey string    `bson:"s3ObjectKey"`
	CreatedAt   time.Time `bson:"createdAt"`
	IsPublic    bool      `bson:"isPublic"`
	Size        int64     `bson:"size"`
	ContentType string    `bson:"contentType"`
}
