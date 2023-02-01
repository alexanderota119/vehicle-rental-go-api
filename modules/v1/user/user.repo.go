package user

import (
	"github.com/rfauzi44/vehicle-rental/database/orm/model"
	"gorm.io/gorm"
)

type user_repo struct {
	database *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}

}

func (r *user_repo) GetAll() (*model.Users, error) {
	var data model.Users
	err := r.database.Model(model.User{}).Find(&data).Error
	if err != nil {
		return nil, err
	}

	return &data, nil

}

func (r *user_repo) Add(data *model.User) (*model.User, error) {
	err := r.database.Create(data).Error
	if err != nil {
		return nil, err
	}

	return data, nil

}
