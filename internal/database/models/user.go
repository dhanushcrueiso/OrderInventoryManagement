package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"unique;not null"`
	Password string    `gorm:"not null"`
	Email    string    `gorm:"unique;not null"`
	Role     string    `gorm:"default:user"`
	Mobile   string
	Name     string
}

type Account struct {
	ID        uuid.UUID
	Username  string
	Password  string
	Email     string
	Role      string
	Mobile    string
	Name      string
	ExpiresAt time.Time
}
