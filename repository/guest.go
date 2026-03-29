package repository

import (
	"github.com/mubarik-siraji/booking-system/models"
	"gorm.io/gorm"
)

type GuestRepo struct {
	DB *gorm.DB
}

// repository/guest.go
func RegisterGuestRepo(db *gorm.DB) *GuestRepo {
	return &GuestRepo{DB: db}
}

func (repo *GuestRepo) CreateGuest(data models.Guest) error {
	return repo.DB.Create(&data).Error
}

func (repo *GuestRepo) GetGuestByEmail(email string) (models.Guest, error) {
	var Guest models.Guest

	err := repo.DB.Where("email_address = ?", email).First(&Guest).Error
	if err != nil {
		return models.Guest{}, err
	}
	return Guest, nil
}
