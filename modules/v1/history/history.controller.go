package history

import (
	"net/http"

	"github.com/rfauzi44/vehicle-rental/interfaces"
)

type history_controller struct {
	service interfaces.HistoryServiceIF
}

func NewController(service interfaces.HistoryServiceIF) *history_controller {
	return &history_controller{service}

}

func (c *history_controller) GetById(w http.ResponseWriter, r *http.Request) {

	UserID := r.Context().Value("user")
	result := c.service.GetById(UserID.(string))
	result.Send(w)

}
