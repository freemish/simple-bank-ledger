package persistence

import (
	"testing"

	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/processes"
)

func TestSelectCustomerByUsernameDoesNotExist(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
	_, err := imc.SelectCustomerByUsername("doesnotexist")
	if err != processes.ErrCustomerDoesNotExist {
		t.Error(err)
	}
}

func TestSelectCustomerByUsername(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
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
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
	imc.customers["goopy"] = &entities.Customer{Username: "goopy"}
	err := imc.InsertCustomer(imc.customers["goopy"])
	if err != processes.ErrCustomerAlreadyExists {
		t.Errorf("Expected customer already exists error and got: %s", err)
	}
}

func TestInsertCustomer(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
	err := imc.InsertCustomer(&entities.Customer{Username: "goopy"})
	if err != nil {
		t.Errorf("Expected nil error while inserting customer and got: %s", err)
	}
}

func TestCreateAccount(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
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
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
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
