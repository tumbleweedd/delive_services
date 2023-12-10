package models

import "github.com/google/uuid"

type App struct {
	ID   uuid.UUID
	Name string
	// чтобы подписывать токины и на стороне клиентсокого приложения валидировать их
	Secret string
}
