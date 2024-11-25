package config

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var FirestoreClient *firestore.Client

func InitFirebase() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	credentials := os.Getenv("FIREBASE_CREDENTIALS")
	if credentials == "" {
		log.Fatalf("FIREBASE_CREDENTIALS environment variable not set in .env")
	}

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", credentials)

	client, err := firestore.NewClient(context.Background(), os.Getenv("PROJECT_ID"))
	if err != nil {
		log.Fatalf("Error creating Firestore client: %v", err)
	}
	FirestoreClient = client

}
