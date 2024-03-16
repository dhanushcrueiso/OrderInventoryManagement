package models

import "github.com/google/uuid"

type Customer struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"unique;not null"`
	Password string
	Email    string
	Mobile   string
	Name     string
}
