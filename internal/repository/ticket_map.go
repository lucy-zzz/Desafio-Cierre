package repository

import (
	"app/internal"
	"context"
	"fmt"
)

func NewRepositoryTicketMap(dbFile map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     dbFile,
		lastId: lastId,
	}
}

type RepositoryTicketMap struct {
	db map[int]internal.TicketAttributes

	lastId int
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}

func (r *RepositoryTicketMap) GetPercentageTicketsByDestinationCountry(c string) (p float64, err error) {
	l := 0
	for _, i := range r.db {
		if c == i.Country {
			l++
		}
	}

	if l == 0 {
		return 0, fmt.Errorf("No such country was found.")
	}

	i := len(r.db)

	p = float64(l) / float64(i)

	return p * 100.00, err
}
