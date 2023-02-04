package vehicle

import (
	"net/http"
)

type vehicle_controller struct {
	repo *vehicle_repo
}

func NewController(repo *vehicle_repo) *vehicle_controller {
	return &vehicle_controller{repo}
}

func (c *vehicle_controller) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (c *vehicle_controller) Add(w http.ResponseWriter, r *http.Request) {

}
