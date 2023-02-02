package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type UserRepoIF interface {
	GetAll() (*model.Users, error)
	Add(data *model.User) (*model.User, error)
}

type UserServiceIF interface {
	GetAll() *lib.Response
	Add(data *model.User) *lib.Response
}
