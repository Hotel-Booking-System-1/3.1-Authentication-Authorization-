package models

import (
	"time"

	"gorm.io/gorm"
)

// Booking represents a hotel booking
type Booking struct {
	gorm.Model
	GuestID      uint      `json:"guest_id"`
	Guest        *Guest    `json:"guest,omitempty"` // pointer to Guest for joins
	RoomID       uint      `json:"room_id"`
	Room         *Room     `json:"room,omitempty"`  // pointer to Room for joins
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `gorm:"default:'Reserved'" json:"status"`
}