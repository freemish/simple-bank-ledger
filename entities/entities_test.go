package entities

import (
	"testing"
)

func printFailedTest(t *testing.T, got string, want string) {
	t.Errorf("\ngot:\n\t%v\nwanted:\n\t%v", got, want)
}

func TestCustomer(t *testing.T) {
	want := "ID: 0\tCreated: 0001-01-01T00:00:00Z\tName: Molly\tUsername: molly\tAccount Number: "
	cust := &Customer{
		Username: "molly",
		Name:     "Molly",
	}
	got := cust.DebugString()
	if got != want {
		printFailedTest(t, got, want)
	}
}

func TestTransaction(t *testing.T) {
	want := "ID: 0\tCreated: 0001-01-01T00:00:00Z\tIsDebit: false\tTranCode: \tAmount: $0.00\tCustomer ID: 0"
	tx := &Transaction{}
	got := tx.DebugString()
	if got != want {
		printFailedTest(t, got, want)
	}
}
