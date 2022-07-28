package adapters

import (
	"testing"

	"github.com/freemish/simple-bank-ledger/processes"
)

func TestRegisterAndLogin(t *testing.T) {
	username := "molly"
	password := "1234"
	imh := NewInMemoryHandler()

	// Successful register
	regCust, regErr := imh.Register("Molly", username, password)
	if regErr != nil {
		t.Error(regErr)
	}
	if regCust == nil {
		t.Error("Registered user object unexpectedly nil")
	}

	// Successful login
	loginCust, loginErr := imh.Login(username, password)
	if loginErr != nil {
		t.Error(loginErr)
	}
	if loginCust != regCust {
		t.Errorf(
			"expected: %v; got: %v; logged in? %v",
			regCust, loginCust, imh.GetLoggedInCustomer(),
		)
	}

	// Make sure logged-in customer is "cached"
	if imh.GetLoggedInCustomer() == nil {
		t.Error("Logged in customer is unexpectedly nil")
	}
	// Successful logout
	imh.Logout(loginCust.Username)
	if imh.GetLoggedInCustomer() != nil {
		t.Error("Failed to log out customer")
	}

	// Unsuccessful register
	regCust, regErr = imh.Register("Molly2", username, password)
	if regErr != processes.ErrCustomerAlreadyExists {
		t.Error(regErr)
	}
	if regCust != nil {
		t.Error("Duplicate registered user object unexpectedly non-nil")
	}

	// Unsuccessful login 1: bad username
	loginCust, loginErr = imh.Login(username+"2", password)
	if loginErr != processes.ErrCustomerDoesNotExist {
		t.Error(loginErr)
	}
	if loginCust != nil {
		t.Errorf("Bad login case one unexpectedly returned customer obj")
	}

	// Unsuccessful login 2: bad password
	loginCust, loginErr = imh.Login(username, password+"2")
	if loginErr != processes.ErrPasswordDoesNotMatch {
		t.Error(loginErr)
	}
	if loginCust != nil {
		t.Errorf("Bad login case two unexpectedly returned customer obj")
	}
}

func TestTransactions(t *testing.T) {
	username := "molly"
	password := "1234"
	imh := NewInMemoryHandler()
	cust, _ := imh.Register("Molly", username, password)

	// Test beginning balance is 0
	bal0, _ := imh.GetBalance(cust)
	if bal0 > 0 {
		t.Errorf("Beginning balance was non-zero: %v", bal0)
	}
	// Test number of transactions is 0
	txs0, _ := imh.GetTransactions(cust)
	if len(txs0) > 0 {
		t.Errorf("More than 0 transactions seeded: %v", txs0)
	}

	// Test balance and tx history affected by recording a transaction
	tx, _ := imh.RecordTransaction(cust, false, 1.01, 0)
	bal1, _ := imh.GetBalance(cust)
	if bal1 != tx.Amount {
		t.Errorf("Beginning balance was not equal to tx amount: %v != %v", tx.Amount, bal1)
	}
	txs1, _ := imh.GetTransactions(cust)
	if len(txs1) != 1 {
		t.Errorf("Length of txs not 1: %v", len(txs1))
	}

	// Test summation and history with second tx
	imh.RecordTransaction(cust, true, 1.0, 0)

	bal2, _ := imh.GetBalance(cust)
	if bal2 != 2.01 {
		t.Errorf("Second balance != 2.01: %v", bal2)
	}
	txs2, _ := imh.GetTransactions(cust)
	if len(txs2) != 2 {
		t.Errorf("Length of txs not 2: %v", len(txs2))
	}
}
