# FingerprintJS
> A take-home assignment for FingerprintJS.

## General info
This project takes two files as input:
- ./input/api_latencies.json: contains latency data for country codes
- ./input/transactions.csv: contains the transaction data for processing

The aim it to maximize the transaction value that is being processed in a given time.

## Technologies
* Go - version 1.17

## Getting started
First get the source code by cloning the repository from GitHub:

```git clone https://github.com/pbreedt/fingerprintjs.git```

...and change into the newly created directory.

Then, either build the app:

```go build```

...or run it straight from source code:

```go run . 1000```

### Usage:
```
./fingerprintjs timeLimit [latency file] [transactions file]  
        timeLimit               - number of total time allowed for processing  
        latency file            - path to the latency JSON file (optional: default to ./input/api_latencies.json)  
        transactions file       - path to the transactions CSV file (optional: default to ./input/transactions.csv)  
```

## Questions & answers:
The [question that was posted](https://gist.github.com/Valve/70d01f51885b4d0cdfeea1a576a23850) was:  
> What is the max USD value that can be processed in 50ms, 60ms, 90ms, 1000ms?

My findings are as follows:  
| Time frame (ms) | Total Value | New Total|Actual ms|
|:----------------|------------:|---------:|:--------|
| 50              |     3637.98 | 4139.43  |  (50)   |
| 60              |     4362.01 | 4624.86  |  (56)   |
| 90              |     6870.48 | 6972.29  |  (90)   |
| 1000            |    35471.81 | 35471.81 |  (1000) |

__The algorithm__

The basic principle is to calculate the 'value per second' for each transaction.  This is, the USD value that will be processed for every 1s spend processing the particular transaction.

Next, I sort all the transactions according to the 'value per second' values, with the highest amounts taking preference if any two transactions have the same 'value per second'.  

IMPROVEMENT:  
Keep track of the time difference between processing time and total available time.  Try and fill this time a little bit better.

__General code__

I tried to stay as close as possible to the originally provided function signatures, thus keeping the same input and return parameters.

Various 'debug' output lines have been commented out, but left in the code.

Some unit tests are included to check for predictable results for controlled input.  

## Contact

Thank you for the opportunity, I had fun doing this assignment!

Written by [@pbreedt](mailto:petrus.breedt@gmail.com)