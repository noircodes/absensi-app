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
	ID        uuid.UUID      `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email" gorm:"unique"`
	Password  string         `json:"password"`
	Role      Role           `json:"role" gorm:"default:ROLE_USER"`
	Status    UserStatus     `json:"status" gorm:"default:ACTIVE"`
	Phone     string         `json:"phone"`
	Code      string         `json:"code"`
	Image     string         `json:"image"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type CreateUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required, email"`
	Password string `validate:"required, min=6"`
	Phone    string
}

type LoginUserRequest struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Role   string `json:"role"`
	Image  string `json:"image"`
}
