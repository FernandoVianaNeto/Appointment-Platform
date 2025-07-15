package storage_adapter

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	"context"
	"io"
	"time"
)

type StorageAdapterInterface interface {
	UploadMedia(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, contentType string) error
	GeneratePresignedURL(ctx context.Context, bucketName string, objectName string, expiration time.Duration) (string, error)
}
