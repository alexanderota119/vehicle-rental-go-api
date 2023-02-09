package reservation

import (
	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/middleware"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/reservation").Subrouter()

	repo := NewRepo(db)
	service := NewService(repo)
	controller := NewController(service)

	route.HandleFunc("/", middleware.Handle(controller.Add, middleware.AuthMiddleware("user"))).Methods("POST")

	route.HandleFunc("/all", middleware.Handle(controller.GetAll, middleware.AuthMiddleware("admin"))).Methods("GET")
	route.HandleFunc("/", middleware.Handle(controller.Update, middleware.AuthMiddleware("admin", "user"))).Methods("PUT")
	route.HandleFunc("/", middleware.Handle(controller.Delete, middleware.AuthMiddleware("admin"))).Methods("DELETE")

}
