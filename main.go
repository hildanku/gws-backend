package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/config"
	"github.com/gws-app/gws-backend/routes"
	"github.com/joho/godotenv"
)

func main() {
	// Init .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Init Firestore
	config.InitFirebase()

	// Init Fiber
	app := fiber.New()

	// Init Route
	routes.Routes(app)

	// Init news
	routes.NewsRoutes(app)

	port := os.Getenv("PORT")
	log.Fatal(app.Listen(":" + port))
}
