package calendar

import "time"

// Event entity
type Event struct {
	ShortDescription string    `json:"short_description"`
	Date             time.Time `json:"date"`
	Duration         int64     `json:"duration"`
	FullDescription  string    `json:"full_description"`
	RemindBefore     int64     `json:"remind_before"`
}

// CreateEventRequest for CreateEvent calendar domain logic method
type CreateEventRequest struct {
	UserID string `json:"user_id"`
	Event
}

// UpdateEventRequest for UpdateEvent calendar domain logic method
type UpdateEventRequest struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
	Event
}

// DeleteEventRequest for DeleteEvent calendar domain logic method
type DeleteEventRequest struct {
	UserID  string `json:"user_id"`
	EventID string `json:"event_id"`
}

// DateRequest for GetDailyEvent, GetWeeklyEvent, GetMonthlyEvent calendar domain logic method
type DateRequest struct {
	UserID string    `json:"user_id"`
	Date   time.Time `json:"date"`
}
