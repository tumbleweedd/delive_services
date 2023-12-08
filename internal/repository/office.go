package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
)

type OfficeRepository struct {
	db *sqlx.DB
}

func NewOfficeRepository(db *sqlx.DB) *OfficeRepository {
	return &OfficeRepository{
		db: db,
	}
}

func (officeRepo *OfficeRepository) GetOffice(officeUUID uuid.UUID) (*models.Office, error) {
	//TODO implement me
	panic("implement me")
}

func (officeRepo *OfficeRepository) CreateOffice(uuid uuid.UUID, office *models.Office) error {
	//TODO implement me
	panic("implement me")
}

func (officeRepo *OfficeRepository) GetOfficeList() ([]*models.Office, error) {
	//TODO implement me
	panic("implement me")
}

func (officeRepo *OfficeRepository) UpdateOffice(office *models.Office) error {
	//TODO implement me
	panic("implement me")
}

func (officeRepo *OfficeRepository) DeleteOffice(officeUUID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
