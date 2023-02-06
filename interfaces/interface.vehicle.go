package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type VehicleRepoIF interface {
	Add(data *model.Vehicle) (*model.Vehicle, error)
	GetAll() (*model.Vehicles, error)
}

type VahicleServiceIF interface {
	Add(data *model.Vehicle) *lib.Response
	GetAll() *lib.Response
}
