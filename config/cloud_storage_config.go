package config

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/joho/godotenv"
)

var CloudStorageClient *storage.Client

func InitCloudStorage() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if credentials == "" {
		log.Fatalf("GOOGLE_APPLICATION_CREDENTIALS environment variable not set in .env")
	}

	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	CloudStorageClient = client
	log.Println("Cloud Storage Success")
}
