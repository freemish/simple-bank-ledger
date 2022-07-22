package persistence

import (
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/processes"
)

type InMemoryCacheStore struct {
	customers map[string]*entities.Customer // username mapped to the customer object
}

func (imc InMemoryCacheStore) SelectCustomerByUsername(username string) (*entities.Customer, error) {
	customer, exists := imc.customers[username]
	if !exists {
		return nil, processes.ErrCustomerDoesNotExist
	}
	return customer, nil
}

func (imc InMemoryCacheStore) InsertCustomer(c *entities.Customer) error {
	_, err := imc.SelectCustomerByUsername(c.Username)
	if err == nil {
		return processes.ErrCustomerAlreadyExists
	}
	c.ID = len(imc.customers) + 1
	imc.customers[c.Username] = c
	return nil
}

func (imc InMemoryCacheStore) UpdateLastLogin(c *entities.Customer) error {
	c.LastLoginDate = time.Now()
	return nil
}
