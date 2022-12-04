package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type SignedDetails struct {
	ID       uuid.UUID
	Username string
	Role     string
	jwt.StandardClaims
}

var JWT_SECRET_KEY string = os.Getenv("JWT_SECRET_KEY")

func GenerateAllTokens(username string, id uuid.UUID) (signedToken string, signedRefreshToken string, err error) {
	claims := SignedDetails{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(JWT_SECRET_KEY))
	if err != nil {
		return "", "", err
	}

	return token, refreshToken, err
}
