package repository

import (
	"github.com/mubarik-siraji/booking-system/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func RegisterUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) CreateUser(data models.User) error {
	return repo.DB.Create(&data).Error

}

func (repo *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := repo.DB.Where("email_address =?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *UserRepo) GetAllUsers(page, limit int, role, isActive string) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := repo.DB.Model(&models.User{})

	if role != "" {
		query = query.Where("role = ?", role)
	}

	if isActive != "" {
		isActiveBool := isActive == "true"
		query = query.Where("is_active = ?", isActiveBool)
	}

	// Count total records before pagination
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Fetch paginated results
	err = query.Limit(limit).Offset(offset).Order("created_at desc").Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (repo *UserRepo) GetUserByID(id uint) (models.User, error) {
	var user models.User
	err := repo.DB.First(&user, id).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}