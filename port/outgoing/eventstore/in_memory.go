package eventstore

import (
	"fmt"

	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type InMemoryEventStore struct {
	eventStore map[uuid.UUID][]model.Event
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{
		eventStore: make(map[uuid.UUID][]model.Event),
	}
}

func (es *InMemoryEventStore) Store(aggregateID uuid.UUID, newEvents []model.Event, baseVersion int) error {
	events := es.eventStore[aggregateID]
	// TODO: consider optimistic locking
	es.eventStore[aggregateID] = append(events, newEvents...)
	for _, e := range newEvents {
		logrus.WithField("aggregate", e.AggregateID).Debug(fmt.Sprintf("stored event %v", e))
	}
	logrus.Debug(fmt.Sprintf("events store contains %v", es.eventStore))

	return nil
}

func (es *InMemoryEventStore) Load(aggregateID uuid.UUID) ([]model.Event, error) {
	events := es.eventStore[aggregateID]

	copyOf := make([]model.Event, len(events))
	copy(copyOf, events)

	return copyOf, nil
}
