package notifier

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/teploff/otus/sender/internal/domain/service"
)

// makeSendNotificationEndpoint create SendNotification endpoint.
func makeSendNotificationEndpoint(service service.NotifierService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SendNotificationRequest)
		return nil, service.Send(req.Announcement)
	}
}

type Endpoints struct {
	SendNotification endpoint.Endpoint
}

func MakeNotifierEndpoints(svc service.NotifierService) Endpoints {
	return Endpoints{
		SendNotification: makeSendNotificationEndpoint(svc),
	}
}
