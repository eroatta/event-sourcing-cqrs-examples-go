package rest

import (
	"errors"
	"net/http"

	"github.com/eroatta/event-sourcing-cqrs-examples-go/domain/model/client"
	"github.com/eroatta/event-sourcing-cqrs-examples-go/service"
	"github.com/gin-gonic/gin"
)

func NewServer(clientService *service.ClientService) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// TODO: this can be improved, by receiving a list of resources, each resource specifying
	// its supported method, endpoint and handler function
	r.POST("/clients", func(c *gin.Context) {
		clientsPostHandler(c, clientService)
	})

	return r
}

func clientsPostHandler(c *gin.Context, clientService *service.ClientService) {
	var cmd clientsPostCommand
	if err := c.ShouldBindJSON(&cmd); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New("invalid values"))
	}

	enrollClientCommand := service.EnrollClientCommand{
		Name:  cmd.Name,
		Email: cmd.Email,
	}

	client, err := clientService.Process(enrollClientCommand)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	response := toDTO(*client)
	c.JSON(http.StatusCreated, response)
}

type clientsPostCommand struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type clientDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func toDTO(entity client.Client) clientDTO {
	return clientDTO{
		ID:    entity.ID().String(),
		Name:  entity.Name(),
		Email: entity.Email(),
	}
}
