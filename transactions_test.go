package main

import (
	"fmt"
	"testing"
)

var (
	lat map[string]int = map[string]int{"aa": 1, "bb": 2}
	tx1 Transaction    = Transaction{ID: "a-a", BankCountryCode: "aa", Amount: 10.0} // highest value for 1s
	tx2 Transaction    = Transaction{ID: "a-b", BankCountryCode: "aa", Amount: 5.0}
	tx3 Transaction    = Transaction{ID: "b-a", BankCountryCode: "bb", Amount: 20.0} // highest value for 2s
	tx4 Transaction    = Transaction{ID: "b-b", BankCountryCode: "bb", Amount: 5.0}  // least desirable txn
	p   Processor      = Processor{Latency: lat}
)

func TestPrioritize1s(t *testing.T) {
	txns := []Transaction{tx1, tx2, tx3, tx4}
	prioritizeXs(t, txns, 1, []Transaction{tx1}, []Transaction{tx2, tx3, tx4})
}

func TestPrioritize2s(t *testing.T) {
	txns := []Transaction{tx1, tx2, tx3, tx4}
	prioritizeXs(t, txns, 2, []Transaction{tx3}, []Transaction{tx1, tx2, tx4})
}

func TestPrioritize3s(t *testing.T) {
	txns := []Transaction{tx1, tx2, tx3, tx4}
	prioritizeXs(t, txns, 3, []Transaction{tx1, tx3}, []Transaction{tx2, tx4})
}

func TestPrioritize4s(t *testing.T) {
	txns := []Transaction{tx1, tx2, tx3, tx4}
	prioritizeXs(t, txns, 4, []Transaction{tx1, tx2, tx3}, []Transaction{tx4})
}

func prioritizeXs(t *testing.T, txns []Transaction, timeLimit int, shouldContain []Transaction, shouldNotContain []Transaction) {
	ptxns := p.Prioritize(txns, timeLimit)
	fmt.Println("prioritized txns", ptxns)
	for _, tx := range shouldContain {
		if !contains(ptxns, tx) {
			t.Fatalf("prioritized transactions should have contained transaction %v", tx)
		}
	}

	for _, tx := range shouldNotContain {
		if contains(ptxns, tx) {
			t.Fatalf("prioritized transactions should not contain transaction %v", tx)
		}
	}

}

func contains(txs []Transaction, tx Transaction) bool {
	for _, t := range txs {
		if t == tx {
			return true
		}
	}
	return false
}
