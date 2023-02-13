package model

import (
	"time"
)

type Reservation struct {
	ReservationID string    `gorm:"type:uuid; primaryKey; default:uuid_generate_v4()" json:"reservation_id,omitempty" valid:"-"`
	FVehicleID    string    `json:"vehicle_id" valid:"type(string)"`
	FUserID       string    `json:"user_id,omitempty" valid:"-"`
	StartDate     time.Time `json:"start_date" valid:"-"`
	EndDate       time.Time `json:"end_date" valid:"-"`
	Duration      int       `json:"duration,omitempty" valid:"-"`
	Quantity      int       `json:"quantity,omitempty" valid:"-"`
	Total         int       `json:"total,omitempty" valid:"-"`
	IsPaid        bool      `json:"is_paid,omitempty" valid:"-"`
	CreatedAt     time.Time `json:"created_at" valid:"-"`
	UpdatedAt     time.Time `json:"updated_at,omitempty" valid:"-" `
	VehicleDetail Vehicle   `gorm:"foreignKey:FVehicleID;references:VehicleID;" json:"vehicle_detail,omitempty" valid:"-"`
}

type Reservations []Reservation
