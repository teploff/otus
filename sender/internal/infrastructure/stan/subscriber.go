package stan

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"

	"github.com/nats-io/stan.go"
)

// Subscriber wraps an endpoint and provides stan.MsgHandler.
type Subscriber struct {
	e            endpoint.Endpoint
	dec          DecodeRequestFunc
	enc          EncodeResponseFunc
	before       []RequestFunc
	after        []SubscriberResponseFunc
	errorEncoder ErrorEncoder
	finalizer    []SubscriberFinalizerFunc
	errorHandler transport.ErrorHandler
}

// NewSubscriber constructs a new subscriber, which provides stan.MsgHandler and wraps
// the provided endpoint.
func NewSubscriber(
	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...SubscriberOption,
) *Subscriber {
	s := &Subscriber{
		e:            e,
		dec:          dec,
		enc:          enc,
		errorEncoder: DefaultErrorEncoder,
		errorHandler: transport.NewLogErrorHandler(log.NewNopLogger()),
	}
	for _, option := range options {
		option(s)
	}
	return s
}

// SubscriberOption sets an optional parameter for subscribers.
type SubscriberOption func(*Subscriber)

// SubscriberBefore functions are executed on the publisher request object before the
// request is decoded.
func SubscriberBefore(before ...RequestFunc) SubscriberOption {
	return func(s *Subscriber) { s.before = append(s.before, before...) }
}

// SubscriberAfter functions are executed on the subscriber reply after the
// endpoint is invoked, but before anything is published to the reply.
func SubscriberAfter(after ...SubscriberResponseFunc) SubscriberOption {
	return func(s *Subscriber) { s.after = append(s.after, after...) }
}

// SubscriberErrorEncoder is used to encode errors to the subscriber reply
// whenever they're encountered in the processing of a request. Clients can
// use this to provide custom error formatting. By default,
// errors will be published with the DefaultErrorEncoder.
func SubscriberErrorEncoder(ee ErrorEncoder) SubscriberOption {
	return func(s *Subscriber) { s.errorEncoder = ee }
}

// SubscriberErrorLogger is used to log non-terminal errors. By default, no errors
// are logged. This is intended as a diagnostic measure. Finer-grained control
// of error handling, including logging in more detail, should be performed in a
// custom SubscriberErrorEncoder which has access to the context.
// Deprecated: Use SubscriberErrorHandler instead.
func SubscriberErrorLogger(logger log.Logger) SubscriberOption {
	return func(s *Subscriber) { s.errorHandler = transport.NewLogErrorHandler(logger) }
}

// SubscriberErrorHandler is used to handle non-terminal errors. By default, non-terminal errors
// are ignored. This is intended as a diagnostic measure. Finer-grained control
// of error handling, including logging in more detail, should be performed in a
// custom SubscriberErrorEncoder which has access to the context.
func SubscriberErrorHandler(errorHandler transport.ErrorHandler) SubscriberOption {
	return func(s *Subscriber) { s.errorHandler = errorHandler }
}

// SubscriberFinalizer is executed at the end of every request from a publisher through NATS Streaming.
// By default, no finalizer is registered.
func SubscriberFinalizer(f ...SubscriberFinalizerFunc) SubscriberOption {
	return func(s *Subscriber) { s.finalizer = f }
}

// ServeMsg provides stan.MsgHandler.
func (s Subscriber) ServeMsg(sc *stan.Conn) func(msg *stan.Msg) {
	return func(msg *stan.Msg) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		if len(s.finalizer) > 0 {
			defer func() {
				for _, f := range s.finalizer {
					f(ctx, msg)
				}
			}()
		}

		for _, f := range s.before {
			ctx = f(ctx, msg)
		}

		// acknowledgment despite errors
		defer func() {
			if err := msg.Ack(); err != nil {
				s.errorHandler.Handle(ctx, err)
				s.errorEncoder(ctx, err)
				return
			}
		}()

		request, err := s.dec(ctx, msg)
		if err != nil {
			s.errorHandler.Handle(ctx, err)
			s.errorEncoder(ctx, err)
			return
		}

		response, err := s.e(ctx, request)
		if err != nil {
			s.errorHandler.Handle(ctx, err)
			s.errorEncoder(ctx, err)
			return
		}

		for _, f := range s.after {
			ctx = f(ctx, sc)
		}

		if err = s.enc(ctx, sc, response); err != nil {
			s.errorHandler.Handle(ctx, err)
			s.errorEncoder(ctx, err)
			return
		}
	}
}

// ErrorEncoder is responsible for encoding an error to the subscriber reply.
// Users are encouraged to use custom ErrorEncoders to encode errors to
// their replies, and will likely want to pass and check for their own error
// types.
type ErrorEncoder func(ctx context.Context, err error)

// ServerFinalizerFunc can be used to perform work at the end of an request
// from a publisher, after the response has been written to the publisher. The principal
// intended use is for request logging.
type SubscriberFinalizerFunc func(ctx context.Context, msg *stan.Msg)

// NopRequestDecoder is a DecodeRequestFunc that can be used for requests that do not
// need to be decoded, and simply returns nil, nil.
func NopRequestDecoder(_ context.Context, _ *stan.Msg) (interface{}, error) {
	return nil, nil
}

// EncodeJSONResponse is a EncodeResponseFunc that serializes the response as a
// JSON object to the subscriber reply. Many JSON-over services can use it as
// a sensible default.
func EncodeJSONResponse(_ context.Context, sc *stan.Conn, response interface{}) error {
	_, err := json.Marshal(response)
	if err != nil {
		return err
	}
	//sc.PublishAsync(reply, b)

	return nil
}

// DefaultErrorEncoder writes the error to the subscriber reply.
func DefaultErrorEncoder(_ context.Context, err error) {
	logger := log.NewNopLogger()

	type Response struct {
		Error string `json:"err"`
	}

	var response Response

	response.Error = err.Error()

	_, err = json.Marshal(response)
	if err != nil {
		logger.Log("err", err)
		return
	}
}