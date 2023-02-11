package model

import (
	"time"
)

type Reservation struct {
	ReservationID string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"reservation_id,omitempty" valid:"-"`
	VehicleID     string    `json:"vehicle_id" valid:"type(string)"`
	UserID        string    `json:"user_id,omitempty" valid:"-"`
	StartDate     time.Time `json:"start_date" valid:"-"`
	EndDate       time.Time `json:"end_date" valid:"-"`
	Duration      int       `json:"duration,omitempty" valid:"-"`
	Quantity      int       `json:"quantity,omitempty" valid:"-"`
	Total         int       `json:"total,omitempty" valid:"-"`
	IsPaid        bool      `gorm:"type:bool default:false" json:"is_paid,omitempty" valid:"-"`
	CreatedAt     time.Time `json:"created_at" valid:"-"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" valid:"-" `
	VehicleDetail Vehicle   `gorm:"foreignKey:VehicleID;references:VehicleID" json:"vehicle_detail,omitempty" valid:"-"`
}

type Reservations []Reservation
