package databus

import "github.com/teploff/otus/scheduler/internal/domain/entity"

// DataBus provides transportation for data
type DataBus interface {
	PublishNotification(event *entity.Event) error
}
