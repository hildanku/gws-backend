package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gws-app/gws-backend/models"
)

func GetAllNews(ctx *fiber.Ctx) error {
	resp, err := http.Get("https://sandbox.api.service.nhs.uk/nhs-website-content/mental-health")
	if err != nil {
		log.Println("Error fetching data:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data from the API",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response body",
		})
	}

	// Parse the JSON response
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Println("Error unmarshalling response body:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse JSON response",
		})
	}

	// Extract relevant data
	extractedContents := []models.MentalHealthContent{}
	if hasParts, ok := apiResponse["hasPart"].([]interface{}); ok {
		for _, item := range hasParts {
			if content, ok := item.(map[string]interface{}); ok {
				text := ""
				if hasPart, ok := content["hasPart"].([]interface{}); ok {
					for _, part := range hasPart {
						if pageElement, ok := part.(map[string]interface{}); ok {
							if t, exists := pageElement["text"].(string); exists {
								text = t
								break
							}
						}
					}
				}
				extractedContents = append(extractedContents, models.MentalHealthContent{
					Headline:    content["headline"].(string),
					Description: content["description"].(string),
					URL:         content["url"].(string),
					Text:        text,
				})
			}
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(models.NewsResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   extractedContents,
	})
}

func GetRecommendation(ctx *fiber.Ctx) error {
	resp, err := http.Get("https://sandbox.api.service.nhs.uk/nhs-website-content/mental-health")
	if err != nil {
		log.Println("Error fetching data:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch data from the API",
		})
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to read response body",
		})
	}

	// Parse the JSON response
	var apiResponse map[string]interface{}
	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Println("Error unmarshalling response body:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse JSON response",
		})
	}

	// Extract relevant data
	extractedContents := []models.MentalHealthContent{}
	if hasParts, ok := apiResponse["hasPart"].([]interface{}); ok {
		for _, item := range hasParts {
			if content, ok := item.(map[string]interface{}); ok {
				extractedContents = append(extractedContents, models.MentalHealthContent{
					Headline:    content["headline"].(string),
					Description: content["description"].(string),
					URL:         content["url"].(string),
				})
			}
		}
	}

	// response := models.RecommendationResponse{
	//	Data: extractedContents,
	//	}

	// return ctx.Status(fiber.StatusOK).JSON(response)

	return ctx.Status(fiber.StatusOK).JSON(models.NewsResponse{
		Code:   fiber.StatusOK,
		Status: "success",
		Data:   extractedContents,
	})
}
