package handlers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/config"
	"github.com/gws-app/gws-backend/models"
	"log"
	"time"
)

func CreateMoodEntry(ctx *fiber.Ctx) error {
	mood := new(models.Mood)
	if err := ctx.BodyParser(mood); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Input",
			Data:   nil,
		})
	}

	if mood.Emotion == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Emotion is required",
			Data:   nil,
		})
	}

	mood.CreatedAt = time.Now()
	log.Println("Saving mood entry to Firestore:", mood)
	_, _, err := config.FirestoreClient.Collection("mood_entries").Add(context.Background(), mood)
	if err != nil {
		log.Println("Error saving mood entry:", err) // Log the actual error from Firestore
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.MoodResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Failed to create mood entry",
			Data:   nil,
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(models.MoodResponse{
		Code:   fiber.StatusCreated,
		Status: "Mood Created",
		Data:   mood,
	})

}
