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
	data, err := s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
