package service

import "github.com/teploff/otus/sender/internal/domain/entity"

// NotifierService provides logic to notify user of events
type NotifierService interface {
	Send(announcement entity.Announcement) error
}
