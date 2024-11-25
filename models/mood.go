package models

import "time"

type Mood struct {
	UserID       string              `json:"user_id"`
	Emotion      string              `json:"emotion"`
	Activities   map[string][]string `json:"activities,omitempty"`
	Note         string              `json:"note,omitempty"`
	VoiceNoteURL string              `json:"voice_note_url,omitempty"`
	CreatedAt    time.Time           `json:"created_at,omitempty"`
}
