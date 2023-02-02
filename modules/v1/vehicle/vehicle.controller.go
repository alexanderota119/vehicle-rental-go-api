package vehicle

import (
	"encoding/json"
	"net/http"

	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type vehicle_controller struct {
	repo *vehicle_repo
}

func NewController(repo *vehicle_repo) *vehicle_controller {
	return &vehicle_controller{repo}
}

func (c *vehicle_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.repo.GetAll()
	if err != nil {
		lib.ResponseError(w, http.StatusInternalServerError, err.Error())
	}
	lib.ResponseSuccess(w, http.StatusOK, response)

}

func (c *vehicle_controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	response, err := c.repo.Add(&data)
	if err != nil {
		lib.ResponseError(w, http.StatusInternalServerError, err.Error())
	}

	lib.ResponseSuccess(w, http.StatusOK, response)

}
