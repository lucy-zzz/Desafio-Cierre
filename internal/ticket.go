package internal

import "context"

type TicketAttributes struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Hour    string  `json:"hour"`
	Price   float64 `json:"price"`
}

type Ticket struct {
	Id         int              `json:"id"`
	Attributes TicketAttributes `json:"attributes"`
}

type RepositoryTicket interface {
	Get(ctx context.Context) (t map[int]TicketAttributes, err error)
	GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]TicketAttributes, err error)
	GetPercentageTicketsByDestinationCountry(c string) (p float64, err error)
}

type ServiceTicket interface {
	GetTicketsAmountByDestinationCountry(c string) (t int, err error)
	GetPercentageTicketsByDestinationCountry(c string) (p float64, err error)
}
