package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStruct struct {
	UUID         uuid.UUID
	Email        string
	Name         string
	LastName     string
	OfficeUUID   string
	HashPassword []byte
	LastLoginAt  time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	RefreshToken string
}
