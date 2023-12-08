package models

import "github.com/google/uuid"

type Order struct {
	OrderUUID uuid.UUID
	UserUUID  uuid.UUID
}

type OrderItem struct {
	ID          int
	Count       int
	ProductUUID uuid.UUID
	OrderUUID   uuid.UUID
}

type OrderItemByOffice struct {
	Count       int
	ProductUUID uuid.UUID
}

type OrdersByOffice struct {
	UserUUID      uuid.UUID
	OfficeUUID    uuid.UUID
	OfficeName    string
	OfficeAddress string
	Salads        []*OrderItemByOffice
	Garnishes     []*OrderItemByOffice
	Meats         []*OrderItemByOffice
	Soups         []*OrderItemByOffice
	Drinks        []*OrderItemByOffice
	Desserts      []*OrderItemByOffice
}
