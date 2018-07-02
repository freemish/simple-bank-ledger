package bankledger

import (
	"errors"
	"time"

	"github.com/freemish/errgo"
)

/*
	This file is for abstracting data stores.
	What if I'm not using a database? What if I want to later?
*/

var (
	// CustomerList is a list of all users who have registered
	// (and only exists because of the absence of a database).
	CustomerList []*Customer

	// TransactionMap maps customer pointers to transactions.
	// Also only exists because of arbitrary absence of database.
	TransactionMap = make(map[*Customer]TransactionHistory)
)

// Errors having to do with storing data.
var (
	ErrNilCustomer = errors.New("Nil pointer to Customer")
)

// TransactionHistory represents a set of Transactions.
type TransactionHistory struct {
	Transactions []Transaction
	Balance      float64
}

// --- Customers and Logins ---

// SelectCustomerByAccountNo returns a pointer to a Customer.
// Can return nil if account number does not exist.
func SelectCustomerByAccountNo(accountNo string) *Customer {
	for _, cust := range CustomerList {
		if cust.AccountNumber == accountNo {
			return cust
		}
	}
	return nil
}

// SelectCustomerByUsername returns a pointer to a Customer.
// Can return nil if username does not exist.
func SelectCustomerByUsername(username string) *Customer {
	for _, cust := range CustomerList {
		if cust.Username == username {
			return cust
		}
	}
	return nil
}

// InsertCustomer adds a Customer to the data store.
func InsertCustomer(c *Customer) error {
	if c == nil {
		return errgo.Wrap(ErrNilCustomer)
	}
	CustomerList = append(CustomerList, c)
	return nil
}

// UpdateLastLogin updates customer's last login in data store.
func UpdateLastLogin(c *Customer) error {
	if c == nil {
		return errgo.Wrap(ErrNilCustomer)
	}
	c.LastLoginDate = time.Now()
	return nil
}

// --- Transactions ---

// InsertTransaction enters a new transaction for the logged-in customer.
func InsertTransaction(cust *Customer, t Transaction) error {
	if cust == nil {
		return errgo.Wrap(ErrNilCustomer)
	}
	history := TransactionMap[cust]
	history.Transactions = append(history.Transactions, t)
	history.Balance += t.Amount
	TransactionMap[cust] = history
	return nil
}

// reverseTransactionsDesc just reverses the list of transactions, because
// they will have been inserted in order, and it may be desirable to return
// them in descending order with respect to dates.
func reverseTransactionsList(list []Transaction) []Transaction {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// SelectAllTransactionHistory returns ALL transaction history for
// a customer. This is not a good idea for a web application with a database because
// it does not scale well at all. You will eventually have slow, slow requests.
func SelectAllTransactionHistory(cust *Customer) []Transaction {
	if cust == nil {
		var hist []Transaction
		return hist
	}
	return reverseTransactionsList(TransactionMap[cust].Transactions)
}

// SelectMonthTransactionHistory returns all transactions made in a given month.
func SelectMonthTransactionHistory(cust *Customer, month time.Month) []Transaction {
	if cust == nil {
		var hist []Transaction
		return hist
	}

	// assuming that transactions have been inserted in ascending order with
	// respect to date, we can start from the middle and do a binary search
	// for the first transaction of the month, then from there we can find
	// the last transaction made that month.

	hist := TransactionMap[cust].Transactions
	getFirst := getFirstTransactionOfMonthIndex

	// return subslice - could be of zero length and stuff, so VALIDATE this
	return hist[getFirst(hist, month):getFirst(hist, month+1)]
}

// assumes list sorted in ascended order
func getFirstTransactionOfMonthIndex(list []Transaction, month time.Month) int {
	var firstOfMonthIndex = -1
	middleIndex := len(list) / 2 // should be integer division

	for firstOfMonthIndex == -1 {
		if list[middleIndex].Date.Month() < month { // middle transaction is before desired month
			middleIndex = (middleIndex + len(list)) / 2
		} else if list[middleIndex].Date.Month() > month {
			middleIndex = middleIndex / 2
		} else { // transaction is within month
			// if transaction before this one is not the same month, return!
			firstOfMonthIndex = middleIndex // COME BACK TO THIS
		}
	}

	return firstOfMonthIndex
}

// SelectBalance returns a customer's balance.
func SelectBalance(cust *Customer) float64 {
	if cust == nil {
		return 0
	}
	return TransactionMap[cust].Balance
}
