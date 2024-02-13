package jwt

import (
	"time"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(username string) (string, error) {
	secretKey := JWTKey
	claims := model.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "your-application-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
