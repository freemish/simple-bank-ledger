package adapters

import (
	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/persistence"
	"github.com/freemish/simple-bank-ledger/processes"
)

type InMemoryHandler struct {
	logged_in_customer *entities.Customer
	cache_store        persistence.InMemoryCacheStore
}

func NewInMemoryHandler() InMemoryHandler {
	return InMemoryHandler{
		logged_in_customer: nil,
		cache_store:        persistence.NewInMemoryCacheStore(),
	}
}

func (imh *InMemoryHandler) Login(username, password string) (*entities.Customer, error) {
	cust, err := processes.Login(username, password, imh.cache_store)
	imh.logged_in_customer = cust
	return cust, err
}

func (imh *InMemoryHandler) Register(name, username, password string) (*entities.Customer, error) {
	return processes.CreateAccount(username, name, password, imh.cache_store)
}

func (imh *InMemoryHandler) Logout(username string) error {
	imh.logged_in_customer = nil
	return nil
}

func (imh *InMemoryHandler) GetBalance(cust *entities.Customer) (float64, error) {
	return imh.cache_store.SelectBalanceByUsername(cust.Username)
}

func (imh *InMemoryHandler) GetTransactions(cust *entities.Customer, args ...map[string]interface{}) ([]*entities.Transaction, error) {
	return imh.cache_store.SelectTransactionsByUsername(cust.Username, args...)
}

func (imh *InMemoryHandler) RecordTransaction(cust *entities.Customer, is_debit bool, amount float64, voided_tx int) (*entities.Transaction, error) {
	return processes.RecordTransaction(cust, is_debit, amount, voided_tx, imh.cache_store)
}
