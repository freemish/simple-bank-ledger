package processes

import (
	"fmt"
	"testing"
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
)

var correctPassword = "1234"

// --- Mock helper functions

func mockCheckUsernameExistsReturnTrue(username string) bool {
	return true
}

func mockCheckUsernameExistsReturnFalse(username string) bool {
	return false
}

func mockInsertCustomer(c *entities.Customer) {}

func mockUpdateLastLogin(c *entities.Customer) {}

func mockSelectCustomerByUsernameCorrectPassword(username string) (*entities.Customer, error) {
	phash, _ := generatePasswordHash(correctPassword)
	return &entities.Customer{Username: username, PasswordHash: phash}, nil
}

func mockSelectCustomerByUsernameIncorrectPassword(username string) (*entities.Customer, error) {
	phash, _ := generatePasswordHash(username)
	return &entities.Customer{Username: username, PasswordHash: phash}, nil
}

func mockSelectCustomerByUsernameNotFound(username string) (*entities.Customer, error) {
	return nil, ErrCustomerDoesNotExist
}

// --- Tests

func TestCreateAccountNoPersistence(t *testing.T) {
	cust, err := CreateAccount("molly", "Molly", correctPassword, nil, nil)
	if err != nil {
		t.Error(err)
	}

	want := fmt.Sprintf(
		"ID: 0\tCreated: %s\tName: Molly\tUsername: molly\tAccount Number: %s",
		cust.Created.Format(time.RFC3339Nano),
		cust.AccountNumber,
	)
	got := cust.DebugString()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCreateAccountMockPersistenceAccountCreationFailure(t *testing.T) {
	_, err := CreateAccount("", "", "", mockCheckUsernameExistsReturnTrue, mockInsertCustomer)
	if err != ErrCustomerAlreadyExists {
		t.Errorf("Expected ErrCustomerAlreadyExists: %v", err)
	}
}

func TestCreateAccountMockPersistence(t *testing.T) {
	cust, err := CreateAccount("molly", "Molly", correctPassword, mockCheckUsernameExistsReturnFalse, mockInsertCustomer)
	if err != nil {
		t.Error(err)
	}

	want := fmt.Sprintf(
		"ID: 0\tCreated: %s\tName: Molly\tUsername: molly\tAccount Number: %s",
		cust.Created.Format(time.RFC3339Nano),
		cust.AccountNumber,
	)
	got := cust.DebugString()

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLoginNoPersistenceLoginFailure(t *testing.T) {
	_, err := Login("molly", correctPassword, nil, nil)
	if err != ErrCustomerDoesNotExist {
		t.Error(err)
	}
}

func TestLoginPass(t *testing.T) {
	cust, err := Login("molly", correctPassword, mockSelectCustomerByUsernameCorrectPassword, mockUpdateLastLogin)
	if err != nil {
		t.Error(err)
	}
	if !verifyLoginPassword(correctPassword, cust.PasswordHash) {
		t.Error("Incorrect password but login passed")
	}
}

func TestLoginFail(t *testing.T) {
	_, err := Login("molly", "wrongpassword", mockSelectCustomerByUsernameIncorrectPassword, mockUpdateLastLogin)
	if err != ErrPasswordDoesNotMatch {
		t.Error(err)
	}
}

func TestLoginFailUsernameNotFound(t *testing.T) {
	_, err := Login("golly", correctPassword, mockSelectCustomerByUsernameNotFound, mockUpdateLastLogin)
	if err != ErrCustomerDoesNotExist {
		t.Error(err)
	}
}

func TestPasswordHashNotEqualToPassword(t *testing.T) {
	pHash, err := generatePasswordHash(correctPassword)
	if err != nil {
		t.Error(err)
	}
	if pHash == correctPassword {
		t.Error("Password hash was found to be equal to original password!")
	}
}

func TestPasswordHashNotEqualToSubsequentPasswordHash(t *testing.T) {
	pHash1, err := generatePasswordHash(correctPassword)
	if err != nil {
		t.Error(err)
	}
	pHash2, err := generatePasswordHash(correctPassword)
	if err != nil {
		t.Error(err)
	}
	if pHash1 == pHash2 {
		t.Error("Password hashes are not salted!")
	}
}
