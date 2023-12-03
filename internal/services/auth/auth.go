package auth

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"time"
)

type Auth interface {
	SaveUser(ctx context.Context, name, lastname, email, officeUUID, pwd, pwdConfirm string) (userUUID uuid.UUID, err error)
}

type AuthService struct {
	log            *slog.Logger
	authRepository Auth
	tokenTTL       time.Duration
}

func NewAuthService(
	log *slog.Logger,
	authRepository Auth,
	tokenTTL time.Duration,

) *AuthService {
	return &AuthService{
		log:            log,
		authRepository: authRepository,
		tokenTTL:       tokenTTL,
	}
}

func (authService *AuthService) Login(ctx context.Context, email string, password string, appID int32) (token string, err error) {
	//TODO implement me
	panic("implement me")
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
