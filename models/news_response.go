package models

type NewsResponse struct {
	Code   int                   `json:"code"`
	Status string                `json:"status"`
	Data   []MentalHealthContent `json:"contents"`
}
