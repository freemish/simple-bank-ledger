package processes

import (
	"testing"

	"github.com/freemish/simple-bank-ledger/entities"
)

// --- Mock persistence

type MockTransactionStore struct {
	Called                  bool
	LastInsertedTransaction *entities.Transaction
}

func (mts *MockTransactionStore) SelectTransactionsByUsername(username string, args ...map[string]interface{}) ([]*entities.Transaction, error) {
	return []*entities.Transaction{}, nil
}

func (mts *MockTransactionStore) SelectBalanceByUsername(username string) (float64, error) {
	return 0.0, nil
}

func (mts *MockTransactionStore) InsertTransaction(t *entities.Transaction) error {
	mts.LastInsertedTransaction = t
	mts.Called = true
	return nil
}

// --- Tests

func TestGetTranCode(t *testing.T) {
	type testStruct struct {
		is_debit  bool
		is_voided bool
		want      entities.TranCode
	}
	tests := []testStruct{
		{is_debit: true, is_voided: false, want: entities.Debit},
		{is_debit: false, is_voided: false, want: entities.Credit},
		{is_debit: true, is_voided: true, want: entities.VoidDebit},
		{is_debit: false, is_voided: true, want: entities.VoidCredit},
	}

	for _, test := range tests {
		got := GetTranCode(test.is_debit, test.is_voided)
		if test.want != got {
			t.Errorf("Got %s but expected %s", got, test.want)
		}
	}
}

func TestRecordTransaction(t *testing.T) {
	mts := &MockTransactionStore{}
	cust, _ := CreateAccount("molly", "Molly", "1234", nil)
	RecordTransaction(cust, false, 1.23, 0, mts)
	if !mts.Called {
		t.Errorf("Mock transaction store was not called")
	}
	if mts.LastInsertedTransaction.TranCode != entities.Credit {
		t.Errorf(
			"Expected inserted transaction to be credit, but got %s",
			mts.LastInsertedTransaction.TranCode,
		)
	}
}
