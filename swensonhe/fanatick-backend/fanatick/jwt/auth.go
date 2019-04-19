package jwt

import (
	"time"

	"github.com/swensonhe/fanatick-backend/fanatick"

	"github.com/dgrijalva/jwt-go"
)

// TokenAuth generates JWT tokens.
type TokenAuth struct {
	Secret        string
	TokenDuration time.Duration
}

const (
	issuer = `fanatick`
)

// TokenAuth should implement the fanatick.TokenAuthenticator interface.
var _ fanatick.TokenAuthenticator = &TokenAuth{}

// TokenAuth should implement the fanatick.TokenGenerator interface.
var _ fanatick.TokenGenerator = &TokenAuth{}

// Generate generates an auth token.
func (auth *TokenAuth) Generate(user *fanatick.User) (string, error) {
	expireAt := time.Now().Add(auth.TokenDuration).Unix()

	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(auth.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Authenticate authenticates a token.
func (auth *TokenAuth) Authenticate(tokenString string) (*fanatick.User, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrorUnexpectedSigningMethod
		}
		return []byte(auth.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrorUnauthorized
	}

	return &fanatick.User{ID: claims.UserID}, nil
}
