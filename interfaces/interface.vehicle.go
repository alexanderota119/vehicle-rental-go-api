package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type VehicleRepoIF interface {
	Add(data *model.Vehicle) (*model.Vehicle, error)
	GetAll() (*model.Vehicles, error)
	GetById(uuid string) (*model.Vehicle, error)
	Update(data *model.Vehicle) (*model.Vehicle, error)
	Delete(uuid string) error
}

type VahicleServiceIF interface {
	Add(data *model.Vehicle) *lib.Response
	GetAll() *lib.Response
	GetById(uuid string) *lib.Response
	Update(data *model.Vehicle) *lib.Response
	Delete(uuid string) *lib.Response
}
