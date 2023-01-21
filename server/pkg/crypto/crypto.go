package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// Password hashing functions adapted from https://gowebexamples.com/password-hashing/

func HashPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	hashedPassword = string(bytes)
	return
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
