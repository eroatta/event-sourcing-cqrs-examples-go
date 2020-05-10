package main

import (
	"github.com/eroatta/event-sourcing-cqrs-examples-go/port/incoming/rest"
	"github.com/eroatta/event-sourcing-cqrs-examples-go/port/outgoing/eventstore"
	"github.com/eroatta/event-sourcing-cqrs-examples-go/service"
)

func main() {
	eventStore := eventstore.NewInMemoryEventStore()
	clientService := service.NewClientService(eventStore)

	server := rest.NewServer(clientService)
	server.Run()
}
