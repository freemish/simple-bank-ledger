package adapters

import (
	"github.com/freemish/simple-bank-ledger/entities"
	"github.com/freemish/simple-bank-ledger/persistence"
	"github.com/freemish/simple-bank-ledger/processes"
)

type InMemoryHandler struct {
	LoggedInCustomer *entities.Customer
	cacheStore       persistence.InMemoryCacheStore
}

func NewInMemoryHandler() *InMemoryHandler {
	return &InMemoryHandler{
		LoggedInCustomer: nil,
		cacheStore:       persistence.NewInMemoryCacheStore(),
	}
}

func (imh *InMemoryHandler) GetLoggedInCustomer() *entities.Customer {
	return imh.LoggedInCustomer
}

func (imh *InMemoryHandler) Login(username, password string) (*entities.Customer, error) {
	cust, err := processes.Login(username, password, imh.cacheStore)
	imh.LoggedInCustomer = cust
	return cust, err
}

func (imh *InMemoryHandler) Register(name, username, password string) (*entities.Customer, error) {
	return processes.CreateAccount(username, name, password, imh.cacheStore)
}

func (imh *InMemoryHandler) Logout(username string) error {
	imh.LoggedInCustomer = nil
	return nil
}

func (imh *InMemoryHandler) GetBalance(cust *entities.Customer) (float64, error) {
	return imh.cacheStore.SelectBalanceByUsername(cust.Username)
}

func (imh *InMemoryHandler) GetTransactions(cust *entities.Customer, args ...map[string]interface{}) ([]*entities.Transaction, error) {
	return imh.cacheStore.SelectTransactionsByUsername(cust.Username, args...)
}

func (imh *InMemoryHandler) RecordTransaction(cust *entities.Customer, is_debit bool, amount float64, voided_tx int) (*entities.Transaction, error) {
	return processes.RecordTransaction(cust, is_debit, amount, voided_tx, imh.cacheStore)
}
