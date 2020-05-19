package notifier

import (
	"github.com/francoispqt/gojay"
	"github.com/teploff/otus/sender/internal/domain/entity"
	"time"
)

//SendNotificationRequest request for SendNotification endpoint.
type SendNotificationRequest struct {
	entity.Announcement
}

func (s *SendNotificationRequest) UnmarshalJSONObject(dec *gojay.Decoder, key string) error {
	switch key {
	case "user_id":
		return dec.String(&s.UserID)
	case "short_description":
		return dec.String(&s.ShortDescription)
	case "full_description":
		return dec.String(&s.FullDescription)
	case "date":
		return dec.Time(&s.Date, time.RFC3339)
	case "duration":
		return dec.Int64(&s.Duration)
	}

	return nil
}

func (s *SendNotificationRequest) NKeys() int {
	return 5
}
