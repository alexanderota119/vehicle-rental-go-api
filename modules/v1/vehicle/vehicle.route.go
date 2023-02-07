package vehicle

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/", middleware.Handle(controller.Add, middleware.AuthMiddleware("admin"))).Methods("POST")
	route.HandleFunc("/", controller.GetAll).Methods("GET")

}