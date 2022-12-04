package auth

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword is used to encrypt password before transport
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

//VerifyPassword checks the input password while verifying it with the password
func VerifyPassword(userPassword string, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	if err != nil {
		return err
	}
	return nil
}
