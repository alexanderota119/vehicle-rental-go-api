package user

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type user_service struct {
	repo interfaces.UserRepoIF
}

func NewService(repo interfaces.UserRepoIF) *user_service {
	return &user_service{repo}

}

func (s user_service) GetAll() *lib.Response {
	data, err := s.repo.GetAll()
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}

func (s user_service) Add(data *model.User) *lib.Response {
	hashPassword, err := lib.HashPassword(data.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	data.Password = hashPassword
	data, err = s.repo.Add(data)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
