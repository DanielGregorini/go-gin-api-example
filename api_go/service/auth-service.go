package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/DanielGregorini/go-api-gin/entity"
	"github.com/golang-jwt/jwt/v4"
)

type AuthService interface {
	GenerateToken(user entity.User) (string, error)
	ValidateToken(tokenStr string) (*jwt.Token, error)
	IsTokenValid(tokenStr string) bool
	IsIDInToken(tokenStr string, userID int) (bool, error)
}

type authService struct {
	secretKey string
}

func NewAuthService(secret string) AuthService {
	return &authService{secretKey: secret}
}

func (s *authService) GenerateToken(user entity.User) (string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   strconv.Itoa(user.ID),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *authService) ValidateToken(tokenStr string) (*jwt.Token, error) {
	// ParseWithClaims já faz o parsing direto em RegisteredClaims
	return jwt.ParseWithClaims(tokenStr, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verifica o método de assinatura
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})
}

func (s *authService) IsTokenValid(tokenStr string) bool {
	token, err := s.ValidateToken(tokenStr)
	if err != nil {
		return false
	}

	return token.Valid
}

func (s *authService) IsIDInToken(tokenStr string, userID int) (bool, error) {
	token, err := s.ValidateToken(tokenStr)
	if err != nil {
		return false, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok || !token.Valid {
		return false, fmt.Errorf("token inválido ou claims inesperadas")
	}

	return claims.Subject == strconv.Itoa(userID), nil
}
