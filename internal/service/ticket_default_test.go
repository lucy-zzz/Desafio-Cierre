package service_test

import (
	"app/internal"
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceTicketDefault_GetTotalTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		rp := repository.NewRepositoryTicketMock()
		rp.FuncGet = func() (ta map[int]internal.TicketAttributes, err error) {
			ta = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return ta, err
		}

		sv := service.NewServiceTicketDefault(rp)

		total, err := sv.GetTotalTickets()

		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestServiceTicketDefault_GetTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get total tickets amount by country", func(t *testing.T) {
		rp := repository.NewRepositoryTicketMock()
		rp.FuncGetTicketsByDestinationCountry = func(c string) (ta map[int]internal.TicketAttributes, err error) {
			ta = map[int]internal.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return ta, err
		}

		sv := service.NewServiceTicketDefault(rp)

		total, err := sv.GetTicketsAmountByDestinationCountry("")

		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
