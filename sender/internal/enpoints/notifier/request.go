package notifier

import (
	"time"
)

//SendNotificationRequest request for SendNotification endpoint.
//easyjson:json
type SendNotificationRequest struct {
	UserID           string    `json:"user_id"`
	ShortDescription string    `json:"short_description"`
	FullDescription  string    `json:"full_description"`
	Date             time.Time `json:"date"`
	Duration         int64     `json:"duration"`
}
