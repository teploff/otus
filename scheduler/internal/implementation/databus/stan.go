package databus

import (
	"github.com/francoispqt/gojay"
	"github.com/nats-io/stan.go"
	"github.com/teploff/otus/scheduler/internal/domain/databus"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
)

type stanDataBus struct {
	conn stan.Conn
}

func NewStanDataBus(stanConn stan.Conn) databus.DataBus {
	return &stanDataBus{
		conn: stanConn,
	}
}

// PublishMeasurement provides transportation for measurement to stan
func (s *stanDataBus) PublishNotification(event *entity.Event) error {
	msg, err := gojay.MarshalJSONObject(NewEncodeNotification(event))
	if err != nil {
		return err
	}
	_, err = s.conn.PublishAsync("notifications", msg, func(_ string, _ error) {})

	return err
}
