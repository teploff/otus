package entity

import "time"

// Announcement entity of notification
type Announcement struct {
	UserID           string
	ShortDescription string
	FullDescription  string
	Date             time.Time
	Duration         int64
}
