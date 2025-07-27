package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type JWTClaims struct {
	UserID     int      `json:"user_id"`
	Email      string   `json:"email"`
	Role       string   `json:"role"`
	Permission []string `json:"permission"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, permission []string, email, role string, secretKey string, expiresIn time.Duration) (string, error) {
	fmt.Println("EXPIRES at", jwt.NewNumericDate(time.Now().Add(expiresIn)))
	claims := JWTClaims{
		UserID:     userID,
		Email:      email,
		Role:       role,
		Permission: permission,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(tokenString string, secretKey string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		zap.L().Info("token data", zap.Any("claims", claims))
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
