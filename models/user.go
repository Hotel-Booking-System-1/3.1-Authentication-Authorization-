package models

import "time"

type Role string

const (
	RoleAdmin        Role = "ADMIN"
	RoleReceptionist Role = "RECEPTIONIST"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	EmailAddress     string    `json:"emailAddress"`
	Password  string    `json:"password"`
	Role      Role      `json:"role"`
	Is_Active bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
