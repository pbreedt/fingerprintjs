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
func (t Transaction) ValPerSec(latency map[string]int) float64 {
	if lat, ok := latency[t.BankCountryCode]; ok {
		return t.Amount / float64(lat)
	} else {
		return 0
	}
}

func (t Transaction) String() string {
	return fmt.Sprintf("ID:%s: Amount:%f, BankCountry:%s", t.ID, t.Amount, t.BankCountryCode)
}

type Processor struct {
	//stores latency data read from file
	Latency map[string]int
}

func (processor *Processor) ProcessTransactions(transactions []Transaction) []Result {
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
// that will maximize the USD value and fit the transactions under given time limit
func (processor *Processor) Prioritize(transactions []Transaction, totalTime int) []Transaction {
	//sort by value per second
	sort.Slice(transactions, func(i, j int) bool {
		valPerSecI := transactions[i].ValPerSec(processor.Latency)
		valPerSecJ := transactions[j].ValPerSec(processor.Latency)
		if valPerSecI == valPerSecJ {
			return transactions[i].Amount > transactions[j].Amount
		}
		// return valPerSecI > valPerSecJ
		return valPerSecI > valPerSecJ
	})

	priority := make([]Transaction, 0)
	countTime := 0
	countAmount := 0.0
	for _, txn := range transactions {
		// fmt.Printf("Txn:%v (value/s:%f)\n", txn, txn.ValPerSec(processor.Latency))
		//keep adding transactions as long as time limit allows
		if (countTime + processor.Latency[txn.BankCountryCode]) <= totalTime {
			countTime += processor.Latency[txn.BankCountryCode]
			countAmount += txn.Amount
			priority = append(priority, txn)
		}
	}

	fmt.Printf("Total processing time:%d\n", countTime)
	fmt.Printf("Total value processed:%.2f\n", countAmount)
	return priority
}
