package processes

import (
	"fmt"
	"testing"
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
)

func mockCheckUsernameExistsReturnTrue(username string) bool {
	return true
}

func mockCheckUsernameExistsReturnFalse(username string) bool {
	return false
}

func mockInsertCustomer(c *entities.Customer) {

}

func TestCreateAccountNoPersistence(t *testing.T) {
	cust, err := CreateAccount("molly", "Molly", "1234", nil, nil)
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
	cust, err := CreateAccount("molly", "Molly", "1234", mockCheckUsernameExistsReturnFalse, mockInsertCustomer)
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

func TestPasswordHashNotEqualToPassword(t *testing.T) {
	password := "1234"
	pHash, err := generatePasswordHash(password)
	if err != nil {
		t.Error(err)
	}
	if pHash == password {
		t.Error("Password hash was found to be equal to original password!")
	}
}

func TestPasswordHashNotEqualToSubsequentPasswordHash(t *testing.T) {
	password := "1234"
	pHash1, err := generatePasswordHash(password)
	if err != nil {
		t.Error(err)
	}
	pHash2, err := generatePasswordHash(password)
	if err != nil {
		t.Error(err)
	}
	if pHash1 == pHash2 {
		t.Error("Password hashes are not salted!")
	}
}
