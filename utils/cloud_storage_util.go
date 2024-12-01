package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"cloud.google.com/go/storage"
)

func UploadGCS(file multipart.File, filename string) (string, error) {

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("BUCKET_NAME environment variable not set")
	}

	client, err := storage.NewClient(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}
	defer client.Close()

	bucket := client.Bucket(bucketName)
	object := bucket.Object(filename)
	writer := object.NewWriter(context.Background())

	// Copy file
	if _, err := io.Copy(writer, file); err != nil {
		return "", fmt.Errorf("failed to upload file to GCS: %v", err)
	}

	// close resource
	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %v", err)
	}

	// Generate URL
	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", "gws-storage", filename)
	return fileURL, nil
}
