package calendar

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
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

		fmt.Println("Get: ", req)

		if err = svc.CreateEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeUpdateEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UpdateEventRequest)

		fmt.Println("Get: ", req)

		if err = svc.UpdateEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeDeleteEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DeleteEventRequest)

		fmt.Println("Get: ", req)

		if err = svc.DeleteEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeGetDailyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		fmt.Println("Get: ", req)

		if err = svc.GetDailyEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeGetWeeklyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		fmt.Println("Get: ", req)

		if err = svc.GetWeeklyEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

func makeGetMonthlyEventEndpoint(svc service.CalendarService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(DateRequest)

		fmt.Println("Get: ", req)

		if err = svc.GetMonthlyEvent(); err != nil {
			return nil, err
		}

		return EmptyResponse{}, nil
	}
}

// MakeCalendarEndpoints provides endpoints for admin-panel.
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
