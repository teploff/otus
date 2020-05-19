package service

import (
	"fmt"
	"github.com/teploff/otus/sender/internal/domain/entity"
	"github.com/teploff/otus/sender/internal/domain/service"
	"go.uber.org/zap"
)

// notifierService implements NotifierService interface
type notifierService struct {
	logger *zap.Logger
}

// NewNotifierService get instance of notifier service
func NewNotifierService(logger *zap.Logger) service.NotifierService {
	return &notifierService{logger: logger}
}

// Send logging notifications
func (n notifierService) Send(a entity.Announcement) error {
	n.logger.Info(fmt.Sprintf("notification for %s. Date: %s. Short description: %s. Duration: %d",
		a.UserID, a.Date, a.ShortDescription, a.Duration))

	return nil
}
