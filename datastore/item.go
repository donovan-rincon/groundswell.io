package datastore

import "unsafe"

type Item struct {
	Key  string
	Size uint32
	Data any
	// Could add expiration time and creation time
}

func NewItem(key string, data any) *Item {
	item := &Item{
		Key:  key,
		Data: data,
		Size: uint32(unsafe.Sizeof(data)),
	}
	return item
}
