package service

import (
	"errors"

	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model"
	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model/client"
	"github.com/google/uuid"
)

var (
	ErrClientNotFound = errors.New("Client not found")
)

type ClientService struct {
	eventStore model.EventStore
}

func NewClientService(eventStore model.EventStore) *ClientService {
	return &ClientService{
		eventStore: eventStore,
	}
}

func (cs *ClientService) LoadClient(id uuid.UUID) (*client.Client, error) {
	eventStream, err := cs.eventStore.Load(id)
	if err != nil {
		// TODO: log the error
		return nil, err
	}

	if len(eventStream) == 0 {
		return nil, ErrClientNotFound
	}

	return client.NewClientFromEvents(id, eventStream)
}

// TODO: this can be improved by using a func that forces to process a command, always returning
// a Client and an error...
func (cs *ClientService) Process(command interface{}) (*client.Client, error) {
	switch cmd := command.(type) {
	case EnrollClientCommand:
		c, err := client.NewClient(uuid.New(), cmd.Name, cmd.Email)
		if err != nil {
			// TODO: log the error
			return nil, err
		}
		err = cs.storeEvents(c)
		if err != nil {
			// TODO: log the error
			return nil, err
		}

		return c, nil
	}

	return nil, nil
}

func (cs *ClientService) storeEvents(c *client.Client) error {
	return cs.eventStore.Store(c.ID(), c.NewEvents(), c.BaseVersion())
}
