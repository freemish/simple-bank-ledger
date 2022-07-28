package adapters

import (
	"github.com/freemish/simple-bank-ledger/entities"
)

type IAuthProcessHandler interface {
	Login(username, password string) (*entities.Customer, error)
	Register(name, username, password string) (*entities.Customer, error)
	Logout(username string) error
}

type ITransactionProcessHandler interface {
	GetBalance(cust *entities.Customer) (float64, error)
	GetTransactions(cust *entities.Customer, args ...map[string]interface{}) ([]*entities.Transaction, error)
	RecordTransaction(cust *entities.Customer, is_debit bool, amount float64, voided_tx int) (*entities.Transaction, error)
}
