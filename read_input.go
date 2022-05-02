package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadLatencyFile(filepath string) (map[string]int, error) {
	dat, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading latency file %s: %s", filepath, err.Error())
	}

	lat := make(map[string]int, 0)
	err = json.Unmarshal(dat, &lat)
	if err != nil {
		return nil, fmt.Errorf("error parsing latency file data:%s", err.Error())
	}

	return lat, nil
}

func ReadTransactionFile(filepath string, latency map[string]int) ([]Transaction, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("error opening transaction file %s:%s", filepath, err.Error())
	}
	defer file.Close()

	// could have used csv.NewReader(f).ReadAll(), but this is simple enough
	transactions := make([]Transaction, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if len(line) > 0 {
			fields := strings.Split(line, ",")
			if len(fields) >= 3 {
				amt, err := strconv.ParseFloat(fields[1], 64)
				if err != nil {
					fmt.Printf("WARNING: amount value '%s' is not a valid amount. (%s)\n", fields[1], line)
				} else {
					txn := Transaction{ID: fields[0], Amount: amt, BankCountryCode: fields[2], Latency: latency[fields[2]]}
					transactions = append(transactions, txn)
				}
			} else {
				fmt.Printf("WARNING: input line is invalid - does not have 3 fields. (%s)\n", line)
			}
		} else {
			fmt.Printf("WARNING: input line is empty. (%s)\n", line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("transaction file encountered error: %s", err.Error())
		return transactions, err
	}

	return transactions, nil
}
