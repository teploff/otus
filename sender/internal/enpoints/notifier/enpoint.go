package notifier

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/teploff/otus/sender/internal/domain/entity"
	"github.com/teploff/otus/sender/internal/domain/service"
)

// Endpoints of NotifierService
type Endpoints struct {
	SendNotification endpoint.Endpoint
}

// makeSendNotificationEndpoint create SendNotification endpoint.
func makeSendNotificationEndpoint(service service.NotifierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendNotificationRequest)
		return nil, service.Send(entity.Announcement{
			UserID:           req.UserID,
			ShortDescription: req.ShortDescription,
			FullDescription:  req.FullDescription,
			Date:             req.Date,
			Duration:         req.Duration,
		})
	}
}

func MakeNotifierEndpoints(svc service.NotifierService) Endpoints {
	return Endpoints{
		SendNotification: makeSendNotificationEndpoint(svc),
	}
}
