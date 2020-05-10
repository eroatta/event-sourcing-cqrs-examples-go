package model

import "github.com/google/uuid"

type Aggregate interface {
	ID() uuid.UUID
	BaseVersion() int
	NewEvents() []Event
	NextVersion() int
	ApplyNewEvent(event Event) error
	Apply(event Event)
}

// TODO: add NewAggregateFromEvents
