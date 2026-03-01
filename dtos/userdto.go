package dtos

import "github.com/mubarik-siraji/booking-system/models"

type CreateUserDto struct {
	EmailAddress string      `json:"emailAddress" binding:"required,email"`
	Password     string      `json:"password" binding:"required,min=8,max=128"`
	FullName     string      `json:"fullName" binding:"required"`
	Role         models.Role `json:"role" binding:"required,oneof=admin  receptionist"`
}


type LoginUserRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8,max=128"`
}

type LoginUserResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

type UserResponse struct {
	ID        uint        `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	Role      models.Role `json:"role"`
	IsActive  bool        `json:"is_active"`
	CreatedAt string      `json:"created_at"`
}

type PaginatedUsersResponse struct {
	Items      []UserResponse `json:"items"`
	Total      int64          `json:"total"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalPages int            `json:"total_pages"`
}