package calendar

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/teploff/otus/calendar/domain/entity"
	"github.com/teploff/otus/calendar/domain/service"
)

// Endpoints for calendar service.
type Endpoints struct {
	CreateEvent     endpoint.Endpoint
	UpdateEvent     endpoint.Endpoint
	DeleteEvent     endpoint.Endpoint
	GetDailyEvent   endpoint.Endpoint
	GetWeeklyEvent  endpoint.Endpoint
	GetMonthlyEvent endpoint.Endpoint
}

func makeCreateEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(CreateEventRequest)

		if err = svc.CreateEvent(ctx, entity.Event{
			ShortDescription: req.ShortDescription,
			Date:             req.Date,
			Duration:         req.Duration,
			FullDescription:  req.FullDescription,
			RemindBefore:     req.RemindBefore,
			UserID:           req.UserID,
		}); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeUpdateEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateEventRequest)

		if err = svc.UpdateEvent(ctx, entity.Event{
			ID:               req.EventID,
			ShortDescription: req.ShortDescription,
			Date:             req.Date,
			Duration:         req.Duration,
			FullDescription:  req.FullDescription,
			RemindBefore:     req.RemindBefore,
			UserID:           req.UserID,
		}); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeDeleteEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteEventRequest)

		if err = svc.DeleteEvent(ctx, req.EventID, req.UserID); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeGetDailyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		events, err := svc.GetDailyEvent(ctx, req.UserID, req.Date)
		if err != nil {
			return nil, err
		}

		return GetEventResponse{Events: events}, nil
	}
}

func makeGetWeeklyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		events, err := svc.GetWeeklyEvent(ctx, req.UserID, req.Date)
		if err != nil {
			return nil, err
		}

		return GetEventResponse{Events: events}, nil
	}
}

func makeGetMonthlyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		events, err := svc.GetMonthlyEvent(ctx, req.UserID, req.Date)
		if err != nil {
			return nil, err
		}

		return GetEventResponse{Events: events}, nil
	}
}

// MakeCalendarEndpoints provides endpoints for calendar service.
func MakeCalendarEndpoints(svc service.CalendarService) Endpoints {
	return Endpoints{
		CreateEvent:     makeCreateEventEndpoint(svc),
		UpdateEvent:     makeUpdateEventEndpoint(svc),
		DeleteEvent:     makeDeleteEventEndpoint(svc),
		GetDailyEvent:   makeGetDailyEventEndpoint(svc),
		GetWeeklyEvent:  makeGetWeeklyEventEndpoint(svc),
		GetMonthlyEvent: makeGetMonthlyEventEndpoint(svc),
	}
}
