package processes

import (
	"github.com/freemish/simple-bank-ledger/entities"
)

type ICustomerStore interface {
	UpdateLastLogin(*entities.Customer) error
	InsertCustomer(*entities.Customer) error
	SelectCustomerByUsername(string) (*entities.Customer, error)
}

type ITransactionStore interface {
	SelectTransactionsByUsername(string, ...map[string]interface{}) ([]*entities.Transaction, error)
	InsertTransaction(*entities.Transaction) error
	SelectBalanceByUsername(string) (float64, error)
}
