package processes

import (
	"github.com/freemish/simple-bank-ledger/entities"
)

func RecordTransaction(cust *entities.Customer, is_debit bool, amount float64, voided_tx int, its ITransactionStore) (*entities.Transaction, error) {
	tx := &entities.Transaction{
		Customer: *cust,
		ID:       voided_tx,
		IsDebit:  is_debit,
		Amount:   amount,
		TranCode: GetTranCode(is_debit, bool(voided_tx == 0)),
	}
	err := its.InsertTransaction(tx)
	return tx, err
}

func GetTranCode(is_debit, is_voided bool) entities.TranCode {
	if is_debit && !is_voided {
		return entities.Debit
	}
	if !is_debit && !is_voided {
		return entities.Credit
	}
	if is_debit {
		return entities.VoidDebit
	}
	return entities.VoidCredit
}
