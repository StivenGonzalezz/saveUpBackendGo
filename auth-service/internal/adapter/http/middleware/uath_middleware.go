package middleware

import (
	"auth-service/pkg/jwt"
)

func AuthMiddleware(tokenString string) (bool, error) {
	return jwt.ValidateToken(tokenString)
}
