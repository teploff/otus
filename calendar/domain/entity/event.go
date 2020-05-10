package entity

import (
	"time"
)

// Event entity
type Event struct {
	ID               string    `json:"id"`
	ShortDescription string    `json:"short_description"`
	Date             time.Time `json:"date"`
	Duration         int64     `json:"duration"`
	FullDescription  string    `json:"full_description"`
	RemindBefore     int64     `json:"remind_before"`
	UserID           string    `json:"user_id"`
}
