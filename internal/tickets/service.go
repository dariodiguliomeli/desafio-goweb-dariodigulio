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

func (s *TicketService) AverageDestination(destination string) (interface{}, error) {
	return nil, nil
}
