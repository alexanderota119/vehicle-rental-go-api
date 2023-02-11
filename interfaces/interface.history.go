package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type HistoryRepoIF interface {
	GetById(uuid string) (*model.Reservations, error)
}

type HistoryServiceIF interface {
	GetById(uuid string) *lib.Response
}
