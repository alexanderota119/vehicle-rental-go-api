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

func (s *vehicle_service) GetById(uuid string) *lib.Response {
	data, err := s.repo.GetById(uuid)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *vehicle_service) Update(data *model.Vehicle) *lib.Response {
	data, err := s.repo.Update(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *vehicle_service) Delete(uuid string) *lib.Response {
	data, err := s.repo.Delete(uuid)

	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *vehicle_service) GetByCategory(category string) *lib.Response {
	data, err := s.repo.GetByCategory(category)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *vehicle_service) Sort(by string, order string) *lib.Response {
	data, err := s.repo.Sort(by, order)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}

func (s *vehicle_service) GetBySlug(slug string) *lib.Response {
	data, err := s.repo.GetBySlug(slug)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}

	return lib.NewRes(data, 200, false)
}
