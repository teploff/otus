package calendar

import "github.com/teploff/otus/calendar/domain/entity"

// EmptyResponse for CreateEvent, DeleteEventRequest, UpdateEvent calendar domain logic method
type EmptyResponse struct {
}

// GetEventResponse for GetDailyEvent, GetWeeklyEvent, GetMonthlyEvent calendar domain logic method
type GetEventResponse struct {
	Events []entity.Event `json:"events"`
}
