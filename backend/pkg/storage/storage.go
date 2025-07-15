package storage

import (
	configs "appointment-platform-backend-backend/cmd/config"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient(host string, user string, password string) (*minio.Client, error) {
	secure := true

	if configs.ApplicationCfg.Env == "local" {
		secure = false
	}

	return minio.New(host, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: secure,
	})
}

func CreateBucketIfNotExists(ctx context.Context, client *minio.Client, bucketName string) error {
	exists, err := client.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	err = client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
	if err != nil {
		return err
	}

	return nil
}
