package dtos

import "time"


type CreateBookingRequest struct {
	GuestID      uint      `json:"guest_id" binding:"required"`
	RoomID       uint      `json:"room_id" binding:"required"`
	CheckInDate  time.Time `json:"check_in_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
	CheckOutDate time.Time `json:"check_out_date" binding:"required" time_format:"2006-01-02T15:04:05Z07:00"`
}


type UpdateBookingRequest struct {
	RoomID       uint       `json:"room_id"`
	CheckInDate  *time.Time `json:"check_in_date" time_format:"2006-01-02T15:04:05Z07:00"`
	CheckOutDate *time.Time `json:"check_out_date" time_format:"2006-01-02T15:04:05Z07:00"`
}


type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=Reserved Checked-in Checked-out Cancelled"`
}


type BookingFilterParams struct {
	GuestName string `form:"guest_name"`
	Status    string `form:"status"`
	Page      int    `form:"page,default=1"`
	Limit     int    `form:"limit,default=10"`
}


type BookingResponse struct {
	ID           uint      `json:"id"`
	GuestName    string    `json:"guest_name"`
	RoomNumber   string    `json:"room_number"`
	CheckInDate  time.Time `json:"check_in_date"`
	CheckOutDate time.Time `json:"check_out_date"`
	TotalPrice   float64   `json:"total_price"`
	Status       string    `json:"status"`
}


type PaginatedBookingResponse struct {
	Data     []BookingResponse `json:"data"`
	Total    int64             `json:"total"`
	Page     int               `json:"page"`
	LastPage int               `json:"last_page"`
}