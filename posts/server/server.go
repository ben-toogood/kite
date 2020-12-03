package server

import (
	"cloud.google.com/go/storage"
	"gorm.io/gorm"
)

// Posts implements the posts handler interface
type Posts struct {
	DB     *gorm.DB
	Bucket *storage.BucketHandle
}
