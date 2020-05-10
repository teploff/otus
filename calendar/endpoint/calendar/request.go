package calendar

import "time"

type Event struct {
	ShortDescription string        `json:"short_description"`
	Date             time.Time     `json:"date"`
	Duration         time.Duration `json:"duration"`
	FullDescription  string        `json:"full_description"`
	RemindBefore     time.Duration `json:"remind_before"`
}

type CreateEventRequest struct {
	UserID string `json:"user_id"`
	Event
}

type UpdateEventRequest struct {
	EventID string `json:"event_id"`
	Event
}

type DeleteEventRequest struct {
	EventID string `json:"event_id"`
}

type DateRequest struct {
	Date time.Time `json:"date"`
}
