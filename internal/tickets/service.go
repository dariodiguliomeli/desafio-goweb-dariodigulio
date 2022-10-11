package domain

import (
	"desafio-go-web/internal/domain"
)

type TicketService struct {
	Repository Repository
}

func (s *TicketService) GetTotalTickets(destination string) ([]domain.Ticket, error) {
	all, err := s.Repository.GetAll()
	var ticketsByDestination []domain.Ticket
	for _, ticket := range all {
		if ticket.Country == destination {
			ticketsByDestination = append(ticketsByDestination, ticket)
		}
	}
	return ticketsByDestination, err
}

func (s *TicketService) AverageDestination(destination string) (float64, error) {
	tickets, err := s.GetTotalTickets(destination)
	totalPrice := 0.0
	for _, ticket := range tickets {
		totalPrice += ticket.Price
	}
	avg := totalPrice / float64(len(tickets))
	return avg, err
}
