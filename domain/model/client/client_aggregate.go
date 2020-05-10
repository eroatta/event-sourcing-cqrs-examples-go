package client

import (
	"errors"
	"fmt"
	"time"

	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model"
	"github.com/google/uuid"
)

type Client struct {
	id          uuid.UUID
	baseVersion int
	newEvents   []model.Event
	name        string
	email       string
}

func NewClient(id uuid.UUID, name, email string) (*Client, error) {
	client := &Client{
		id:        id,
		newEvents: make([]model.Event, 0),
	}

	event := NewClientEnrolledEvent(id, time.Now(), client.NextVersion(), name, email)
	err := client.ApplyNewEvent(event)
	if err != nil {
		// TODO: log the error
		return nil, err
	}

	return client, nil
}

func NewClientFromEvents(id uuid.UUID, eventsStream []model.Event) (*Client, error) {
	client := &Client{
		id:        id,
		newEvents: make([]model.Event, 0),
	}

	for _, e := range eventsStream {
		err := client.Apply(e)
		if err != nil {
			// TODO: log the error
			return nil, err
		}
		client.baseVersion = e.Version()
	}

	return client, nil
}

func (c Client) ID() uuid.UUID {
	return c.id
}

func (c Client) BaseVersion() int {
	return c.baseVersion
}

func (c Client) NewEvents() []model.Event {
	return c.newEvents
}

func (c Client) Name() string {
	return c.name
}

func (c Client) Email() string {
	return c.email
}

func (c Client) NextVersion() int {
	return c.baseVersion + len(c.newEvents) + 1
}

func (c *Client) ApplyNewEvent(event model.Event) error {
	if event.Version() != c.NextVersion() {
		return fmt.Errorf("New event version '%d' does not match expected next version '%d'",
			event.Version(), c.NextVersion())
	}

	err := c.Apply(event)
	if err != nil {
		// TODO: log the error
		return err
	}

	c.newEvents = append(c.newEvents, event)
	return nil
}

func (c *Client) Apply(event model.Event) error {
	switch ev := event.(type) {
	case ClientEnrolledEvent:
		err := clientEnrolledEventApplier(ev, c)
		if err != nil {
			// TODO: log the error
			return err
		}
	default:
		return errors.New("Unprocessable event")
	}

	return nil
}

func clientEnrolledEventApplier(event ClientEnrolledEvent, client *Client) error {
	client.name = event.Name()
	client.email = event.Email()

	return nil
}
