package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	timeLimit, latencyFile, transactionsFile := processArgs()

	latency, err := ReadLatencyFile(latencyFile)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Println("Latency:", latency)

	txns, err := ReadTransactionFile(transactionsFile, latency)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Total number of transactions:", len(txns))

	processor := Processor{Latency: latency}
	ptxns := processor.Prioritize(txns, timeLimit)
	r := processor.ProcessTransactions(ptxns)
	fmt.Println("Number of results", len(r))
}

func processArgs() (int, string, string) {
	if len(os.Args) <= 1 {
		usage()
	}

	timeLimit := 1000
	var err error
	if len(os.Args[1]) > 0 {
		timeLimit, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("First argumant is invalid - must be a number")
			usage()
		}
		fmt.Println("Using processing time limit of:", timeLimit)
	}

	latencyFile := "./input/api_latencies.json"
	if len(os.Args) > 2 && len(os.Args[2]) > 0 {
		latencyFile = os.Args[2]
		fmt.Println("Using latency file from arguments:", latencyFile)
	}

	txnsFile := "./input/transactions.csv"
	if len(os.Args) > 3 && len(os.Args[3]) > 0 {
		txnsFile = os.Args[3]
		fmt.Println("Using transactions file from arguments:", txnsFile)
	}

	return timeLimit, latencyFile, txnsFile
}

func usage() {
	fmt.Println("Usage:")
	fmt.Println(" ./fingerprintjs timeLimit [latency file] [transactions file]")
	fmt.Println("\ttimeLimit\t\t- number of total time allowed for processing")
	fmt.Println("\tlatency file\t\t- path to the latency JSON file (optional: default to ./input/api_latencies.json)")
	fmt.Println("\ttransactions file\t- path to the transactions CSV file (optional: default to ./input/transactions.csv)")
	os.Exit(1)
}
