package bankledger

import (
	"errors"
	"time"

	"github.com/freemish/errgo"  // propagates stack traces for debugging/logging
	"golang.org/x/crypto/bcrypt" // for simplifying password hashing/salting
)

// Errors related to customers.
var (
	ErrCustomerAlreadyExists = errors.New("Customer already exists")
	ErrCustomerDoesNotExist  = errors.New("Customer does not exist")
	ErrPasswordDoesNotMatch  = errors.New("Password does not match")
)

// Customer defines a user of the bank ledger
// who can make transactions on a bank account.
type Customer struct {
	Name          string
	Username      string
	AccountNumber string

	// login stuff
	PasswordHash  string
	LastLoginDate time.Time
}

// CreateAccount accepts a username and password, hashes/salts the password,
// generates an account number, and stores the account.
func CreateAccount(name, username, password string) (*Customer, error) {
	// make sure account does not already exist
	if SelectCustomerByUsername(username) != nil {
		return nil, errgo.Wrap(ErrCustomerAlreadyExists)
	}

	// make sure password can be hashed/salted
	hash, err := generatePasswordHash(password)
	if err != nil {
		return nil, errgo.Wrap(err)
	}

	// build Customer type
	cust := &Customer{
		Username:      username,
		PasswordHash:  hash,
		AccountNumber: generateAccountNumber(),
		LastLoginDate: time.Now(),
		Name:          name,
	}

	// store Customer
	err = InsertCustomer(cust)
	if err != nil {
		return nil, errgo.Wrap(err)
	}

	return cust, nil
}

// Login checks that the username exists, then checks that the password
// is correct for the customer.
func Login(username, password string) (*Customer, error) {
	cust := SelectCustomerByUsername(username)

	if cust == nil {
		return cust, errgo.Wrap(ErrCustomerDoesNotExist)
	}

	if cust.verifyLoginPassword(password) {
		UpdateLastLogin(cust)
		return cust, nil
	}

	return nil, errgo.Wrap(ErrPasswordDoesNotMatch)
}

// verifyLoginPassword returns true if password belongs to the
// Customer, false otherwise.
func (cust *Customer) verifyLoginPassword(password string) bool {
	if cust == nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(cust.PasswordHash), []byte(password))
	if err != nil {
		return false
	}
	return true
}

// generateAccountNumber is supposed to generate a "unique"
// account number that is 10-12 digits long.
func generateAccountNumber() string {
	return ""
}

// generatePasswordHash uses the bcrypt library to generate a salted hash.
func generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	//return string(bytes), errgo.Wr
	if err != nil {
		return string(bytes), errgo.Wrap(err)
	}
	return string(bytes), nil
}
