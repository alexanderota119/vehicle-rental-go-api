package auth

import "github.com/rfauzi44/vehicle-rental/interfaces"

type auth_controller struct {
	service interfaces.AuthServiceIF
}

func NewController(service interfaces.AuthServiceIF) *auth_controller {
	return &auth_controller{service}

}
