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

// CreateAccount creates an account if username does not exist.
// checkUsernameExists and insertAccount are both nullable.
func CreateAccount(
	username, name, password string,
	ics ICustomerStore,
) (*entities.Customer, error) {
	if ics != nil {
		cust, _ := ics.SelectCustomerByUsername(username)
		if cust != nil {
			return nil, ErrCustomerAlreadyExists
		}
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

	if ics != nil {
		return cust, ics.InsertCustomer(cust)
	}

	return cust, nil
}

// Login checks that the username exists, then checks that the password
// is correct for the customer.
func Login(
	username, password string,
	ics ICustomerStore,
) (*entities.Customer, error) {
	if ics == nil {
		return nil, ErrCustomerDoesNotExist
	}

	cust, _ := ics.SelectCustomerByUsername(username)
	if cust == nil {
		return cust, ErrCustomerDoesNotExist
	}

	if !verifyLoginPassword(password, cust.PasswordHash) {
		return nil, ErrPasswordDoesNotMatch
	}

	return cust, ics.UpdateLastLogin(cust)

}

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

// verifyLoginPassword returns true if password matches password hash.
func verifyLoginPassword(password, passwordHash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil
}
