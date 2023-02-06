package model

import (
	"time"
)

type Vehicle struct {
	VehicleID   string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"vehicle_id,omitempty" valid:"-"`
	Name        string    `json:"name" valid:"type(string)"`
	Location    string    `json:"location" valid:"type(string)"`
	Price       int       `json:"price" valid:"type(int)"`
	Description string    `json:"description,omitempty" valid:"-"`
	Category    string    `json:"category" valid:"type(string)"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" valid:"-"`
}

type Vehicles []Vehicle
