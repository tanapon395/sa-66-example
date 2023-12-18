package service

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.RegisteredClaims
}

// Generate Token generates a jwt token
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.ExpirationHours) * time.Hour)),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

// Validate Token validates the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return

}
