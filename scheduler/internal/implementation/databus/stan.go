package databus

import (
	"github.com/francoispqt/gojay"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/scheduler/internal/domain/databus"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
)

// stanDataBus implements DataBase interface to provide write to stan (nats-streaming)
type stanDataBus struct {
	conn stan.Conn
}

// NewStanDataBus gets databus instance
func NewStanDataBus(stanConn stan.Conn) databus.DataBus {
	return &stanDataBus{
		conn: stanConn,
	}
}

// PublishNotification provides transportation for notification to stan
func (s *stanDataBus) PublishNotification(event *entity.Event) error {
	msg, err := gojay.MarshalJSONObject(NewEncodeNotification(event))
	if err != nil {
		return err
	}
	_, err = s.conn.PublishAsync("notifications", msg, func(_ string, _ error) {})

	return err
}
