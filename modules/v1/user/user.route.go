package user

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/user").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/", middleware.Handle(controller.GetById, middleware.AuthMiddleware("user"))).Methods("GET")
	route.HandleFunc("/all", middleware.Handle(controller.GetAll, middleware.AuthMiddleware("admin"))).Methods("GET")

}
