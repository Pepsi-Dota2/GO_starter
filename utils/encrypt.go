package utils

import "golang.org/x/crypto/bcrypt"

// CheckPasswordHash compares the hashed password with the plain password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
