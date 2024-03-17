package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	CustomerID      uuid.UUID `gorm:"primaryKey"`
	ID              uuid.UUID `gorm:"primaryKey"`
	ProductID       uuid.UUID `gorm:"primaryKey"`
	QuantityOrdered int
	TotalPrice      float64
	OrderId         string
	OrderDate       time.Time
}
