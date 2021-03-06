package service

import (
	"context"
	"github.com/teploff/otus/scheduler/internal/domain/databus"
	"github.com/teploff/otus/scheduler/internal/domain/repository"
	"github.com/teploff/otus/scheduler/internal/domain/service"
	"go.uber.org/zap"
	"time"
)

// tickerReminder implements ReminderService interface
type tickerReminder struct {
	tk         *time.Ticker
	done       chan struct{}
	repository repository.EventRepository
	dataBus    databus.DataBus
	logger     *zap.Logger
}

// NewTickerReminder gets ticker reminder instance
func NewTickerReminder(duration time.Duration, eventRepository repository.EventRepository, dataBus databus.DataBus,
	logger *zap.Logger) service.ReminderService {
	return &tickerReminder{
		tk:         time.NewTicker(duration),
		done:       make(chan struct{}, 1),
		repository: eventRepository,
		dataBus:    dataBus,
		logger:     logger,
	}
}

// Run starts ticker reminder. Ticker ticks with duration interval
func (t tickerReminder) Run(ctx context.Context) {
	for {
		select {
		case <-t.tk.C:
			events, err := t.repository.GetEventsRequiringNotice(ctx)
			if err != nil {
				t.logger.Error("on query to get events which required notice", zap.Error(err))
			}

			for _, event := range events {
				if err = t.dataBus.PublishNotification(&event); err != nil {
					t.logger.Error("on push notification via data bus", zap.Error(err))
				}
			}

			if len(events) > 0 {
				if err = t.repository.ConfirmEvents(ctx, events); err != nil {
					t.logger.Error("on query to confirm events", zap.Error(err))
				}
			}

			if err = t.repository.CleanExpiredEvents(ctx); err != nil {
				t.logger.Error("on query to clean expired events", zap.Error(err))
			}
		case <-t.done:
			return
		default:
			time.Sleep(time.Millisecond * 100)
		}
	}
}

// Stop closing ticker reminder
func (t tickerReminder) Stop() {
	t.done <- struct{}{}
}
