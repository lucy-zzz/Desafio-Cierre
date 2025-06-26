package repository

import (
	"app/internal"
	"context"
)

func NewRepositoryTicketMock() *RepositoryTicketMock {
	return &RepositoryTicketMock{}
}

type RepositoryTicketMock struct {
	FuncGet                                      func() (t map[int]internal.TicketAttributes, err error)
	FuncGetTicketsByDestinationCountry           func(country string) (t map[int]internal.TicketAttributes, err error)
	FuncGetPercentageTicketsByDestinationCountry func(c string) (p float64, err error)

	Spy struct {
		Get                                      int
		GetTicketsByDestinationCountry           int
		GetPercentageTicketsByDestinationCountry int
	}
}

func (r *RepositoryTicketMock) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	r.Spy.Get++

	t, err = r.FuncGet()
	return
}

func (r *RepositoryTicketMock) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	r.Spy.GetTicketsByDestinationCountry++

	t, err = r.FuncGetTicketsByDestinationCountry(country)
	return
}

func (r *RepositoryTicketMock) GetPercentageTicketsByDestinationCountry(c string) (p float64, err error) {
	r.Spy.GetPercentageTicketsByDestinationCountry++

	p, err = r.FuncGetPercentageTicketsByDestinationCountry(c)
	return
}
