package entities

import (
	"testing"
)

func TestCustomer(t *testing.T) {
	want := "ID: 0\tCreated: 0001-01-01T00:00:00Z\tName: Molly\tUsername: molly\tAccount Number: "
	cust := &Customer{
		Username: "molly",
		Name:     "Molly",
	}
	got := cust.DebugString()
	if got != want {
		t.Errorf("got:\n\t%v,\nwanted:\n\t%v", got, want)
	}
}

func TestTransaction(t *testing.T) {

}
