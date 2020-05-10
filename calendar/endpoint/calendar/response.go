package calendar

import "github.com/teploff/otus/calendar/domain/entity"

type EmptyResponse struct {
}

type GetEventResponse struct {
	Events []entity.Event `json:"events"`
}
