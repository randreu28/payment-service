package jwt

import (
	"context"
	"errors"
	"net/http"
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

// Used to avoid collisions
type authContext string

const JwtPayloadKey authContext = "jwtPayload"

// AuthMiddleware validates JWT tokens for requests, attatchs the payload to the context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		authHeader := req.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(res, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		jwtPayload, err := ParseToken(authHeader)
		if err != nil {
			http.Error(res, "Invalid or expired JWT token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(req.Context(), JwtPayloadKey, jwtPayload)
		reqWithCtx := req.WithContext(ctx)

		next.ServeHTTP(res, reqWithCtx)
	})
}
