package vehicle

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRoute(rt *mux.Router, db *gorm.DB) {
	route := rt.PathPrefix("/vehicle").Subrouter()
	repo := NewRepo(db)
	controller := NewController(repo)

	route.HandleFunc("/", controller.GetAll).Methods("GET")
	route.HandleFunc("/", controller.Add).Methods("POST")

}
