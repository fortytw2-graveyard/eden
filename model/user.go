package model

import (
	"log"
	"os"
	"time"

	"github.com/fortytw2/abdi"
)

// ensure abdi has the filesystem key for HMAC
// this key is loaded into env from .env by github.com/joho/godotenv
func init() {
	abdi.Key = []byte(os.Getenv("SECRET_KEY"))
	// ensure we have a secret key
	if abdi.Key == nil {
		log.Fatalln("FATAL: secret key not present")
	}
}

// User model
type User struct {
	ID           int       `json:"-"`
	Username     string    `json:"username"`
	Email        string    `json:"-"`
	PasswordHash string    `json:"-"`
	Banned       bool      `json:"-"`
	Admin        bool      `json:"-"`
	Confirmed    bool      `json:"-"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// CreateUser creates a new, validated user
func CreateUser(username string, email string, password string) (user *User, err error) {
	var hash string
	hash, err = abdi.Hash(password)
	if err != nil {
		return
	}

	user = &User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
		Admin:        false,
		Confirmed:    false,
	}

	return
}

// CheckPassword checks a users password against the password hash and returns
// a bool and any errors
func (u *User) CheckPassword(password string) bool {
	if err := abdi.Check(password, u.PasswordHash); err != nil {
		return false
	}
	return true
}

// GenConfirmationCode creates a confirmationcode using crypto
func (u *User) GenConfirmationCode() (*string, error) {
	return nil, nil
}
