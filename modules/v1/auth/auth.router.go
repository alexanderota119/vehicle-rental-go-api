package auth

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/modules/v1/user"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/auth").Subrouter()

	repo := user.NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/login", controller.Login).Methods("POST")
	route.HandleFunc("/register", controller.Register).Methods("POST")

}
