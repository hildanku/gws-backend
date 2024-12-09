package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/handlers"
)

func Routes(app *fiber.App) {
	api := app.Group("/api/moods")

	api.Post("/", handlers.CreateMoodEntry)
	api.Get("/", handlers.GetAllMood)              // ambil semua data
	api.Get("/:user_id", handlers.GetDataByUserId) // ambil semua data by userId
}
