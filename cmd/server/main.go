package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/domain"
	tickets "desafio-go-web/internal/tickets"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
)

func main() {
	list, err := LoadTicketsFromFile("../../tickets.csv")
	if err != nil {
		panic("Couldn't load tickets")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	ticketService := createTicketService(list)
	ticketsRouter := r.Group("tickets/")
	{
		ticketsRouter.GET("getByCountry/:destination", ticketService.GetTicketsByCountry())
		ticketsRouter.GET("getAverage/:destination", ticketService.AverageDestination())
	}
	// GET - “/ticket/getAverage/:dest”
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func createTicketService(list []domain.Ticket) *handler.Service {
	repository := tickets.NewRepository(list)
	ticketService := tickets.TicketService{Repository: repository}
	service := handler.NewService(handler.Service{Service: ticketService})
	return service
}

func LoadTicketsFromFile(path string) ([]domain.Ticket, error) {
	var ticketList []domain.Ticket
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	for _, row := range data {
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return []domain.Ticket{}, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
	}
	return ticketList, nil
}
