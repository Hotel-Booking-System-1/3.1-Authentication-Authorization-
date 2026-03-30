package service

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/mubarik-siraji/booking-system/dtos"
	"github.com/mubarik-siraji/booking-system/models"
	"github.com/mubarik-siraji/booking-system/repository"
)

type BookingService interface {
	CreateBooking(req dtos.CreateBookingRequest) (*models.Booking, error)
	GetAllBookings(params dtos.BookingFilterParams) (dtos.PaginatedBookingResponse, error)
	UpdateBooking(id uint, req dtos.UpdateBookingRequest) (*models.Booking, error)
	UpdateStatus(id uint, status string) error
	CancelBooking(id uint) error
}

type bookingService struct {
	repo     repository.BookingRepository
	roomRepo repository.RoomRepository
}
// Gudaha service/booking_service.go

func RegisterBookingService(r repository.BookingRepository, rr repository.RoomRepository) BookingService {
    return &bookingService{
        repo:     r,  // for Bookings
        roomRepo: rr, //  for rooms (Room Price)
    }
}

func NewBookingService(r repository.BookingRepository, rr repository.RoomRepository) BookingService {
	return &bookingService{repo: r, roomRepo: rr}
}

// --------------------------------------
// Helper: Calculate nights between two dates
func calculateNights(checkIn, checkOut time.Time) int {
	nights := int(checkOut.Sub(checkIn).Hours() / 24)
	if checkOut.Sub(checkIn).Hours()/24 > float64(nights) {
		nights++
	}
	if nights < 1 {
		nights = 1
	}
	return nights
}

// --------------------------------------
// 1. Create Booking + Price Calculation + Validation
func (s *bookingService) CreateBooking(req dtos.CreateBookingRequest) (*models.Booking, error) {
	// Validate dates
	if !req.CheckOutDate.After(req.CheckInDate) {
		return nil, errors.New("check-out must be after check-in")
	}

	// Check room availability
	available, err := s.repo.IsRoomAvailable(req.RoomID, req.CheckInDate, req.CheckOutDate,0)
	if err != nil {
		return nil, fmt.Errorf("failed to check room availability: %w", err)
	}
	if !available {
		return nil, errors.New("selected room is not available for the chosen dates")
	}

	// Get room price
	room, err := s.roomRepo.GetByID(req.RoomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}

	nights := calculateNights(req.CheckInDate, req.CheckOutDate)
	booking := &models.Booking{
		GuestID:      req.GuestID,
		RoomID:       req.RoomID,
		CheckInDate:  req.CheckInDate,
		CheckOutDate: req.CheckOutDate,
		TotalPrice:   float64(nights) * room.PricePerNight,
		Status:       "Reserved",
	}

	if err := s.repo.Create(booking); err != nil {
		return nil, fmt.Errorf("failed to create booking: %w", err)
	}

	return booking, nil
}

// --------------------------------------
// 2. Read Bookings with Pagination & Filtering
func (s *bookingService) GetAllBookings(params dtos.BookingFilterParams) (dtos.PaginatedBookingResponse, error) {
	offset := (params.Page - 1) * params.Limit
	data, total, err := s.repo.GetAll(params.Status, params.GuestName, offset, params.Limit)
	if err != nil {
		return dtos.PaginatedBookingResponse{}, fmt.Errorf("failed to get bookings: %w", err)
	}

	// Map to DTOs
	responseData := make([]dtos.BookingResponse, 0, len(data))
	for _, b := range data {
		responseData = append(responseData, dtos.BookingResponse{
			ID:           b.ID,
			GuestName:    b.Guest.FullName,
			RoomNumber:   b.Room.RoomNumber,
			CheckInDate:  b.CheckInDate,
			CheckOutDate: b.CheckOutDate,
			TotalPrice:   b.TotalPrice,
			Status:       b.Status,
		})
	}

	return dtos.PaginatedBookingResponse{
		Data:     responseData,
		Total:    total,
		Page:     params.Page,
		LastPage: int(math.Ceil(float64(total) / float64(params.Limit))),
	}, nil
}
func (s *bookingService) UpdateBooking(id uint, req dtos.UpdateBookingRequest) (*models.Booking, error) {

	booking, err := s.repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get booking: %w", err)
	}

	newIn := booking.CheckInDate
	newOut := booking.CheckOutDate

	// Update if provided
	if req.CheckInDate != nil {
		newIn = *req.CheckInDate
	}
	if req.CheckOutDate != nil {
		newOut = *req.CheckOutDate
	}

	// Validate dates
	if !newOut.After(newIn) {
		return nil, errors.New("check-out must be after check-in")
	}

	// Always assign (THIS IS THE FIX)
	booking.CheckInDate = newIn
	booking.CheckOutDate = newOut

	// Check availability
	available, err := s.repo.IsRoomAvailable(booking.RoomID, newIn, newOut,id)
	if err != nil {
		return nil, fmt.Errorf("failed to check room availability: %w", err)
	}
	if !available {
		return nil, errors.New("selected room is not available for the new dates")
	}

	// Recalculate price
	room, err := s.roomRepo.GetByID(booking.RoomID)
	if err != nil {
		return nil, fmt.Errorf("failed to get room: %w", err)
	}

	nights := calculateNights(newIn, newOut)
	booking.TotalPrice = float64(nights) * room.PricePerNight

	// Save
	if err := s.repo.Update(booking); err != nil {
		return nil, fmt.Errorf("failed to update booking: %w", err)
	}

	return booking, nil
}
// --------------------------------------
// 3. Update Booking with Re-validation

func (s *bookingService) UpdateStatus(id uint, status string) error {
	booking, err := s.repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to get booking: %w", err)
	}

	// Optional: validate status transitions
	validStatuses := map[string]bool{
		"Reserved":   true,
		"Checked-in": true,
		"Checked-out": true,
		"Cancelled":  true,
	}

	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	// Example business rule (optional)
	if booking.Status == "Cancelled" {
		return errors.New("booking already cancelled")
	}

	// Update status
	booking.Status = status

	// Save to DB
	if err := s.repo.Update(booking); err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}

	return nil
}
// --------------------------------------
// 5. Cancel Booking
func (s *bookingService) CancelBooking(id uint) error {
	return s.UpdateStatus(id, "Cancelled")
}
