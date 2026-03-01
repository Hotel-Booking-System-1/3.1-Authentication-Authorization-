package helpers

import (
	"time"

	"github.com/mubarik-siraji/booking-system/infra"
	"github.com/mubarik-siraji/booking-system/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(role models.Role, sub string, expiresIn int64, isRefreshOken bool) (string, error) {
	config := infra.Configurations
	// jwtSecret := []byte(config.JWTSECRET)
	var jwtSecret []byte

	if isRefreshOken {
		jwtSecret = []byte(config.RefreshJwtToKenSecret)
	} else {
		jwtSecret = []byte(config.AccsessJwtToKenSecret)
	}

	claims := jwt.MapClaims{
		"sub":      sub,
		"nbf":            time.Now().Unix(),
		"exp":            expiresIn,
		"isRefreshToken": isRefreshOken,
		"role":           role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}