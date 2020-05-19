package service

import (
	"fmt"
	"github.com/teploff/otus/sender/internal/domain/entity"
	"github.com/teploff/otus/sender/internal/domain/service"
	"go.uber.org/zap"
)

type notifierService struct {
	logger *zap.Logger
}

func NewNotifierService(logger *zap.Logger) service.NotifierService {
	return &notifierService{logger: logger}
}

func (n notifierService) Send(a entity.Announcement) error {
	n.logger.Info(fmt.Sprintf("notification for %s. Date: %s. Short description: %s. Duration: %d",
		a.UserID, a.Date, a.ShortDescription, a.Duration))

	return nil
}
