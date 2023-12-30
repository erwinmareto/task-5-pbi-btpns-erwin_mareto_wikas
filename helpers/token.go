package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	Id uint `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(userId uint, email string) string {

	claims := &TokenClaims{
		Id: userId,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return err.Error()
	}
	return tokenString
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)

	if !ok || !token.Valid {
		return nil, err
	} 

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, err
	}

	return claims, nil
}
