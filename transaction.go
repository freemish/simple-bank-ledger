package bankledger

import (
	"errors"
	"time"

	"github.com/freemish/errgo"
)

var (
	// MinimumBalance represents the lowest bank balance a Customer is allowed.
	// For this bank, it's going to be 0.
	MinimumBalance = 0
)

// Errors related to transactions.
var (
	ErrBalanceTooLow     = errors.New("Balance is too low to authorize transaction")
	ErrFailedToAuthorize = errors.New("Customer does not have an active session")
)

// Transaction represents a withdrawal or deposit
// from a bank account.
type Transaction struct {
	Date        time.Time
	Name        string
	Description string
	Amount      int
}

// TransactionHistory represents a set of Transactions.
type TransactionHistory struct {
	Transactions []Transaction
	Balance      int
}

// RecordTransaction adds another transaction to a customer's history.
// Rejects transaction if balance would drop below minimum or customer isn't logged in.
func (cust *Customer) RecordTransaction(name, descr string, amount int) error {
	if cust == nil {
		return errgo.Wrap(ErrFailedToAuthorize)
	}

	balance := SelectBalance(cust)
	if (balance + amount) < MinimumBalance {
		return errgo.Wrap(ErrBalanceTooLow)
	}

	transaction := Transaction{
		Date:        time.Now(),
		Name:        name,
		Description: descr,
		Amount:      amount,
	}

	err := InsertTransaction(cust, transaction)
	if err != nil {
		return errgo.Wrap(err)
	}
	return nil
}

// GetAllHistory gets all history for a customer.
func (cust *Customer) GetAllHistory() []Transaction {
	return SelectAllTransactionHistory(cust)
}

// GetMonthHistory eventually only fetches transactions in given month.
func (cust *Customer) GetMonthHistory(month time.Month) []Transaction {
	return cust.GetAllHistory() // COME BACK TO THIS
}
