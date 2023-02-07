package vehicle

import (
	"encoding/json"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type vehicle_controller struct {
	service interfaces.VahicleServiceIF
}

func NewController(service interfaces.VahicleServiceIF) *vehicle_controller {
	return &vehicle_controller{service}

}

func (c *vehicle_controller) Add(w http.ResponseWriter, r *http.Request) {

	var data model.Vehicle
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

func (c *vehicle_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	result := c.service.GetAll()
	result.Send(w)
}