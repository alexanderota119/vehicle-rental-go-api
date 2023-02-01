package user

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()
	repo := NewRepo(db)
	controller := NewController(repo)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")

}
