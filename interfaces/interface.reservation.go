package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type ReservationRepoIF interface {
	Add(data *model.Reservation) (*model.Reservation, error)
}

type ReservationServiceIF interface {
	Add(data *model.Reservation) *lib.Response
}
