package gcs

import (
	"context"
	"io"

	"cloud.google.com/go/storage"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

func NewGCS(pathToKey, bucketName string) (*gcs, error) {
	client, err := storage.NewClient(context.Background(), option.WithCredentialsFile(pathToKey))
	if err != nil {
		return nil, err
	}

	bucket := client.Bucket(bucketName)

	return &gcs{
		client: client,
		bucket: bucket,
	}, nil
}

type gcs struct {
	client *storage.Client
	bucket *storage.BucketHandle
}

func (g *gcs) Upload(ctx context.Context, name string, file io.Reader) error {
	object := g.bucket.Object(name)

	writer := object.NewWriter(ctx)

	_, err := io.Copy(writer, file)
	if err != nil {
		logrus.WithError(err).Errorf("failed to write %s to gcs", name)
		return err
	}

	if err := writer.Close(); err != nil {
		logrus.WithError(err).Errorf("failed to flush %s to gcs", name)
		return err
	}

	return nil
}
