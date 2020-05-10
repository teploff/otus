package service

import (
	"context"
	"github.com/teploff/otus/calendar/domain/entity"
	"github.com/teploff/otus/calendar/domain/repository"
	"time"
)

type calendarService struct {
	repository repository.EventRepository
}

func NewCalendarService(repository repository.EventRepository) *calendarService {
	return &calendarService{repository: repository}
}

func (c calendarService) CreateEvent(ctx context.Context, event entity.Event) error {
	if err := c.repository.InsertEvent(ctx, event); err != nil {
		return err
	}

	return nil
}

func (c calendarService) UpdateEvent(ctx context.Context, event entity.Event) error {
	if err := c.repository.UpdateEvent(ctx, event); err != nil {
		return nil
	}
	return nil
}

func (c calendarService) DeleteEvent(ctx context.Context, eventID, userID string) error {
	if err := c.repository.DeleteEvent(ctx, eventID, userID); err != nil {
		return err
	}
	return nil
}

func (c calendarService) GetDailyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error) {
	events, err := c.repository.GetEvents(ctx, userID, date, time.Hour*24)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c calendarService) GetWeeklyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error) {
	events, err := c.repository.GetEvents(ctx, userID, date, time.Hour*24*7)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (c calendarService) GetMonthlyEvent(ctx context.Context, userID string, date time.Time) ([]entity.Event, error) {
	events, err := c.repository.GetEvents(ctx, userID, date, time.Hour*24*30)
	if err != nil {
		return nil, err
	}

	return events, nil
}
