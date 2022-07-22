package processes

import (
	"github.com/freemish/simple-bank-ledger/entities"
)

type ICustomerStore interface {
	UpdateLastLogin(*entities.Customer) error
	InsertCustomer(*entities.Customer) error
	SelectCustomerByUsername(string) (*entities.Customer, error)
}
