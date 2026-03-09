package models

import (
	"time"

	"gorm.io/gorm"
)

// RoomType represents the category of the room
type RoomType string

const (
	Single RoomType = "Single"
	Double RoomType = "Double"
	Suite  RoomType = "Suite"
)

// RoomStatus represents the current state of the room
type RoomStatus string

const (
	Available   RoomStatus = "Available"
	Maintenance RoomStatus = "Maintenance"
	Inactive    RoomStatus = "Inactive"
)

// Room is the central entity for the Room Management module
type Room struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	RoomNumber    string         `gorm:"uniqueIndex;not null;size:50" json:"room_number"`
	RoomType      RoomType       `gorm:"type:varchar(20);not null" json:"room_type"`
	PricePerNight float64        `gorm:"type:decimal(10,2);not null" json:"price_per_night"`
	Status        RoomStatus     `gorm:"type:varchar(20);default:'Available'" json:"status"`
	Description   string         `gorm:"type:text" json:"description"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"` // Soft Delete field
}