package stan

import (
	"context"

	"github.com/nats-io/stan.go"
)

// DecodeRequestFunc extracts a user-domain request object from a publisher
// request object. It's designed to be used in NATS Streaming subscribers, for subscriber-side
// endpoints. One straightforward DecodeRequestFunc could be something that
// JSON decodes from the request body to the concrete response type.
type DecodeRequestFunc func(context.Context, *stan.Msg) (request interface{}, err error)

// EncodeRequestFunc encodes the passed request object into the NATS Streaming request
// object. It's designed to be used in NATS Streaming publishers, for publisher-side
// endpoints. One straightforward EncodeRequestFunc could something that JSON
// encodes the object directly to the request payload.
type EncodeRequestFunc func(context.Context, *stan.Msg, interface{}) error

// EncodeResponseFunc encodes the passed response object to the subscriber reply.
// It's designed to be used in NATS Streaming subscribers, for subscriber-side
// endpoints. One straightforward EncodeResponseFunc could be something that
// JSON encodes the object directly to the response body.
type EncodeResponseFunc func(context.Context, *stan.Conn, interface{}) error

// DecodeResponseFunc extracts a user-domain response object from an NATS Streaming
// response object. It's designed to be used in NATS Streaming publisher, for publisher-side
// endpoints. One straightforward DecodeResponseFunc could be something that
// JSON decodes from the response payload to the concrete response type.
type DecodeResponseFunc func(context.Context, *stan.Msg) (response interface{}, err error)
