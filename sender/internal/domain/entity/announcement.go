package entity

import "time"

type Announcement struct {
	UserID           string
	ShortDescription string
	FullDescription  string
	Date             time.Time
	Duration         int64
}
