package dtos

import "github.com/mubarik-siraji/booking-system/models"

type CreateGuestDto struct {
	EmailAddress string      `json:"emailAddress" binding:"required,email"`
	Password     string      `json:"password" binding:"required,min=8,max=128"`
	FullName     string      `json:"fullname" binding:"required"`
	Role         models.Role `json:"role" binding:"oneo=ADMIN STUDENT CASHIER STUDENT_AFFEIRS"`
}

type LoginGuestRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8,max=128"`
}

type LoginGuestResponse struct {
	AccessToken  string
	RefreshToken string
	Guest        models.Guest
}
