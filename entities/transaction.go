package entities

import (
	"fmt"
	"time"
)

type TranCode string

const (
	Debit      TranCode = "DEB"
	VoidDebit  TranCode = "VOID_DEB"
	Credit     TranCode = "CRD"
	VoidCredit TranCode = "VOID_CRD"
)

// Transaction represents a withdrawal or deposit
// from a bank account.
type Transaction struct {
	ID          int
	Created     time.Time
	IsDebit     bool
	TranCode    TranCode
	Description string
	Layer       string
	Amount      float64
	Customer    Customer
}

func (t Transaction) DebugString() string {
	return fmt.Sprintf(
		"ID: %d\tCreated: %s\tIsDebit: %v\tTranCode: %s\tAmount: $%.2f\tCustomer ID: %d",
		t.ID, t.Created.Format(time.RFC3339Nano), t.IsDebit, t.TranCode, t.Amount, t.Customer.ID,
	)
}
