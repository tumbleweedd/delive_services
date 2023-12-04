package auth

import (
	"context"
	"github.com/tumbleweedd/delive_services/sso/internal/repository"
	"golang.org/x/crypto/bcrypt"

	"github.com/google/uuid"
	"log/slog"
	"time"
)

type AuthService struct {
	log            *slog.Logger
	authRepository repository.Auth
	tokenTTL       time.Duration
}

func NewAuthService(
	log *slog.Logger,
	authRepository repository.Auth,
	tokenTTL time.Duration,

) *AuthService {
	return &AuthService{
		log:            log,
		authRepository: authRepository,
		tokenTTL:       tokenTTL,
	}
}

// TODO: создать табличку
// TODO: оформить методы репозитория (+ изменить интерфейсы его)
func (authService *AuthService) Login(ctx context.Context, email string, password string, appID int32) (token string, err error) {
	//const op = "Auth.Login"
	//
	//user, err := authService.authRepository.GetByEmail(email)
	//if err != nil {
	//	if errors.Is(err, custom_errors.ErrUserNotFound) {
	//		authService.log.Warn("user not found", sl.Err(err))
	//		return "", fmt.Errorf("%s: %w", op, custom_errors.ErrInvalidCredentials)
	//	}
	//	authService.log.Error("failed to get userRepo by email: %v", err)
	//	return "", err
	//}
	//
	//if isPasswordValid := authService.checkPasswordHash(password, user.Password); !isPasswordValid {
	//	authService.log.Errorf()
	//}
	return "", err
}

func (authService *AuthService) checkPasswordHash(password, hash string) bool {
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(password))
	if err != nil {
		return false
	}
	return true
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
