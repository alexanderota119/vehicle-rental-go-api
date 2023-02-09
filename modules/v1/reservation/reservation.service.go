package reservation

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type reservation_service struct {
	repo interfaces.ReservationRepoIF
}

func NewService(repo interfaces.ReservationRepoIF) *reservation_service {
	return &reservation_service{repo}

}

func (s *reservation_service) Add(data *model.Reservation) *lib.Response {

	duration := data.EndDate.Sub(data.StartDate)
	data.Duration = int(duration.Hours() / 24)

	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *reservation_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *reservation_service) Update(data *model.Reservation) *lib.Response {
	duration := data.EndDate.Sub(data.StartDate)
	data.Duration = int(duration.Hours() / 24)
	data, err := s.repo.Update(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s *reservation_service) Delete(uuid string) *lib.Response {
	err := s.repo.Delete(uuid)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes("success", 200, false)

}
