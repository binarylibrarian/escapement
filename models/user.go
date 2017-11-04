package models

import "golang.org/x/crypto/bcrypt"
import "log"

// User represents a registered user of the service
type User struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
}

// SetPassword uses the users salt to hash and store their password
func (u *User) SetPassword(password string) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	u.PasswordHash = hash
}

// ValidatePassword check the given password against the users password hash
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password))
	if err != nil {
		return false
	}
	return true
}
