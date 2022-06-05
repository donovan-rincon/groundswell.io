# GROUNDSWELL.IO

This program is intended to simulate a REDIS like db written only with go and no third party libraries or packages.

# Approach

The solution used was to implement a sync.Map structure in order to store the values with the allowed commands and transactions. 
The reason that sync.Map was used is for future prof of concurrent connections to the server/db and multiple write/reads while reducing the probability of errors with locks and writing same value at a time


# Allowed Commands and Transactions
## Commands
- SET name value - Set the corresponding name to the related value. Neither var names or values can contain spaces.
- GET name - Prints out the value of the variable name, or Nil if the variable is not set.
- UNSET name - unset the variable name. Meaning if a var is set, then unset - getting said var will return Nil.
- NUMEQUALTO value - returns the number of variables that are set to the value in question. Empty results should return 0.

Example
```
Input                        Output
SET test-var-name 100
GET test-var-name             100 
UNSET test-var-name
GET test-var-name             Nil
SET test-var-name-1 50
SET test-var-name-2 50
NUMEQUALTO 50                 2
SET test-var-name-2 10
NUMEQUALTO 50                 1
END
```

## Transactions
- BEGIN - Opens a new transaction block. Transactions can be nested.
- ROLLBACK - Undo all of the commands in the contextual transaction block. If no transaction is open your program should print “NO TRANSACTION”.
- COMMIT - Close all open transaction blocks, committing all changes. Print nothing if the transaction committing is successful. Print “NO TRANSACTION” if a transaction block is not open.

Transactions can be nested also

Examples

```
Input                         Output
GET test-var-name             Nil 
BEGIN
SET test-var-name 100
GET test-var-name             100 
COMMIT
GET test-var-name             100
```

```
Input                         Output 
GET test-var-name             Nil 
BEGIN
SET test-var-name 100
GET test-var-name             100 
BEGIN
SET test-var-name 120
GET test-var-name             120 
BEGIN
SET test-var-name 150
GET test-var-name             150 
ROLLBACK
COMMIT
GET test-var-name             120
```

# How to run the program

## Prerequisites
Go should be installed in order to build the code

## Steps
- Open a terminal and run the cmd `go build` from the root folder where the code was downloaded in order to build the binary with the output of `groundswell.io` (the name of the module)

- Run the binary file with `./groundswell.io`

- Start adding the allowed commands and transactions

# Next Steps
- Add to the server the connection params (host, port, etc) and start a server instead of running the binary file
- Implement clients in order to connect to the server and allow multiple connections instead of only one from the running program
- Store data in a file so it can be loaded when starting or after a crash
