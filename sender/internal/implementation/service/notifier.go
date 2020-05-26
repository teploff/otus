package service

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/teploff/otus/sender/internal/domain/entity"
	"github.com/teploff/otus/sender/internal/domain/service"
)

// notifierService implements NotifierService interface
type notifierService struct {
	logger log.Logger
}

// NewNotifierService get instance of notifier service
func NewNotifierService(logger log.Logger) service.NotifierService {
	return &notifierService{logger: logger}
}

// Send logging notifications
func (n notifierService) Send(a entity.Announcement) error {
	fmt.Println("got")
	return n.logger.Log(fmt.Sprintf("notification for %s. Date: %s. Short description: %s. Duration: %d",
		a.UserID, a.Date, a.ShortDescription, a.Duration))
}
