package repository

import (
	"errors"
	"github.com/mubarik-siraji/booking-system/models"
	"github.com/mubarik-siraji/booking-system/dtos"
	"gorm.io/gorm"
)

type RoomRepository interface {
	Create(room *models.Room) error
	GetAll(filter dtos.RoomFilter) ([]models.Room, int64, error)
	GetByID(id uint) (*models.Room, error)
	Update(room *models.Room) error
	Delete(id uint) error
}

type roomRepository struct {
	DB *gorm.DB
}
func RegisterRoomRepo(db *gorm.DB) *roomRepository {
	return &roomRepository{
		DB: db,
	}
}
func NewRoomRepository(db *gorm.DB) RoomRepository {
	return &roomRepository{DB: db}
}

func (r *roomRepository) Create(room *models.Room) error {
	return r.DB.Create(room).Error
}

func (r *roomRepository) GetAll(f dtos.RoomFilter) ([]models.Room, int64, error) {
	var rooms []models.Room
	var total int64

	query := r.DB.Model(&models.Room{})

	// Filtering Logic
	if f.Status != "" {
		query = query.Where("status = ?", f.Status)
	}
	if f.RoomType != "" {
		query = query.Where("room_type = ?", f.RoomType)
	}
	if f.RoomNumber != "" {
		query = query.Where("room_number LIKE ?", "%"+f.RoomNumber+"%")
	}

	// Count total records before pagination
	query.Count(&total)

	// Pagination Logic
	offset := (f.Page - 1) * f.Limit
	err := query.Offset(offset).Limit(f.Limit).Order("room_number asc").Find(&rooms).Error

	return rooms, total, err
}

func (r *roomRepository) GetByID(id uint) (*models.Room, error) {
	var room models.Room
	err := r.DB.First(&room, id).Error
	return &room, err
}

func (r *roomRepository) Update(room *models.Room) error {
	// Updates the specific instance; GORM handles updated_at automatically
	return r.DB.Save(room).Error
}

func (r *roomRepository) Delete(id uint) error {
	// Business Rule: Check for active bookings before soft delete
	var activeBookings int64
	// Note: Replace "bookings" with your actual table name
	r.DB.Table("bookings").
		Where("room_id = ? AND status NOT IN ('Cancelled', 'CheckedOut')", id).
		Count(&activeBookings)

	if activeBookings > 0 {
		return errors.New("cannot delete room: active bookings exist")
	}

	// Performs soft delete because models.Room contains gorm.DeletedAt
	return r.DB.Delete(&models.Room{}, id).Error
}