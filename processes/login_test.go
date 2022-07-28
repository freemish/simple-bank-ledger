package processes

import (
	"fmt"
	"testing"
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
)

var correctPassword = "1234"

// --- Mock helpers

type MockCustomerStore struct {
	usernameExists bool
	password       string
}

func (mcs MockCustomerStore) SelectCustomerByUsername(username string) (*entities.Customer, error) {
	if !mcs.usernameExists {
		return nil, ErrCustomerDoesNotExist
	}

	pHash, _ := generatePasswordHash(mcs.password)
	return &entities.Customer{Username: username, PasswordHash: pHash}, nil
}

func (mcs MockCustomerStore) UpdateLastLogin(c *entities.Customer) error {
	return nil
}

func (mcs MockCustomerStore) InsertCustomer(c *entities.Customer) error {
	return nil
}

// --- Tests

func TestCreateAccountNoPersistence(t *testing.T) {
	cust, err := CreateAccount("molly", "Molly", correctPassword, nil)
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
	_, err := CreateAccount("", "", "", MockCustomerStore{usernameExists: true})
	if err != ErrCustomerAlreadyExists {
		t.Errorf("Expected ErrCustomerAlreadyExists: %v", err)
	}
}

func TestCreateAccountMockPersistence(t *testing.T) {
	cust, err := CreateAccount("molly", "Molly", correctPassword, MockCustomerStore{usernameExists: false})
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
	_, err := Login("molly", correctPassword, nil)
	if err != ErrCustomerDoesNotExist {
		t.Error(err)
	}
}

func TestLoginPass(t *testing.T) {
	cust, err := Login("molly", correctPassword, MockCustomerStore{usernameExists: true, password: correctPassword})
	if err != nil {
		t.Error(err)
	}
	if cust == nil {
		t.Error("Customer should not be nil")
		return
	}
	if !verifyLoginPassword(correctPassword, cust.PasswordHash) {
		t.Error("Incorrect password but login passed")
	}
}

func TestLoginFail(t *testing.T) {
	_, err := Login("molly", "wrongpassword", MockCustomerStore{usernameExists: true, password: correctPassword})
	if err != ErrPasswordDoesNotMatch {
		t.Error(err)
	}
}

func TestLoginFailUsernameNotFound(t *testing.T) {
	_, err := Login("golly", correctPassword, MockCustomerStore{usernameExists: false})
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
