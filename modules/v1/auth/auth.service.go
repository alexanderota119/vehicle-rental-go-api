package auth

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/interfaces"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type auth_service struct {
	repo interfaces.UserRepoIF
}

type tokenResponse struct {
	Token string `json:token`
}

func NewService(repo interfaces.UserRepoIF) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(body *model.User) *lib.Response {
	user, err := s.repo.FindEmail(body.Email)
	if err != nil {
		return lib.NewRes("Email not registered", 401, true)
	}

	if lib.CheckPassword(body.Password, user.Password) {
		return lib.NewRes("Wrong password", 401, true)

	}
	jwt := lib.NewToken(user.UserID, user.Role)
	token, err := jwt.CreateToken()
	if err != nil {
		return lib.NewRes(err, 501, true)
	}
	return lib.NewRes(tokenResponse{Token: token}, 200, false)

}

func (s *auth_service) Register(body *model.User) *lib.Response {
	hashPassword, err := lib.HashPassword(body.Password)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	body.Password = hashPassword
	data, err := s.repo.Add(body)
	if err != nil {
		return lib.NewRes(err.Error(), 400, true)
	}
	return lib.NewRes(data, 200, false)

}
