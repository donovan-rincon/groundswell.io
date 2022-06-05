package operations

import (
	"testing"

	"groundswell.io/server"
)

func TestDetermineOperation_SET_Command(t *testing.T) {
	server.Init()
	tests := []struct {
		name string
		args []string
	}{
		{"set cmd", []string{"set", "var_test", "10"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation(tt.args)
			item, exists := server.Ser.DB.Get(tt.args[1])
			if !exists {
				t.Errorf("variable should exists")
			}
			if tt.args[2] != item.Data {
				t.Errorf("%s should have a value of %v, it has %v", tt.args[1], tt.args[2], item.Data)
			}

		})
	}
}

func TestDetermineOperation_GET_Command(t *testing.T) {
	server.Init()
	tests := []struct {
		name string
		args []string
	}{
		{"get cmd", []string{"get", "var_test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation([]string{"set", "var_test", "10"})
			DetermineOperation(tt.args)
			item, exists := server.Ser.DB.Get(tt.args[1])
			if !exists {
				t.Errorf("variable should exists")
			}
			if item.Data != "10" {
				t.Errorf("%s should have a value of %v, it has %v", tt.args[1], "10", item.Data)
			}

		})
	}
}

func TestDetermineOperation_UNSET_Command(t *testing.T) {
	server.Init()
	tests := []struct {
		name string
		args []string
	}{
		{"unset cmd", []string{"unset", "var_test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation([]string{"set", "var_test", "10"})
			item, exists := server.Ser.DB.Get("var_test")
			if !exists {
				t.Errorf("variable should exists")
			}
			if item.Data != "10" {
				t.Errorf("%s should have a value of %v, it has %v", tt.args[1], "10", item.Data)
			}
			DetermineOperation(tt.args)
			_, exists = server.Ser.DB.Get(tt.args[1])
			if exists {
				t.Errorf("variable should not exists")
			}
		})
	}
}

func TestDetermineOperation_NUMEQUALTO_Command(t *testing.T) {
	server.Init()
	tests := []struct {
		name string
		args []string
	}{
		{"numequalto cmd", []string{"numequalto", "10"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation([]string{"set", "var_test1", "10"})
			DetermineOperation([]string{"set", "var_test2", "10"})

			DetermineOperation(tt.args)
			count := server.Ser.DB.NumEqualTo(tt.args[1])
			if count != 2 {
				t.Errorf("two variables should have a value of 10")
			}
		})
	}
}

func TestDetermineOperation_Transaction_Rollback(t *testing.T) {
	server.Init()
	tests := []struct {
		name          string
		args          []string
		expectedValue string
	}{
		{"rollback txn", []string{"get", "var_test"}, "10"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation([]string{"set", "var_test", "10"})
			DetermineOperation([]string{"begin"})
			DetermineOperation([]string{"set", "var_test", "15"})
			DetermineOperation([]string{"rollback"})

			DetermineOperation(tt.args)
			item, exists := server.Ser.DB.Get(tt.args[1])
			if !exists {
				t.Errorf("var_test should exists")
			}
			if item.Data != tt.expectedValue {
				t.Errorf("var_test should have a value of %v, but it has %v", tt.expectedValue, item.Data)

			}
		})
	}
}

func TestDetermineOperation_Transaction_Commit(t *testing.T) {
	server.Init()
	tests := []struct {
		name          string
		args          []string
		expectedValue string
	}{
		{"commit txn", []string{"get", "var_test"}, "15"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DetermineOperation([]string{"set", "var_test", "10"})
			DetermineOperation([]string{"begin"})
			DetermineOperation([]string{"set", "var_test", "15"})
			DetermineOperation([]string{"commit"})

			DetermineOperation(tt.args)
			item, exists := server.Ser.DB.Get(tt.args[1])
			if !exists {
				t.Errorf("var_test should exists")
			}
			if item.Data != tt.expectedValue {
				t.Errorf("var_test should have a value of %v, but it has %v", tt.expectedValue, item.Data)

			}
		})
	}
}
