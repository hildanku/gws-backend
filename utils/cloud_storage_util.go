package utils

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

func UploadGCS(file multipart.File, userID string) (string, error) {
	log.Println("at UPLOADGCS", userID)

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" {
		return "", fmt.Errorf("BUCKET_NAME environment variable not set")
	}

	client, err := storage.NewClient(context.Background())
	if err != nil {
		return "", fmt.Errorf("failed to create storage client: %v", err)
	}
	defer client.Close()

	timestamp := time.Now().Format("20060102150405") // Format: YYYYMMDDHHMMSS
	filename := fmt.Sprintf("%s/vn-%s.mp3", userID, timestamp)

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

	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, filename)
	return fileURL, nil
}
