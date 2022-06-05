package datastore

import (
	"testing"
)

func TestMap_Set_Get(t *testing.T) {
	tests := []struct {
		name   string
		fields *Map
		args   *Item
	}{
		{"store item", NewMap(), NewItem("test", "10")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				items: tt.fields.items,
				nSize: tt.fields.nSize,
			}
			// Test set
			m.Set(tt.args)

			// Test get with set
			res, exists := m.Get(tt.args.Key)
			if !exists {
				t.Errorf("key should exists")
			}
			if res.Data != tt.args.Data {
				t.Errorf("data is not the same")
			}
		})
	}
}

func TestMap_Get(t *testing.T) {
	tests := []struct {
		name   string
		fields *Map
	}{
		{"not value exists item", NewMap()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				items: tt.fields.items,
				nSize: tt.fields.nSize,
			}
			// Test get no value exists
			_, exists := m.Get("var")
			if exists {
				t.Errorf("var should not exists should exists")
			}
		})
	}
}

func TestMap_Unset(t *testing.T) {
	tests := []struct {
		name   string
		fields *Map
		args   *Item
		want   int64
	}{
		{"unset existing variable", NewMap(), NewItem("test", "10"), 1},
		{"unset not existing variable", NewMap(), nil, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				items: tt.fields.items,
				nSize: tt.fields.nSize,
			}
			if tt.args == nil {
				if got := m.Unset("not_exists"); got != tt.want {
					t.Errorf("variable should not exists")
				}
				return
			}
			m.items.Store(tt.args.Key, tt.args.Data)
			if got := m.Unset(tt.args.Key); got != tt.want {
				t.Errorf("variable should be loaded and unset")
			}
		})
	}
}

func TestMap_Numequalto(t *testing.T) {
	tests := []struct {
		name   string
		fields *Map
		args   []*Item
		want   int64
	}{
		{"one variable with same value", NewMap(), []*Item{NewItem("test", "10")}, 1},
		{"two variables same value", NewMap(), []*Item{NewItem("test", "10"), NewItem("test2", "10")}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Map{
				items: tt.fields.items,
				nSize: tt.fields.nSize,
			}
			if tt.args == nil {
				if got := m.Unset("not_exists"); got != tt.want {
					t.Errorf("variable should not exists")
				}
				return
			}
			for _, item := range tt.args {
				m.Set(item)
			}

			if got := m.NumEqualTo("10"); got != tt.want {
				t.Errorf("Expected: %d , got: %d", tt.want, got)
			}
		})
	}
}
