package customer

import (
	"github.com/google/uuid"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
)

type CustomerRepository struct {
}

func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{}
}

func (customerRepo *CustomerRepository) GetUser(userUUID uuid.UUID) *models.UserStruct {
	return nil
}
