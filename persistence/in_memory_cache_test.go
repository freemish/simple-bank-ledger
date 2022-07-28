package persistence

import (
	"testing"

	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/processes"
)

func getInMemoryCacheStore() *InMemoryCacheStore {
	return NewInMemoryCacheStore()
}

func TestSelectCustomerByUsernameDoesNotExist(t *testing.T) {
	var imc = getInMemoryCacheStore()
	_, err := imc.SelectCustomerByUsername("doesnotexist")
	if err != processes.ErrCustomerDoesNotExist {
		t.Error(err)
	}
}

func TestSelectCustomerByUsername(t *testing.T) {
	var imc = getInMemoryCacheStore()
	cust, err := processes.CreateAccount("molly", "Molly", "1234", imc)
	if err != nil {
		t.Error(err)
	}
	if cust == nil {
		t.Error("Customer not returned from CreateAccount function")
		return
	}
	if cust.Username != "molly" {
		t.Errorf("Unknown error while creating a customer: %v", cust)
	}
	cust, err = imc.SelectCustomerByUsername("molly")
	if err != nil {
		t.Errorf("Failed to select customer by username: %v", err)
	}
	if cust.Username != "molly" {
		t.Errorf("Unknown error while selecting newly created customer: %v", cust)
	}
}

func TestInsertCustomerAlreadyExists(t *testing.T) {
	var imc = getInMemoryCacheStore()
	imc.customers["goopy"] = &entities.Customer{Username: "goopy"}
	err := imc.InsertCustomer(imc.customers["goopy"])
	if err != processes.ErrCustomerAlreadyExists {
		t.Errorf("Expected customer already exists error and got: %s", err)
	}
}

func TestInsertCustomer(t *testing.T) {
	var imc = getInMemoryCacheStore()
	err := imc.InsertCustomer(&entities.Customer{Username: "goopy"})
	if err != nil {
		t.Errorf("Expected nil error while inserting customer and got: %s", err)
	}
}

func TestCreateAccount(t *testing.T) {
	var imc = getInMemoryCacheStore()
	cust, err := processes.CreateAccount("moo", "Molly", "1234", imc)
	if err != nil {
		t.Errorf("Unexpected error while creating account: %v", err.Error())
		return
	}
	if cust == nil {
		t.Error("Customer should not be nil")
		return
	}
	if cust.Username != "moo" {
		t.Errorf("Unexpected customer username: %v", cust.Username)
		return
	}
}

func TestLogin(t *testing.T) {
	var imc = getInMemoryCacheStore()
	processes.CreateAccount("moo", "Molly", "1234", imc)
	cust, err := processes.Login("moo", "1234", imc)
	if err != nil {
		t.Errorf("Login should have passed, but errored out: %v", err)
	}
	if cust == nil {
		t.Error("Customer was nil from Login")
	} else if cust.Username != "moo" {
		t.Errorf("Logged in wrong user: %v", cust)
	}
}

func TestSelectTransactions(t *testing.T) {
	var imc = getInMemoryCacheStore()
	cust, _ := processes.CreateAccount("moo", "Molly", "1234", imc)
	cust2, _ := processes.CreateAccount("moo2", "Molly", "1234", imc)

	txs, err := imc.SelectTransactionsByUsername("moo")
	if err != nil {
		t.Errorf("Got unexpected error from selecting transactions: %v", err)
	}
	if len(txs) != 0 {
		t.Errorf("Did not expect to have any txs returned from selection! %v", txs)
	}

	err = imc.InsertTransaction(&entities.Transaction{
		Customer:    *cust2,
		Amount:      2.09,
		Description: "This is a test",
	})
	if err != nil {
		t.Errorf("Expected to have no problem inserting transaction: %v", err)
	}

	imc.InsertTransaction(&entities.Transaction{
		Customer:    *cust,
		Amount:      -1.09,
		Description: "This is a test",
	})
	txs, err = imc.SelectTransactionsByUsername("moo")
	if err != nil {
		t.Errorf("Got unexpected error from selecting transactions: %v", err)
	}
	if len(txs) != 1 {
		t.Errorf("Expected exactly one transaction in returned list! Got: %v", txs)
	}
	if txs[0].Amount != -1.09 {
		t.Errorf("Unexpected transaction value: %v", txs[0])
	}

	bal, err := imc.SelectBalanceByUsername(cust.Username)
	if err != nil {
		t.Errorf("Unexpected error from balance checking: %v", err)
	}
	if bal != -1.09 {
		t.Errorf("Unexpected balance; wanted -1.09, got %v", bal)
	}

	imc.InsertTransaction(&entities.Transaction{
		Customer:    *cust,
		Amount:      1.09,
		Description: "This is a second test",
	})
	bal, err = imc.SelectBalanceByUsername(cust.Username)
	if err != nil {
		t.Errorf("Unexpected error from balance checking: %v", err)
	}
	if bal != 0.0 {
		t.Errorf("Unexpected balance; wanted 0.0, got %v", bal)
	}
}
