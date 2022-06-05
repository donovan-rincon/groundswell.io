package operations

import (
	"fmt"
	"strings"

	"groundswell.io/datastore"
	"groundswell.io/server"
)

const (
	BEGIN    = "BEGIN"
	COMMIT   = "COMMIT"
	ROLLBACK = "ROLLBACK"
)

type Transaction struct {
	Commads []Command
	TmpDB   *datastore.Map
}

var Transactions []*Transaction

func InitTransactions() {
	Transactions = []*Transaction{}
}

func parseTransaction(params []string) {
	txn := params[0]
	switch strings.ToUpper(txn) {
	case BEGIN:
		Transactions = append(Transactions, &Transaction{
			TmpDB: datastore.NewMap(),
		})
	case COMMIT:
		if len(Transactions) == 0 {
			fmt.Println("NO TRANSACTION")
			return
		}
		lastTxn := Transactions[len(Transactions)-1]
		for _, cmd := range lastTxn.Commads {
			cmd.DB = server.Ser.DB
			cmd.processCommand(false)
		}
		Transactions = Transactions[:len(Transactions)-1]
	case ROLLBACK:
		if len(Transactions) == 0 {
			fmt.Println("NO TRANSACTION")
			return
		}
		Transactions = Transactions[:len(Transactions)-1]
	}
}
