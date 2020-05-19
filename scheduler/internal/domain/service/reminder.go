package service

import "context"

// ReminderService provides reminder to db query
type ReminderService interface {
	Run(ctx context.Context)
	Stop()
}
