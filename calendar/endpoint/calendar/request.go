package calendar

import "time"

type Event struct {
	ShortDescription string    `json:"short_description"`
	Date             time.Time `json:"date"`
	Duration         int64     `json:"duration"`
	FullDescription  string    `json:"full_description"`
	RemindBefore     int64     `json:"remind_before"`
}

type CreateEventRequest struct {
	UserID string `json:"user_id"`
	Event
}

type UpdateEventRequest struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Event
}

type DeleteEventRequest struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
}

type DateRequest struct {
	UserID string    `json:"user_id"`
	Date   time.Time `json:"date"`
}
