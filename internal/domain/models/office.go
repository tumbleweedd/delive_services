package models

import (
	"github.com/google/uuid"
	"time"
)

type Office struct {
	Uuid      uuid.UUID
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
