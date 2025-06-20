package service

import (
	"app/internal"
	"context"
)

type ServiceTicketDefault struct {
	rp internal.RepositoryTicket
}

func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

func (s *ServiceTicketDefault) GetTotalTickets() (total int, err error) {
	tickets, err := s.rp.Get(context.Background())

	total = len(tickets)
	return
}

func (s *ServiceTicketDefault) GetTicketsAmountByDestinationCountry(c string) (t int, err error) {
	l, err := s.rp.GetTicketByDestinationCountry(context.Background(), c)
	t = len(l)
	return
}

func (s *ServiceTicketDefault) GetPercentageTicketsByDestinationCountry(c string) (p float64, err error) {
	p, err = s.rp.GetPercentageTicketsByDestinationCountry(c)
	return p, err
}
