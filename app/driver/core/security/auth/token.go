package auth

import (
	os "os"
	time "time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	UserID string `json:"uid"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func jwtSecret() []byte {
	return []byte(os.Getenv("JWT_SECRET"))
}

func CreateAccessToken(userID string, ttl time.Duration) (string, error) {
	claims := UserClaims{
		UserID: userID,
		// Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			// Optional:
			// Issuer:  "my-api",
			// Subject: userID,
			// Audience: []string{"my-client"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret())
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret(), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrTokenInvalidClaims
}
