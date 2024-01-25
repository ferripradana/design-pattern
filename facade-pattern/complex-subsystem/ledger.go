package complex_subsystem

import "fmt"

type Ledger struct {
}

func (l *Ledger) MakeEntry(accountId, txnType string, amount float32) {
	fmt.Printf("Make ledger entry for accountId %s with txnType %s for amount %f\n", accountId, txnType, amount)
	return
}
