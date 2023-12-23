package models

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ProductUUID uuid.UUID
	Name        string
	Description string
	Type        ProductType
	Weight      int
	Price       float64
	CreatedAt   time.Time
}

type ProductType int

const (
	CustomerProductTypeUnspecified ProductType = iota
	CustomerProductTypeFirst
	CustomerProductTypeSecond
	CustomerProductTypeSalad
	CustomerProductTypeDessert
	CustomerProductTypeDrink
)
