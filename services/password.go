package services

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateByteString() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashAndSaltPassword(password string) (string, error) {
	// Generate a salt for the password

	randString, err := GenerateByteString()
	if err != nil {
		return "", err
	}
	salt, err := bcrypt.GenerateFromPassword(randString, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Generate the hashed password using the salt and the given password
	hashedPassword, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashedPasswordString := string(hashedPassword[:])

	return hashedPasswordString, nil
}
