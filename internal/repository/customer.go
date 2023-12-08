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

type CustomerRepository struct {
	db *sqlx.DB
}

func (customerRepo *CustomerRepository) AppInfo(ctx context.Context, appID int) (*models.App, error) {
	const op = "repository.auth.AppInfo"
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
	const op = "repository.auth.SaveUser"
	const query = "INSERT INTO users (name, lastname, email, office_uuid, password, password_confirm) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uuid"

	stmt, err := customerRepo.db.Prepare(query)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}
	defer stmt.Close()

	var userUUID uuid.UUID
	err = stmt.QueryRowContext(ctx, name, lastname, email, officeUUID, pwd).Scan(&userUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return uuid.Nil, fmt.Errorf("%s: %w", op, custom_errors.ErrUserNotFound)
		}
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userUUID, nil
}

func (customerRepo *CustomerRepository) GetUser(ctx context.Context, email string) (*models.UserStruct, error) {
	const op = "repository.auth.GetUser"
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

func (customerRepo *CustomerRepository) IsAdmin(ctx context.Context, userUUID uuid.UUID) (bool, error) {
	const op = "repository.auth.IsAdmin"
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
