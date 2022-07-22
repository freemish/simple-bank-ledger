package processes

import (
	"errors"
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
	"golang.org/x/crypto/bcrypt"
)

// Errors related to customer login and account creation.
var (
	ErrCustomerAlreadyExists = errors.New("customer already exists")
	ErrCustomerDoesNotExist  = errors.New("customer does not exist")
	ErrPasswordDoesNotMatch  = errors.New("password does not match")
)

// generatePasswordHash uses the bcrypt library to generate a salted hash.
func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return string(bytes), err
	}
	return string(bytes), nil
}

func generateAccountNumber() string {
	accNum := "1234567890"
	return accNum
}

// CreateAccount creates an account if username does not exist.
func CreateAccount(
	username, name, password string,
	checkUsernameExists func(string) bool,
	insertAccount func(*entities.Customer),
) (*entities.Customer, error) {
	if checkUsernameExists != nil && checkUsernameExists(username) {
		return nil, ErrCustomerAlreadyExists
	}

	passHash, err := generatePasswordHash(password)
	if err != nil {
		return nil, err
	}

	cust := &entities.Customer{
		Created:       time.Now(),
		Name:          name,
		Username:      username,
		PasswordHash:  passHash,
		AccountNumber: generateAccountNumber(),
	}

	if insertAccount != nil {
		insertAccount(cust)
	}

	return cust, nil
}
