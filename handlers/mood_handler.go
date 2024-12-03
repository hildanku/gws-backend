package handlers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/config"
	"github.com/gws-app/gws-backend/models"
	"github.com/gws-app/gws-backend/utils"
	"google.golang.org/api/iterator"
)

func CreateMoodEntry(ctx *fiber.Ctx) error {
	mood := new(models.Mood)

	mood.UserID = ctx.FormValue("user_id")

	if err := ctx.BodyParser(mood); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Invalid Input",
			Data:   nil,
		})
	}

	// Parsing manual activities
	if activities := ctx.FormValue("activities"); activities != "" {
		if err := json.Unmarshal([]byte(activities), &mood.Activities); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
				Code:   fiber.StatusBadRequest,
				Status: "Invalid activities format",
				Data:   nil,
			})
		}
	}

	if mood.Mood == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Mood is required",
			Data:   nil,
		})
	}

	if mood.UserID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "User ID is required",
			Data:   nil,
		})
	}
	log.Println(mood.UserID)
	//	uid := mood.UserID

	// catch voice note
	voiceNoteHeader, err := ctx.FormFile("voice_note_url")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.MoodResponse{
			Code:   fiber.StatusBadRequest,
			Status: "Voice Note is required",
			Data:   nil,
		})
	}
	//
	voiceNoteFile, err := voiceNoteHeader.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.MoodResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Failed to open voice note file",
			Data:   nil,
		})
	}

	voiceNoteURL, err := utils.UploadGCS(voiceNoteFile, mood.UserID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(models.MoodResponse{
			Code:   fiber.StatusInternalServerError,
			Status: "Failed to upload voiceNote to cloud storage",
			Data:   nil,
		})
	}

	mood.VoiceNoteURL = voiceNoteURL

	mood.CreatedAt = time.Now()

	_, _, err = config.FirestoreClient.Collection("mood_entries").Add(context.Background(), mood)
	if err != nil {
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

// Ambil semua data
func GetAllMood(ctx *fiber.Ctx) error {
	iter := config.FirestoreClient.Collection("mood_entries").Documents(context.Background())
	defer iter.Stop()

	var moods []models.Mood
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(models.MoodResponse{
				Code:   fiber.StatusInternalServerError,
				Status: "Failed to fetch mood entries",
				Data:   nil,
			})
		}

		var mood models.Mood
		doc.DataTo(&mood)
		moods = append(moods, mood)
	}

	if len(moods) == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(models.MoodResponse{
			Code:   fiber.StatusNotFound,
			Status: "No mood entries found",
			Data:   nil,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.MoodResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   moods,
	})
}
