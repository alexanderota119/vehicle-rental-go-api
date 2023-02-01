package user

import (
	"encoding/json"
	"net/http"

	"github.com/rfauzi44/vehicle-rental/database/orm/model"
)

type user_controller struct {
	repo *user_repo
}

func NewController(repo *user_repo) *user_controller {
	return &user_controller{repo}
}

func (c *user_controller) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	response, err := c.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}

func (c *user_controller) Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var data model.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	response, err := c.repo.Add(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(response)

}
