package auth_service

import (
	"context"
	"errors"
	"fmt"
	custom_errors "github.com/tumbleweedd/delive_services/sso/internal/lib/errors"
	"github.com/tumbleweedd/delive_services/sso/internal/lib/jwt"
	"github.com/tumbleweedd/delive_services/sso/internal/lib/logger/sl"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"log/slog"
	"time"
)

type Auth interface {
	Login(ctx context.Context, email string, password string, appID int) (accessToken string, refreshToken string, err error)
	RegisterNewUser(ctx context.Context, name string, lastname string, email string, officeUUID string, pwd string, pwdConfirm string) (userUUID uuid.UUID, err error)
	IsAdmin(ctx context.Context, userUUID string) (bool, error)
	Logout(ctx context.Context, token string) (success bool) // TODO: непонятно пока, как реализовывать
	RefreshToken(ctx context.Context, userUUID string, refreshToken string) (aToken string, rToken string, err error)
}

type AuthService struct {
	log      *slog.Logger
	tokenTTL time.Duration

	customerRepository repository.Customer
}

func NewAuthService(log *slog.Logger, customerRepository repository.Customer, tokenTTL time.Duration) *AuthService {
	return &AuthService{
		log:                log,
		customerRepository: customerRepository,
		tokenTTL:           tokenTTL,
	}
}

func (authService *AuthService) Login(ctx context.Context, email string, password string, appID int) (string, string, error) {
	const op = "services.auth.Login"

	log := authService.log.With(slog.String("op", op), slog.String("username", email))

	user, err := authService.customerRepository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			log.Error("user not found", sl.Err(err))
			return "", "", fmt.Errorf("%s: %w", op, custom_errors.ErrInvalidCredentials)
		}
		log.Error("failed to get user by email: %v", sl.Err(err))
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	if err = authService.comparePasswordAndHash(password, user.HashPassword); err != nil {
		log.Error("failed to check password hash: %v", sl.Err(err))
		return "", "", fmt.Errorf("%s: %w", op, custom_errors.ErrInvalidCredentials)
	}

	// getting app that user interacted with
	app, err := authService.customerRepository.AppInfo(ctx, appID)
	if err != nil {
		log.Error("failed to get app info: %v", sl.Err(err))
		return "", "", fmt.Errorf("%s: %w", op, err)
	}

	access, refresh, err := jwt.NewToken(user, app, authService.tokenTTL)
	if err != nil {
		log.Error("failed to create token: %v", sl.Err(err))
		return "", "", fmt.Errorf("%s: %w", op, err)
	}
	return access, refresh, nil
}

func (authService *AuthService) comparePasswordAndHash(password, hash string) error {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (authService *AuthService) RegisterNewUser(ctx context.Context, name string, lastname string, email string, officeUUID string, pwd string, pwdConfirm string) (userUUID uuid.UUID, err error) {
	const op = "services.auth.RegisterNewUser"

	log := authService.log.With(slog.String("op", op), slog.String("email", email))

	if pwd != pwdConfirm {
		log.Error("passwords do not match", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	hashPass, err := authService.hashPassword(pwd)
	if err != nil {
		log.Error("failed to generate password hash", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	parsedOfficeUUID, err := uuid.Parse(officeUUID)
	if err != nil {
		log.Error("failed to parse office uuid", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	userUUID, err = authService.customerRepository.SaveUser(ctx, parsedOfficeUUID, name, lastname, email, hashPass)
	if err != nil {
		log.Error("failed to save user", sl.Err(err))
		return uuid.Nil, fmt.Errorf("%s: %w", op, err)
	}

	return userUUID, nil
}

func (authService *AuthService) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func (authService *AuthService) IsAdmin(ctx context.Context, userUUID string) (bool, error) {
	const op = "services.auth.IsAdmin"

	log := authService.log.With(slog.String("op", op), slog.String("userGUID", userUUID))

	userUUIDParsed, err := uuid.Parse(userUUID)
	if err != nil {
		log.Error("failed to parse user uuid", sl.Err(err))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	isAdmin, err := authService.customerRepository.IsAdmin(ctx, userUUIDParsed)
	if err != nil {
		log.Error("failed to check if user is admin", sl.Err(err))
		return false, fmt.Errorf("%s: %w", op, err)
	}

	return isAdmin, nil
}

func (authService *AuthService) Logout(ctx context.Context, token string) (success bool) {
	//TODO implement me
	panic("implement me")
}

func (authService *AuthService) RefreshToken(ctx context.Context, userUUID string, refreshToken string) (aToken string, rToken string, err error) {
	//TODO implement me
	panic("implement me")
}
