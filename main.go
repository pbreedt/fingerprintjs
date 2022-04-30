package main

import "fmt"

func main() {
	// var err error
	latency, err := ReadLatencyFile("./input/api_latencies.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println("Latency:", latency)

	txns, err := ReadTransactionFile("./input/transactions.csv")
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("All Transactions:", len(txns))

	processor := Processor{Latency: latency}
	ptxns := processor.Prioritize(txns, 50)
	r := processor.ProcessTransactions(ptxns)
	fmt.Println("Results", r)
}
