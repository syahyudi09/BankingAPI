// auth.go

package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(customerID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("key-rahasia")


func (s *jwtService) GenerateToken(customerID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["customer_id"] = customerID
	claim["exp"] = time.Now().Add(time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		// Periksa metode tanda tangan token
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid Token")
		}

		// Kembalikan secret key sebagai interface{}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	// Periksa apakah token valid
	if !token.Valid {
		return nil, errors.New("Token is not valid")
	}

	return token, nil
}


func NewServiceJWT() Service {
	return &jwtService{}
}