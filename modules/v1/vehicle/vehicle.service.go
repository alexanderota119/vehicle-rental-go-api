package vehicle

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type vehicle_service struct {
	repo interfaces.VehicleRepoIF
}

func NewService(repo interfaces.VehicleRepoIF) *vehicle_service {
	return &vehicle_service{repo}

}

func (s *vehicle_service) Add(data *model.Vehicle) *lib.Response {
	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *vehicle_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
