package reader

import (
	"bufio"
	"io"
	"strings"

	"groundswell.io/operations"
)

func ReadInput(stdin io.Reader) {
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		operations.DetermineOperation(tokens)
	}
}
