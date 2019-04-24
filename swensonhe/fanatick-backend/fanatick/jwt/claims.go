package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims are the token claims.
type Claims struct {
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}
