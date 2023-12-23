package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tumbleweedd/delive_services/sso/internal/domain/models"
	"time"
)

// NewToken generates access and refresh tokens for a user.
//
// Parameters:
// - user: A pointer to a UserStruct representing the user.
// - app: A pointer to an App representing the app.
// - tokenTTL: A time.Duration representing the token's time to live.
// TODO: покрыть тестами данную функцию
func NewToken(user *models.UserStruct, app *models.App, tokenTTL time.Duration) (string, string, error) {
	// Create new access and refresh tokens
	accessToken, refreshToken := jwt.New(jwt.SigningMethodHS256), jwt.New(jwt.SigningMethodHS256)
	claims, rClaims := accessToken.Claims.(jwt.MapClaims), refreshToken.Claims.(jwt.MapClaims)

	// Set claims for access token
	claims["uuid"] = user.UUID
	claims["email"] = user.Email
	claims["app_id"] = app.ID
	claims["exp"] = time.Now().Add(tokenTTL).Unix()

	// Set claims for refresh token
	rClaims["uuid"] = user.UUID
	rClaims["email"] = user.Email
	rClaims["app_id"] = app.ID
	rClaims["refresh_exp"] = time.Now().Add(tokenTTL).AddDate(0, 0, 30).Unix()

	// TODO: перенесь секрет куда-то из модели
	// Generate signed access token string
	accessTokenStr, err := accessToken.SignedString([]byte(app.Secret))
	if err != nil {
		return "", "", fmt.Errorf("access_token generating error: %s", err)
	}

	// Generate signed refresh token string
	refreshTokenStr, err := refreshToken.SignedString([]byte(app.Secret))
	if err != nil {
		return "", "", fmt.Errorf("refresh_token generating error: %s", err)
	}

	return accessTokenStr, refreshTokenStr, nil
}
