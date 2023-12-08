package services

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

// TODO: создать табличку
// TODO: оформить методы репозитория (+ изменить интерфейсы его)
func (authService *AuthService) Login(ctx context.Context, email string, password string, appID int) (token string, err error) {
	const op = "services.auth.Login"

	log := authService.log.With(slog.String("op", op), slog.String("username", email))

	user, err := authService.customerRepository.GetUser(ctx, email)
	if err != nil {
		if errors.Is(err, custom_errors.ErrUserNotFound) {
			authService.log.Error("user not found", sl.Err(err))
			return "", fmt.Errorf("%s: %w", op, custom_errors.ErrInvalidCredentials)
		}
		authService.log.Error("failed to get user by email: %v", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err = authService.checkPasswordHash(password, user.HashPassword); err != nil {
		authService.log.Error("failed to check password hash: %v", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, custom_errors.ErrInvalidCredentials)
	}

	app, err := authService.customerRepository.AppInfo(ctx, appID)
	if err != nil {
		authService.log.Error("failed to get app info: %v", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	token, err := jwt.NewToken(user, app, authService.tokenTTL)
	if err != nil {
		authService.log.Error("failed to create token: %v", sl.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil
}

func (authService *AuthService) checkPasswordHash(password, hash string) error {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (authService *AuthService) RegisterNewUser(ctx context.Context, name string, lastname string, email string, officeUUID string, pwd string, pwdConfirm string) (userUUID uuid.UUID, err error) {
	//TODO implement me
	panic("implement me")
}

func (authService *AuthService) IsAdmin(ctx context.Context, userGUID string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (authService *AuthService) Logout(ctx context.Context, token string) (success bool) {
	//TODO implement me
	panic("implement me")
}

func (authService *AuthService) RefreshToken(ctx context.Context, userUUID string, refreshToken string) (aToken string, rToken string, err error) {
	//TODO implement me
	panic("implement me")
}
