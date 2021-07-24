package password

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	MIN_REQUIRED_PASSWORD_LENGTH     = 8
	MinCost                      int = 4  // the minimum allowable cost as passed in to GenerateFromPassword
	MaxCost                      int = 31 // the maximum allowable cost as passed in to GenerateFromPassword
	DefaultCost                  int = 10 // the cost that will actually be set if a cost below MinCost is passed into GenerateFromPassword
)

func CheckPasswordLength(password string) bool {
	passwordLength := len(password)

	if passwordLength < MIN_REQUIRED_PASSWORD_LENGTH {
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePasswords(hashed, password string) error {

	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return errors.New("Invalid Password")
	}

	return nil
}
