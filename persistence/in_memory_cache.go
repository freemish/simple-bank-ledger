package persistence

import (
	"errors"
	"time"

	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/processes"
)

var (
	ErrTransactionNotMappedToCustomer = errors.New("cannot store a transaction without a valid customer")
)

type InMemoryCacheStore struct {
	customers                 map[string]*entities.Customer // username mapped to the customer object
	transactions_by_customers map[string][]*entities.Transaction
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
	c.Created = time.Now()
	imc.customers[c.Username] = c
	return nil
}

func (imc InMemoryCacheStore) UpdateLastLogin(c *entities.Customer) error {
	c.LastLoginDate = time.Now()
	return nil
}

func (imc InMemoryCacheStore) SelectTransactionsByUsername(username string, args ...map[string]interface{}) ([]*entities.Transaction, error) {
	txs, exists := imc.transactions_by_customers[username]
	if !exists {
		imc.transactions_by_customers[username] = make([]*entities.Transaction, 0)
		return imc.transactions_by_customers[username], nil
	}
	return txs, nil
}

func (imc InMemoryCacheStore) InsertTransaction(tx *entities.Transaction) error {
	if tx.Customer.ID == 0 {
		return ErrTransactionNotMappedToCustomer
	}
	if _, err := imc.SelectCustomerByUsername(tx.Customer.Username); err != nil {
		return ErrTransactionNotMappedToCustomer
	}
	txs, _ := imc.SelectTransactionsByUsername(tx.Customer.Username)
	tx.ID = tx.Customer.ID*10 + len(txs) + 1
	tx.Created = time.Now()
	imc.transactions_by_customers[tx.Customer.Username] = append(txs, tx)
	return nil
}

func (imc InMemoryCacheStore) SelectBalanceByUsername(username string) (float64, error) {
	txs, _ := imc.SelectTransactionsByUsername(username)
	var balance float64 = 0.0
	for _, tx := range txs {
		balance = balance + tx.Amount
	}
	return balance, nil
}
