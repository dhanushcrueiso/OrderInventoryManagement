package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
	Quantity    int
}

type Inventory struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	ProductID   uuid.UUID `gorm:"not null"`
	Quantity    int       `gorm:"not null"`
	LastUpdated time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type ProductUpdate struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Name        string
	Description string
	Price       float64
}
