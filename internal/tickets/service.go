package domain

import (
	"desafio-go-web/internal/domain"
	"github.com/gin-gonic/gin"
)

type TicketService struct {
	Repository Repository
}

func (s TicketService) GetTotalTickets(c *gin.Context, destination string) ([]domain.Ticket, error) {
	return nil, nil
}

func (s TicketService) AverageDestination(c *gin.Context, destination string) (interface{}, error) {
	return nil, nil
}
