package model

import "time"

type User struct {
	UserId      string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"id,omitempty"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password,omitempty"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address,omitempty"`
	Birthday    time.Time `json:"birthday,omitempty"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdateAt    time.Time `json:"updated_at,omitempty" `
}

func (User) TableName() string {
	return "user"

}

type Users []User
