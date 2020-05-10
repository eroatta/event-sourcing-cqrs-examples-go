package model

import (
	"time"

	"github.com/google/uuid"
)

type Event interface {
	AggregateID() uuid.UUID
	Timestamp() time.Time
	Version() int
}

func NewEvent(aggregateID uuid.UUID, timestamp time.Time, version int) Event {
	return event{
		id:        aggregateID,
		timestamp: timestamp,
		version:   version,
	}
}

type event struct {
	id        uuid.UUID
	timestamp time.Time
	version   int
}

func (e event) AggregateID() uuid.UUID {
	return e.id
}

func (e event) Timestamp() time.Time {
	return e.timestamp
}

func (e event) Version() int {
	return e.version
}

type EventStore interface {
	Store(aggregateID uuid.UUID, newEvents []Event, baseVersion int) error
	Load(aggregateID uuid.UUID) ([]Event, error)
}
