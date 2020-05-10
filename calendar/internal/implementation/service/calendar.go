package service

import (
	"fmt"
	"github.com/teploff/otus/calendar/domain/service"
)

type calendarService struct{}

func NewCalendarService() service.CalendarService {
	return &calendarService{}
}

func (c calendarService) CreateEvent() error {
	fmt.Println("implement me")
	return nil
}

func (c calendarService) UpdateEvent() error {
	fmt.Println("implement me")
	return nil
}

func (c calendarService) DeleteEvent() error {
	fmt.Println("implement me")
	return nil
}

func (c calendarService) GetDailyEvent() error {
	fmt.Println("implement me")
	return nil
}

func (c calendarService) GetWeeklyEvent() error {
	fmt.Println("implement me")
	return nil
}

func (c calendarService) GetMonthlyEvent() error {
	fmt.Println("implement me")
	return nil
}
