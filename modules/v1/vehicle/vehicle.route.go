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
	route.HandleFunc("/all", controller.GetAll).Methods("GET")
	route.HandleFunc("/{uuid}", controller.GetById).Methods("GET")
	route.HandleFunc("/", middleware.Handle(controller.Update, middleware.AuthMiddleware("admin"))).Methods("PUT")
	route.HandleFunc("/", middleware.Handle(controller.Delete, middleware.AuthMiddleware("admin"))).Methods("DELETE")

}
