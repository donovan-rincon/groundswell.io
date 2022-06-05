package operations

import (
	"fmt"
	"os"
	"strings"

	"groundswell.io/server"
)

/* type Transaction struct {
	Transaction
} */

func DetermineOperation(params []string) {
	if len(params) == 1 && strings.ToUpper(params[0]) == END {
		// Finish program
		os.Exit(1)
	}

	// If there is only one parameter different to END cmdn, it should be a transaction command <BEGIN> <ROLLACK> <COMMIT>
	if len(params) == 1 {
		parseTransaction(params)
		return
	}

	if len(Transactions) > 0 {
		txn := Transactions[len(Transactions)-1]
		cmd := parseCommand(params, txn.TmpDB)
		txn.Commads = append(txn.Commads, cmd)
		cmd.processCommand(true)
		return
	}

	cmd := parseCommand(params, server.Ser.DB)
	err := cmd.processCommand(true)
	if err != nil {
		fmt.Println(err)
	}
}
