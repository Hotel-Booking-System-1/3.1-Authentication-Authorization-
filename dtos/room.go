package dtos

import (
	"github.com/mubarik-siraji/booking-system/models"

)

type CreateRoomRequest struct {
	RoomNumber    string            `json:"room_number" validate:"required"`
	RoomType      models.RoomType   `json:"room_type" validate:"required,oneof=Single Double Suite"`
	PricePerNight float64           `json:"price_per_night" validate:"required,gt=0"`
	Status        models.RoomStatus `json:"status" validate:"required,oneof=Available Maintenance Inactive"`
	Description   string            `json:"description"`
}

type UpdateRoomRequest struct {
	PricePerNight *float64           `json:"price_per_night" validate:"omitempty,gt=0"`
	Status        *models.RoomStatus `json:"status" validate:"omitempty,oneof=Available Maintenance Inactive"`
	Description   *string            `json:"description"`
}

type RoomResponse struct {
	ID            uint              `json:"id"`
	RoomNumber    string            `json:"room_number"`
	RoomType      models.RoomType   `json:"room_type"`
	PricePerNight float64           `json:"price_per_night"`
	Status        models.RoomStatus `json:"status"`
	Description   string            `json:"description"`
	CreatedAt     string            `json:"created_at"`
	UpdatedAt     string            `json:"updated_at"`
}

type RoomFilter struct {
	RoomNumber string `json:"room_number"`
	RoomType   string `json:"room_type"`
	Status     string `json:"status"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
}

type PaginatedRoomResponse struct {
	Data       []RoomResponse `json:"data"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}