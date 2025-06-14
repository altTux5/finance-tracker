package helpers

import (
	"time"

	"github.com/altTux5/expense-tracker/config"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(config.JWTConf.JWTSecret)

func GenerateJWT(userID string) (string, error) {
	// set expiration time to 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	// set claims - what the token contains
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": expirationTime,
		"iat": time.Now().Unix(),
	}

	// generate the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign the token with the secret key
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
