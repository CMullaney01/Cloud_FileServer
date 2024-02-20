package entities

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	Id          uuid.UUID // Unique identifier for the file
	CreatedAt   time.Time // Timestamp when the file was uploaded
	Name        string    // Name of the file
	ContentType string    // MIME type of the file (e.g., "image/jpeg", "application/pdf")
	Size        int64     // Size of the file in bytes
	Path        string    // Path to the stored file (assuming you're storing it locally)
	// Metadata    map[string]interface{} // Additional metadata about the file (e.g., key-value pairs)
	// Add more fields as needed
}

// the meta data field could be useful in the future for things like this not used for now
// Author: "John Doe"
// Description: "This is an image of a mountain"
// Tags: ["nature", "landscape", "mountain"]
// CreatedBy: "User123"
