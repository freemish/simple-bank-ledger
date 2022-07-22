package persistence

import (
	"fmt"
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

func TestInsertCustomerAlreadyExists(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
	imc.customers["goopy"] = &entities.Customer{Username: "goopy"}
	err := imc.InsertCustomer(imc.customers["goopy"])
	if err != processes.ErrCustomerAlreadyExists {
		t.Errorf("Expected customer already exists error and got: %s", err)
	}
}

func TestCreateAccount(t *testing.T) {
	var imc = InMemoryCacheStore{customers: make(map[string]*entities.Customer)}
	imc.customers["goopy"] = &entities.Customer{Username: "Gooper"}
	fmt.Println(imc.customers)
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
