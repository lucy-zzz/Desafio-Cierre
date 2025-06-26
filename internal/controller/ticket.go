package controller

import (
	"app/internal"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewTicketDefault(sv internal.ServiceTicket) *TicketDefault {
	return &TicketDefault{sv: sv}
}

type TicketDefault struct {
	sv internal.ServiceTicket
}

func (c *TicketDefault) GetTotalTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := c.sv.GetTotalTickets()

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "error",
				"status":  http.StatusInternalServerError,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "ok",
			"data":    data,
		})
	}
}

func (c *TicketDefault) GetByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")
		data, err := c.sv.GetTicketsAmountByDestinationCountry(dest)

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "error",
				"status":  http.StatusInternalServerError,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "ok",
			"data":    data,
		})
	}
}

func (c *TicketDefault) GetAverage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dest := chi.URLParam(r, "dest")

		data, err := c.sv.GetPercentageTicketsByDestinationCountry(dest)

		if err != nil {
			response.JSON(w, http.StatusInternalServerError, map[string]any{
				"message": "error",
				"status":  http.StatusInternalServerError,
			})
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "ok",
			"data":    data,
		})
	}
}
