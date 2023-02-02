package model

import (
	"time"
)

type User struct {
	ID          string     `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"user_id,omitempty"`
	FullName    string     `json:"full_name"`
	Email       string     `json:"email"`
	Password    string     `json:"password,omitempty"`
	Gender      string     `json:"gender"`
	Address     string     `json:"address,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty"`
	PhoneNumber string     `json:"phone_number"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty" `
}

type Users []User
