package models

import (
	"github.com/google/uuid"
	"time"
)

type UserStruct struct {
	UUID         uuid.UUID `db:"uuid"`
	Email        string    `db:"email"`
	Name         string    `db:"name"`
	LastName     string    `db:"lastname"`
	OfficeUUID   string    `db:"office_uuid"`
	OfficeName   string    `db:"office_name"`
	HashPassword string    `db:"password"`
	LastLoginAt  time.Time `db:"last_login_at"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	RefreshToken string    `db:"refresh_token"`
}
