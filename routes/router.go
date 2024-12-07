package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/handlers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/moods")

	api.Post("/", handlers.CreateMoodEntry)
	api.Get("/", handlers.GetAllMood)
}

func NewsRoutes(app *fiber.App) {
	api := app.Group("/api/recommendation")
	//	api.Get("/get", handlers.GetRecommendation)
	api.Get("/", handlers.GetAllNews)
}
