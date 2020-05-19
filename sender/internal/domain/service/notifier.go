package service

import "github.com/teploff/otus/sender/internal/domain/entity"

type NotifierService interface {
	Send(announcement entity.Announcement) error
}
