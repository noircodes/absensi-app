package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "ROLE_ADMIN"
	RoleUser  Role = "ROLE_USER"
)

type UserStatus string

const (
	Active   UserStatus = "ACTIVE"
	Inactive UserStatus = "INACTIVE"
)

type User struct {
	ID         uuid.UUID  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name       string     `json:"name"`
	Email      string     `json:"email" gorm:"unique"`
	Password   string     `json:"password"`
	Role       Role       `json:"role" gorm:"default:ROLE_USER"`
	UserStatus UserStatus `json:"status" gorm:"default:ACTIVE"`
	Phone      string     `json:"phone"`
	Code       string     `json:"code"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type RegisterUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required, email"`
	Password string `validate:"required, min=6"`
	Phone    string `validate:"e164"`
}

type LoginUserRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
