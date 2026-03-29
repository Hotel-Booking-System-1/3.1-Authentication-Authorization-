package models

import "time"

const (
	RoleGuest Role = "GUEST"
)

type Guest struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	FullName     string     `json:"fullName"`
	EmailAddress string     `gorm:"uniqueIndex" json:"emailAddress"`
	Password     string     `json:"-"`           // password ma muuqato marka API response la diro
	Role         Role       `json:"role"`        // default: RoleGuest
	LastLoginAt  *time.Time `json:"lastLoginAt"` // nullable
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `gorm:"index" json:"deletedAt"` // soft delete
}
