package main

import (
	"os"

	"groundswell.io/operations"
	"groundswell.io/reader"
	"groundswell.io/server"
)

func main() {
	server.Init()
	operations.InitTransactions()
	reader.ReadInput(os.Stdin)
}
