package grpc

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/teploff/otus/calendar/endpoint/calendar"
	"github.com/teploff/otus/calendar/transport/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const maxReceivedMsgSize = 1024 * 1024 * 20

type server struct {
	createEvent     kitgrpc.Handler
	updateEvent     kitgrpc.Handler
	deleteEvent     kitgrpc.Handler
	getDailyEvent   kitgrpc.Handler
	getWeeklyEvent  kitgrpc.Handler
	getMonthlyEvent kitgrpc.Handler
}

func (s server) CreateEvent(ctx context.Context, request *pb.CreateRequest) (*pb.EmptyResponse, error) {
	_, response, err := s.createEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response.(*pb.EmptyResponse), nil
}

func (s server) UpdateEvent(ctx context.Context, request *pb.UpdateEventRequest) (*pb.EmptyResponse, error) {
	_, response, err := s.updateEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response.(*pb.EmptyResponse), nil
}

func (s server) DeleteEvent(ctx context.Context, request *pb.DeleteEventRequest) (*pb.EmptyResponse, error) {
	_, response, err := s.deleteEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return response.(*pb.EmptyResponse), nil
}

func (s server) GetDailyEvent(ctx context.Context, request *pb.DateRequest) (*pb.GetEventResponse, error) {
	_, response, err := s.getDailyEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	events := response.(calendar.GetEventResponse).Events
	pbEvents := make([]*pb.DbEvent, 0, len(events))
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.DbEvent{
			Id:               event.ID,
			ShortDescription: event.ShortDescription,
			Date:             event.Date.Unix(),
			Duration:         event.Duration,
			FullDescription:  event.FullDescription,
			RemindBefore:     event.RemindBefore,
		})
	}
	return &pb.GetEventResponse{Events: pbEvents}, nil
}

func (s server) GetWeeklyEvent(ctx context.Context, request *pb.DateRequest) (*pb.GetEventResponse, error) {
	_, response, err := s.getWeeklyEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	events := response.(calendar.GetEventResponse).Events
	pbEvents := make([]*pb.DbEvent, 0, len(events))
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.DbEvent{
			Id:               event.ID,
			ShortDescription: event.ShortDescription,
			Date:             event.Date.Unix(),
			Duration:         event.Duration,
			FullDescription:  event.FullDescription,
			RemindBefore:     event.RemindBefore,
		})
	}
	return &pb.GetEventResponse{Events: pbEvents}, nil
}

func (s server) GetMonthlyEvent(ctx context.Context, request *pb.DateRequest) (*pb.GetEventResponse, error) {
	_, response, err := s.getMonthlyEvent.ServeGRPC(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	events := response.(calendar.GetEventResponse).Events
	pbEvents := make([]*pb.DbEvent, 0, len(events))
	for _, event := range events {
		pbEvents = append(pbEvents, &pb.DbEvent{
			Id:               event.ID,
			ShortDescription: event.ShortDescription,
			Date:             event.Date.Unix(),
			Duration:         event.Duration,
			FullDescription:  event.FullDescription,
			RemindBefore:     event.RemindBefore,
		})
	}
	return &pb.GetEventResponse{Events: pbEvents}, nil
}

// NewGRPCServer instance of gRPC server.
func NewGRPCServer(endpoints calendar.Endpoints, errLogger log.Logger) *grpc.Server {
	options := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transport.NewLogErrorHandler(errLogger)),
	}

	srv := &server{
		createEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.CreateEvent,
			decodeCreateEventRequest,
			encodeEmptyResponse,
			options...,
		), errLogger),
		updateEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.UpdateEvent,
			decodeUpdateEventRequest,
			encodeEmptyResponse,
			options...,
		), errLogger),
		deleteEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.DeleteEvent,
			decodeDeleteEventRequest,
			encodeEmptyResponse,
			options...,
		), errLogger),
		getDailyEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.GetDailyEvent,
			decodeDateRequest,
			encodeGetEventResponse,
			options...,
		), errLogger),
		getWeeklyEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.GetWeeklyEvent,
			decodeDateRequest,
			encodeGetEventResponse,
			options...,
		), errLogger),
		getMonthlyEvent: newRecoveryGRPCHandler(kitgrpc.NewServer(
			endpoints.GetMonthlyEvent,
			decodeDateRequest,
			encodeGetEventResponse,
			options...,
		), errLogger),
	}

	baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor), grpc.MaxRecvMsgSize(maxReceivedMsgSize))
	pb.RegisterCalendarServer(baseServer, srv)

	return baseServer
}

func decodeCreateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.CreateRequest)

	return calendar.CreateEventRequest{
		UserID: request.UserId,
		Event: calendar.Event{
			ShortDescription: request.Event.ShortDescription,
			Date:             time.Unix(request.Event.Date, 0),
			Duration:         request.Event.Duration,
			FullDescription:  request.Event.FullDescription,
			RemindBefore:     request.Event.RemindBefore,
		},
	}, nil
}

func decodeUpdateEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.UpdateEventRequest)

	return calendar.UpdateEventRequest{
		UserID:  request.UserId,
		EventID: request.EventId,
		Event: calendar.Event{
			ShortDescription: request.Event.ShortDescription,
			Date:             time.Unix(request.Event.Date, 0),
			Duration:         request.Event.Duration,
			FullDescription:  request.Event.FullDescription,
			RemindBefore:     request.Event.RemindBefore,
		},
	}, nil
}

func decodeDeleteEventRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.DeleteEventRequest)

	return calendar.DeleteEventRequest{
		UserID:  request.UserId,
		EventID: request.EventId,
	}, nil
}

func decodeDateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	request := grpcReq.(*pb.DateRequest)

	return calendar.DateRequest{
		UserID: request.UserId,
		Date:   time.Unix(request.Date, 0),
	}, nil
}

func encodeEmptyResponse(_ context.Context, _ interface{}) (interface{}, error) {
	return &pb.EmptyResponse{}, nil
}

func encodeGetEventResponse(_ context.Context, grpcResp interface{}) (interface{}, error) {
	return grpcResp.(*pb.GetEventResponse), nil
}

//recoveryGRPCHandler wrap gRPC server, recover them if panic was fired.
type recoveryGRPCHandler struct {
	next   kitgrpc.Handler
	logger log.Logger
}

func newRecoveryGRPCHandler(next kitgrpc.Handler, logger log.Logger) *recoveryGRPCHandler {
	return &recoveryGRPCHandler{next: next, logger: logger}
}

func (rh *recoveryGRPCHandler) ServeGRPC(ctx context.Context, req interface{}) (context.Context, interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				_ = rh.logger.Log("msg", "gRPC server panic recover", "text", err.Error())
			}
		}
	}()

	return rh.next.ServeGRPC(ctx, req)
}
