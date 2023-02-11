package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rfauzi44/vehicle-rental/database/orm"
	"github.com/rfauzi44/vehicle-rental/modules/v1/auth"
	"github.com/rfauzi44/vehicle-rental/modules/v1/history"
	"github.com/rfauzi44/vehicle-rental/modules/v1/reservation"
	"github.com/rfauzi44/vehicle-rental/modules/v1/user"
	"github.com/rfauzi44/vehicle-rental/modules/v1/vehicle"
)

func NewApp() (*mux.Router, error) {
	mainRoute := mux.NewRouter()

	db, err := orm.NewDB()
	if err != nil {
		return nil, err
	}

	var imageFolder = http.FileServer(http.Dir("./public/image"))
	mainRoute.PathPrefix("/public/").Handler(http.StripPrefix("/public/image", imageFolder))

	user.NewRoute(mainRoute, db)
	vehicle.NewRoute(mainRoute, db)
	auth.NewRoute(mainRoute, db)
	reservation.NewRoute(mainRoute, db)
	history.NewRoute(mainRoute, db)

	return mainRoute, nil

}
