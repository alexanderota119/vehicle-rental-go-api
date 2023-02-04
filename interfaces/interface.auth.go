package interfaces

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"github.com/rfauzi44/vehicle-rental/lib"
)

type AuthServiceIF interface {
	Login(id model.User) *lib.Response
}
