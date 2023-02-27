package vehicle

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
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

	image := r.Context().Value("image").(string)
	data.Image = image

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(image)
		lib.NewRes(err.Error(), 400, true).Send(w)
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

func (c *vehicle_controller) GetById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["uuid"]
	c.service.GetById(params).Send(w)
}

func (c *vehicle_controller) Update(w http.ResponseWriter, r *http.Request) {

	var data model.Vehicle

	image := r.Context().Value("image").(string)
	data.Image = image

	err := schema.NewDecoder().Decode(&data, r.MultipartForm.Value)

	if err != nil {
		_ = os.Remove(image)
		lib.NewRes(err.Error(), 400, true).Send(w)
		return
	}

	_, err = govalidator.ValidateStruct(data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Update(&data)
	result.Send(w)

}

func (c *vehicle_controller) Delete(w http.ResponseWriter, r *http.Request) {

	var data model.Vehicle
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		lib.NewRes(err.Error(), 500, true).Send(w)
		return
	}

	result := c.service.Delete(data.VehicleID)
	result.Send(w)

}

func (c *vehicle_controller) GetByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["category"]
	c.service.GetByCategory(params).Send(w)
}

func (c *vehicle_controller) Sort(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	by := query["by"][0]
	order := query["order"][0]
	c.service.Sort(by, order).Send(w)
}

func (c *vehicle_controller) GetBySlug(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["slug"]
	c.service.GetBySlug(params).Send(w)
}

func (c *vehicle_controller) Search(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query()

	s := query["s"][0]

	c.service.Search(s).Send(w)

}
