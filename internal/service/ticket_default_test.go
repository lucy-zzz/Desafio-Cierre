package service_test

import (
	"app/internal/repository"
	"app/internal/service"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {
		rp := repository.NewRepositoryTicketMock()
		// rp.Get = func() (ta map[int]internal.TicketAttributes, err error) {
		// 	ta = map[int]internal.TicketAttributes{
		// 		1: {
		// 			Name:    "John",
		// 			Email:   "johndoe@gmail.com",
		// 			Country: "USA",
		// 			Hour:    "10:00",
		// 			Price:   100,
		// 		},
		// 	}
		// 	return ta, err
		// }

		sv := service.NewServiceTicketDefault(rp)

		total, err := sv.GetTotalTickets()

		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
