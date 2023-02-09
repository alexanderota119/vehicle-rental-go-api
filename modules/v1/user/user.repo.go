package user

import (
	"errors"

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
	err := r.database.Find(&data).Error
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
	data.Password = ""

	return data, nil

}

func (r *user_repo) FindEmail(email string) (*model.User, error) {
	var data model.User

	result := r.database.First(&data, "email = ?", email)
	if result.Error != nil {
		return nil, errors.New("email not found")
	}

	return &data, nil
}

func (r *user_repo) GetById(uuid string) (*model.User, error) {

	var data model.User
	var field = []string{"user_id", "email"}

	result := r.database.Select(field).First(&data, "user_id = ?", uuid)

	if result.Error != nil {
		return nil, errors.New("fail")
	}

	return &data, nil
}

func (r *user_repo) Update(data *model.User) (*model.User, error) {
	err := r.database.Model(&model.User{}).Where("user_id = ?", data.UserID).Updates(data).Error

	if err != nil {
		return nil, err
	}
	data.Password = ""

	return data, nil
}

func (r *user_repo) Delete(uuid string) error {
	err := r.database.Where("user_id = ?", uuid).Delete(&model.User{}).Error

	if err != nil {
		return err
	}

	return nil

}
