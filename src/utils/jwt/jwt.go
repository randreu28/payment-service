package jwt

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	Id    int    `json:"id"`
	Owner string `json:"owner"`
}

type PaymentServiceJwt struct {
	Payload Payload `json:"payload"`
	jwt.RegisteredClaims
}

// CreateJwt creates and signs a new JWT token
func CreateJwt(id int, owner string) (string, error) {
	claims := PaymentServiceJwt{
		Payload: Payload{
			Id:    id,
			Owner: owner,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "payment-service",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 1 day
		},
	}

	unsignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := unsignedToken.SignedString([]byte(jwtSecret))

	return token, err
}

// ParseToken parses a new JWT token
func ParseToken(unparsedToken string) (*PaymentServiceJwt, error) {
	validateToken := func(token *jwt.Token) (interface{}, error) {
		jwtSecret := os.Getenv("JWT_SECRET")
		return []byte(jwtSecret), nil
	}

	token, err := jwt.ParseWithClaims(unparsedToken, &PaymentServiceJwt{}, validateToken)
	if err != nil {
		return nil, err
	}

	jwt, ok := token.Claims.(*PaymentServiceJwt)
	if !token.Valid || !ok {
		return nil, errors.New("invalid token")
	}

	return jwt, nil
}
