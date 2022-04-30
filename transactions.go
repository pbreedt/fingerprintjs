package main

import (
	"fmt"
	"sort"
)

type Transaction struct {
	ID              string
	Amount          float64
	BankCountryCode string
}

type Result struct {
	ID         string
	Fraudulent bool
}

// calculate the value per second for given transaction
func (t Transaction) valPerSec() float64 {
	if lat, ok := latency[t.BankCountryCode]; ok {
		return t.Amount / float64(lat)
	} else {
		return 0
	}
}

func (t Transaction) String() string {
	return fmt.Sprintf("ID:%s: Amount:%f, BankCountry:%s, Latency:%d, ValuePerSec:%f", t.ID, t.Amount, t.BankCountryCode, latency[t.BankCountryCode], t.valPerSec())
}

func processTransactions(transactions []Transaction) []Result {
	result := make([]Result, 0)
	for _, txn := range transactions {
		r := Result{ID: txn.ID}
		r.Fraudulent = processTransaction(txn)
		result = append(result, r)
	}

	return result
}

//fake processing function - without latency applied
//(transactions with ID's starting with 'a' are considered fraudulent)
func processTransaction(transaction Transaction) bool {
	fmt.Println("Process txn: ", transaction)
	return (string(transaction.ID[0]) == "a")
}

// function should return a subset (or full array)
// that will maximize the USD value and fit the transactions under 1 second
func prioritize(transactions []Transaction, totalTime int) []Transaction {
	//sort by value per second
	sort.Slice(transactions, func(i, j int) bool {
		valPerSecI := transactions[i].Amount / float64(latency[transactions[i].BankCountryCode])
		valPerSecJ := transactions[j].Amount / float64(latency[transactions[j].BankCountryCode])
		return valPerSecI > valPerSecJ
	})

	priority := make([]Transaction, 0)
	countTime := 0
	countAmount := 0.0
	for _, txn := range transactions {
		// fmt.Printf("Txn:%v\n", txn)
		//keep adding transactions as long as time limit allows
		if (countTime + latency[txn.BankCountryCode]) <= totalTime {
			countTime += latency[txn.BankCountryCode]
			countAmount += txn.Amount
			priority = append(priority, txn)
		}
	}

	fmt.Println("Time:", countTime, "Amount:", countAmount)
	return priority
}
