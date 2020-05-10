package client

import (
	"time"

	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model"
	"github.com/google/uuid"
)

type ClientEnrolledEvent struct {
	model.Event
	name  string
	email string
}

func NewClientEnrolledEvent(aggregateID uuid.UUID, timestamp time.Time, version int, name, email string) ClientEnrolledEvent {
	return ClientEnrolledEvent{
		Event: model.NewEvent(aggregateID, timestamp, version),
		name:  name,
		email: email,
	}
}

func (e ClientEnrolledEvent) Name() string {
	return e.name
}

func (e ClientEnrolledEvent) Email() string {
	return e.email
}
