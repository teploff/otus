package service

import "context"

type Reminder interface {
	Run(ctx context.Context)
	Stop()
}
