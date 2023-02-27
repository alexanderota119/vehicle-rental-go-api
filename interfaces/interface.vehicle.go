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
	Delete(uuid string) (*model.Vehicle, error)
	GetByCategory(category string) (*model.Vehicles, error)
	Sort(by string, order string) (*model.Vehicles, error)
	GetBySlug(slug string) (*model.Vehicle, error)
	Search(query string) (*model.Vehicles, error)
}

type VahicleServiceIF interface {
	Add(data *model.Vehicle) *lib.Response
	GetAll() *lib.Response
	GetById(uuid string) *lib.Response
	Update(data *model.Vehicle) *lib.Response
	Delete(uuid string) *lib.Response
	GetByCategory(category string) *lib.Response
	Sort(by string, order string) *lib.Response
	GetBySlug(slug string) *lib.Response
	Search(query string) *lib.Response
}
