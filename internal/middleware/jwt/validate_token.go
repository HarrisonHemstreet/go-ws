package jwt

import (
	"context"
	"net/http"
	"strings"

	"github.com/HarrisonHemstreet/go-ws/internal/model"
	jwt_utils "github.com/HarrisonHemstreet/go-ws/internal/utils/jwt"
	"github.com/golang-jwt/jwt/v5"
)

// MiddlewareValidateToken validates the JWT token.
func ValidateToken(next http.Handler, unprotectedMethods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if unprotected(unprotectedMethods, r.Method) {
			next.ServeHTTP(w, r) // Skip JWT validation
			return
		}

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			http.Error(w, "Malformed token", http.StatusUnauthorized)
			return
		}

		tokenString := authHeader[1]
		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwt_utils.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Invalid token signature", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Define a type for the context key to avoid collisions
		type contextKey string

		// Define a value for the context key
		const usernameKey contextKey = "username"

		// Token is valid, put the username in the context
		ctx := context.WithValue(r.Context(), usernameKey, claims.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func unprotected(unprotectedMethods []string, incomingMethod string) bool {
	for _, unprotectedMethod := range unprotectedMethods {
		if strings.EqualFold(unprotectedMethod, incomingMethod) {
			return true
		}
	}
	return false
}
