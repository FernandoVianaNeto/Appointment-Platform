package minio

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

import (
	storage_adapter "appointment-platform-backend-backend/internal/domain/adapters/storage"
	"context"
	"io"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type MinIoAdapter struct {
	Client *minio.Client
}

func NewMinIoAdapter(ctx context.Context, client *minio.Client) storage_adapter.StorageAdapterInterface {
	return &MinIoAdapter{
		Client: client,
	}
}

func (m *MinIoAdapter) UploadMedia(ctx context.Context, bucketName string, objectName string, reader io.Reader, objectSize int64, contentType string) error {
	_, err := m.Client.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		objectSize,
		minio.PutObjectOptions{
			ContentType: contentType,
		},
	)
	return err
}

func (m *MinIoAdapter) GeneratePresignedURL(
	ctx context.Context,
	bucketName string,
	objectName string,
	expiration time.Duration,
) (string, error) {
	reqParams := url.Values{}
	presignedURL, err := m.Client.PresignedGetObject(
		ctx,
		bucketName,
		objectName,
		expiration,
		reqParams,
	)

	if err != nil {
		return "", err
	}

	return presignedURL.String(), nil
}
