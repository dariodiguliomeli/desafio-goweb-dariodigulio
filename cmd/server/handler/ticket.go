package handler

import (
	"desafio-go-web/internal/tickets"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	Service domain.TicketService
}

func NewService(s Service) *Service {
	return &Service{
		Service: s.Service,
	}
}

func (s *Service) GetTicketsByCountry() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		tickets, err := s.Service.GetTotalTickets(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, tickets)
	}
}

func (s *Service) AverageDestination() gin.HandlerFunc {
	return func(c *gin.Context) {
		destination := c.Param("dest")
		avg, err := s.Service.AverageDestination(destination)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error(), nil)
			return
		}
		c.JSON(200, avg)
	}
}
