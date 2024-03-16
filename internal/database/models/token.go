package models

import (
	"time"

	"github.com/google/uuid"
)

type Token struct {
	Id         uuid.UUID `gorm:"PrimarKey"`
	Token      string
	AccountId  uuid.UUID
	Expires_At time.Time
}
