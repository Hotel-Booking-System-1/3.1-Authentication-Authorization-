package repository

import (
	"apdirizakismail/baaackend01-dayone/models"

	"gorm.io/gorm"
)

type GuestRepo struct {
	DB *gorm.DB
}

func NewGuestRepo(db *gorm.DB) *GuestRepo {
	return &GuestRepo{DB: db}
}

func (r *GuestRepo) CreateGuest(data *models.Guest) error {
	return r.DB.Create(data).Error
}

func (r *GuestRepo) GetAllGuests() ([]models.Guest, error) {
	var guests []models.Guest
	err := r.DB.Find(&guests).Error
	return guests, err
}

func (r *GuestRepo) GetGuestByID(id uint) (models.Guest, error) {
	var guest models.Guest
	err := r.DB.First(&guest, id).Error
	return guest, err
}

func (r *GuestRepo) UpdateGuest(id uint, data *models.Guest) error {
	return r.DB.Model(&models.Guest{}).Where("id = ?", id).Updates(data).Error
}

func (r *GuestRepo) DeleteGuest(id uint) error {
	var guest models.Guest
	if err := r.DB.First(&guest, id).Error; err != nil {
		return err
	}

	if guest.HasBooking {
		return gorm.ErrInvalidData
	}
	return r.DB.Delete(&models.Guest{}, id).Error
}
