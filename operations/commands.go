package operations

import (
	"errors"
	"fmt"
	"strings"

	"groundswell.io/datastore"
)

const (
	SET        = "SET"
	GET        = "GET"
	UNSET      = "UNSET"
	NUMEQUALTO = "NUMEQUALTO"
	END        = "END"
)

type Command struct {
	Cmd    string
	Params []string
	DB     *datastore.Map
}

func parseCommand(params []string, db *datastore.Map) Command {
	cmd := Command{
		Cmd:    params[0],
		Params: params[1:],
		DB:     db,
	}
	return cmd
}

func (cmd *Command) processCommand(print bool) error {
	switch strings.ToUpper(cmd.Cmd) {
	case SET:
		if len(cmd.Params) != 2 {
			return errors.New("incorrect set command")
		}
		var item datastore.Item
		item.Key = cmd.Params[0]
		item.Data = cmd.Params[1]
		cmd.DB.Set(&item)
		if print {
			fmt.Println()
		}
	case GET:
		if len(cmd.Params) != 1 {
			return errors.New("incorrect get command")
		}
		item, exists := cmd.DB.Get(cmd.Params[0])
		if !exists {
			fmt.Println("Nil")
			return nil
		}
		if print {
			fmt.Println(item.Data)
		}
	case UNSET:
		if len(cmd.Params) != 1 {
			return errors.New("incorrect unset command")
		}
		cmd.DB.Unset(cmd.Params[0])
		if print {
			fmt.Println()
		}
	case NUMEQUALTO:
		if len(cmd.Params) != 1 {
			return errors.New("incorrect numequalto command")
		}
		count := cmd.DB.NumEqualTo(cmd.Params[0])
		if print {
			fmt.Println(count)
		}
	default:
		return errors.New("unsupported comamnd")
	}
	return nil
}
