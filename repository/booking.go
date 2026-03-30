package repository

import (
	"github.com/mubarik-siraji/booking-system/models"
	"gorm.io/gorm"
	"time"
)

type BookingRepository interface {
	Create(booking *models.Booking) error
	// Updated: Now accepts excludeID to prevent self-collision during updates
	IsRoomAvailable(roomID uint, checkIn, checkOut time.Time, excludeID uint) (bool, error)
	GetByID(id uint) (*models.Booking, error)
	GetAll(status string, guestName string, offset, limit int) ([]models.Booking, int64, error)
	Update(booking *models.Booking) error
	GetBookingsByDateRange(start, end time.Time) ([]models.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func RegisterBookingRepo(db *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

// 1. Create Booking
func (r *bookingRepository) Create(booking *models.Booking) error {
	return r.db.Create(booking).Error
}

// 2. Double Booking Prevention (Corrected Logic)
func (r *bookingRepository) IsRoomAvailable(roomID uint, checkIn, checkOut time.Time, excludeID uint) (bool, error) {
	var count int64
	
	// Rule: (NewCheckIn < ExistingCheckOut) AND (NewCheckOut > ExistingCheckIn)
	query := r.db.Model(&models.Booking{}).
		Where("room_id = ? AND status NOT IN (?)", roomID, []string{"Cancelled", "Checked-out"}).
		Where("check_in_date < ? AND check_out_date > ?", checkOut, checkIn)

	// CRITICAL FIX: If we are updating (excludeID > 0), ignore the current booking record
	if excludeID > 0 {
		query = query.Where("id != ?", excludeID)
	}

	err := query.Count(&count).Error
	return count == 0, err
}

// 3. Get By ID
func (r *bookingRepository) GetByID(id uint) (*models.Booking, error) {
	var booking models.Booking
	// Preload Guest and Room so the Service/Handler has all the data
	err := r.db.Preload("Guest").Preload("Room").First(&booking, id).Error
	return &booking, err
}

// 4. Read Bookings (with Pagination, Filter, and Search)
func (r *bookingRepository) GetAll(status string, guestName string, offset, limit int) ([]models.Booking, int64, error) {
	var bookings []models.Booking
	var total int64

	query := r.db.Model(&models.Booking{}).
		Joins("JOIN guests ON guests.id = bookings.guest_id")

	if status != "" {
		query = query.Where("bookings.status = ?", status)
	}

	if guestName != "" {
		// ILIKE is used for case-insensitive search in PostgreSQL
		query = query.Where("guests.fullname ILIKE ?", "%"+guestName+"%")
	}

	// Count total before applying offset/limit
	query.Count(&total)

	err := query.Preload("Guest").Preload("Room").
		Offset(offset).Limit(limit).
		Order("bookings.created_at DESC").
		Find(&bookings).Error

	return bookings, total, err
}

// 5. Update Booking
func (r *bookingRepository) Update(booking *models.Booking) error {
	// Save updates all fields including associations if configured
	return r.db.Save(booking).Error
}

// 6. Filter by Date Range
func (r *bookingRepository) GetBookingsByDateRange(start, end time.Time) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.db.Where("check_in_date >= ? AND check_in_date <= ?", start, end).
		Preload("Guest").
		Preload("Room").
		Find(&bookings).Error
	return bookings, err
}