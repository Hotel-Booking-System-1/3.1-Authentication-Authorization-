package models

import (
	"time"
)

type Guest struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FullName    string    `json:"full_name" gorm:"not null"`
	PhoneNumber string    `json:"phone_number" gorm:"not null"`
	Email       string    `json:"email" gorm:"not null;unique"`
	IDNumber    *string   `json:"id_number"`
	Address     string    `json:"address"`
	HasBooking  bool      `json:"has_booking"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
