package main

import "fmt"

//stores latency data read from file
var latency map[string]int

func main() {
	var err error
	latency, err = ReadLatencyFile("./input/api_latencies.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println("Latency:", latency)

	txns, err := ReadTransactionFile("./input/transactions.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Transactions:", len(txns))

	ptxns := prioritize(txns, 50)
	r := processTransactions(ptxns)
	fmt.Println("Results", r)
}
