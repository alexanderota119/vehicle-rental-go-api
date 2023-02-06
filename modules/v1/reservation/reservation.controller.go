package reservation

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type reservation_controller struct {
	service interfaces.ReservationServiceIF
}

func NewController(service interfaces.ReservationServiceIF) *reservation_controller {
	return &reservation_controller{service}

}

func (c *reservation_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.Reservation
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Add(&data)
	result.Send(w)

}
