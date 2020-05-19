package stan

import (
	"context"
	"github.com/teploff/otus/sender/internal/enpoints/notifier"
	"time"

	"github.com/francoispqt/gojay"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	"github.com/nats-io/stan.go"
	kitstan "github.com/teploff/otus/sender/internal/infrastructure/stan"
)

// Stan is client to the NATS server.
type Stan struct {
	subscriptions []stan.Subscription
}

// NewStan create instance of Stan.
func NewStan() *Stan {
	return &Stan{}
}

// Serve starts listening stan.
func (s *Stan) Serve(conn stan.Conn, endpoints notifier.Endpoints, errLog log.Logger) error {
	processMeasurementHandler := kitstan.NewSubscriber(
		endpoints.SendNotification,
		DecodeSendNotificationRequest,
		kitstan.EncodeJSONResponse,
		kitstan.SubscriberErrorHandler(transport.NewLogErrorHandler(errLog)),
	)

	subscriptionOptions := []stan.SubscriptionOption{
		stan.SetManualAckMode(),
		stan.AckWait(time.Second * 1),
		stan.MaxInflight(25),
	}

	sub, err := conn.Subscribe("notifications", processMeasurementHandler.ServeMsg(&conn),
		append(subscriptionOptions, stan.DurableName("sender"))...)
	if err != nil {
		return err
	}
	// make unlimited subscription way: unlimited count messages and unlimited message size
	if err = sub.SetPendingLimits(-1, -1); err != nil {
		return err
	}
	s.subscriptions = append(s.subscriptions, sub)

	return nil
}

// Stop closes all subscribe connections.
func (s *Stan) Stop() {
	for _, sub := range s.subscriptions {
		sub.Close()
	}
}

// DecodeSendNotificationRequest decodes SendNotification request.
func DecodeSendNotificationRequest(_ context.Context, msg *stan.Msg) (interface{}, error) {
	var request notifier.SendNotificationRequest

	if err := gojay.UnmarshalJSONObject(msg.Data, &request); err != nil {
		return nil, err
	}

	return request, nil
}
