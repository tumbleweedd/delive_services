package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	custom_errors "github.com/tumbleweedd/delive_services/sso/internal/lib/errors"
)

type CustomerSaver interface {
	SaveUser(ctx context.Context, officeUUID uuid.UUID, name, lastname, email, pwd string) (userUUID uuid.UUID, err error)
}

type CustomerGetter interface {
	GetUser(ctx context.Context, email string) (*models.UserStruct, error)
	GetUserList(ctx context.Context, officeUUID uuid.UUID) ([]*models.UserStruct, error)
	IsAdmin(ctx context.Context, userUUID uuid.UUID) (bool, error)
}

type AppInfoGetter interface {
	AppInfo(ctx context.Context, appID int) (*models.App, error)
}

type Customer interface {
	CustomerGetter
	AppInfoGetter
	CustomerSaver
}

type CustomerRepository struct {
	db *sqlx.DB
}

func (customerRepo *CustomerRepository) AppInfo(ctx context.Context, appID int) (*models.App, error) {
	const op = "repository.customer.AppInfo"
	const query = "SELECT id, name secret FROM apps WHERE id = $1"

	stmt, err := customerRepo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var appInfo models.App
	if err = stmt.QueryRowContext(ctx, appID).Scan(&appInfo.ID, &appInfo.Secret); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, custom_errors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &appInfo, nil
}

func NewCustomerRepository(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (customerRepo *CustomerRepository) SaveUser(ctx context.Context, officeUUID uuid.UUID, name, lastname, email, pwd string) (uuid.UUID, error) {
	const op = "repository.customer.SaveUser"
	const query = "INSERT INTO users (name, lastname, email, office_uuid, password, password_confirm) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid"

	stmt, err := customerRepo.db.Prepare(query)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var userUUID uuid.UUID
	err = stmt.QueryRowContext(ctx, name, lastname, email, officeUUID, pwd).Scan(&userUUID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userUUID, nil
}

func (customerRepo *CustomerRepository) GetUser(ctx context.Context, email string) (*models.UserStruct, error) {
	const op = "repository.customer.GetUser"
	const query = "SELECT * FROM users WHERE email = $1"

	stmt, err := customerRepo.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var user *models.UserStruct
	if err = stmt.QueryRowContext(ctx, email).Scan(&user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, custom_errors.ErrUserNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (customerRepo *CustomerRepository) GetUserList(ctx context.Context, officeUUID uuid.UUID) ([]*models.UserStruct, error) {
	const op = "repository.customer.GetUserList"

	const query = "SELECT uuid, name, lastname, email, office_uuid, office_name, created_at FROM users WHERE office_uuid = $1"

	stmt, err := customerRepo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	//var users []*models.UserStruct

	// TODO: разобраться, как получить всех юзеров
	//if err = stmt.QueryContext(ctx, officeUUID).Scan(&users); err != nil {
	//	return nil, fmt.Errorf("%s: %w", op, err)
	//}
	panic("unimplemented")
}

func (customerRepo *CustomerRepository) IsAdmin(ctx context.Context, userUUID uuid.UUID) (bool, error) {
	const op = "repository.customer.IsAdmin"
	const query = "SELECT is_admin FROM users WHERE uuid = $1"

	stmt, err := customerRepo.db.Prepare(query)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var isAdmin bool
	if err = stmt.QueryRowContext(ctx, userUUID).Scan(&isAdmin); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, custom_errors.ErrUserNotFound)
		}
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}
