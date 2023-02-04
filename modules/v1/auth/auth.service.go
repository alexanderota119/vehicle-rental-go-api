package auth

import (
	"github.com/rfauzi44/vehicle-rental/interfaces"
)

type user_service struct {
	repo interfaces.UserRepoIF
}

func NewService(repo interfaces.UserRepoIF) *user_service {
	return &user_service{repo}

}
