package databus

import (
	"github.com/francoispqt/gojay"
	"github.com/teploff/otus/scheduler/internal/domain/entity"
	"time"
)

// encodeNotification implements MarshalerJSONObject interface for Event encoding
type encodeNotification struct {
	entity.Event
}

func NewEncodeNotification(event *entity.Event) gojay.MarshalerJSONObject {
	return &encodeNotification{*event}
}

func (e *encodeNotification) IsNil() bool {
	return e == nil
}

func (e *encodeNotification) MarshalJSONObject(enc *gojay.Encoder) {
	enc.StringKey("user_id", e.UserID)
	enc.StringKey("short_description", e.ShortDescription)
	enc.StringKey("full_description", e.FullDescription)
	enc.TimeKey("date", &e.Date, time.RFC3339)
	enc.Int64Key("duration", e.Duration)
}
