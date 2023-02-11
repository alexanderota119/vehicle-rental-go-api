package history

import (
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type history_service struct {
	repo interfaces.HistoryRepoIF
}

func NewService(repo interfaces.HistoryRepoIF) *history_service {
	return &history_service{repo}

}

func (s *history_service) GetById(uuid string) *lib.Response {
	data, err := s.repo.GetById(uuid)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
