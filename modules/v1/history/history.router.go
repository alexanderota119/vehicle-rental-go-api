package history

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/history").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/", middleware.Handle(controller.GetById, middleware.AuthMiddleware("user"))).Methods("GET")

}
