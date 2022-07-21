package entities

import (
	"fmt"
	"time"
)

// Customer defines a user of the bank ledger
// who can make transactions on a bank account.
type Customer struct {
	ID           int
	Created      time.Time
	Name         string
	Username     string
	PasswordHash string

	// could consider making a model for tracking
	// individual login session history
	LastLoginDate time.Time

	// could consider separating if one customer can have
	// multiple bank accounts
	AccountNumber string
}

func (c Customer) DebugString() string {
	return fmt.Sprintf(
		"ID: %d\tCreated: %s\tName: %v\tUsername: %s\tAccount Number: %s",
		c.ID, c.Created.Format(time.RFC3339Nano), c.Name, c.Username, c.AccountNumber,
	)
}
