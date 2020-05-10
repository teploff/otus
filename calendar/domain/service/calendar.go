package service

type CalendarService interface {
	CreateEvent() error
	UpdateEvent() error
	DeleteEvent() error
	GetDailyEvent() error
	GetWeeklyEvent() error
	GetMonthlyEvent() error
}
