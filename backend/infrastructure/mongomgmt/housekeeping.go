package mongomgmt

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
)

func (m *FileManager) StartPendingUploadsCleanup(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := m.CleanupPendingUploads(ctx); err != nil {
				log.Println("Error cleaning up pending uploads:", err)
			}
		}
	}
}

func (m *FileManager) CleanupPendingUploads(ctx context.Context) error {
	client, err := m.connectToMongoDB(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to connect to MongoDB")
	}
	defer client.Disconnect(ctx)

	pendingColl := client.Database("user-data").Collection("pending_uploads")
	_, err = pendingColl.DeleteMany(ctx, bson.M{})
	if err != nil {
		return errors.Wrap(err, "failed to delete pending uploads")
	}

	return nil
}
