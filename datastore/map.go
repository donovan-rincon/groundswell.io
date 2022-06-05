package datastore

import (
	"sync"
	"sync/atomic"
)

// Threadsafe map
type Map struct {
	items sync.Map
	nSize int64
}

func NewMap() *Map {
	m := &Map{}
	return m
}

func (m *Map) Set(v *Item) {
	m.items.Store(v.Key, v)
	atomic.AddInt64(&m.nSize, 1)
}

func (m *Map) Get(k string) (*Item, bool) {
	v, ok := m.items.Load(k)
	if !ok {
		return nil, false
	}
	return v.(*Item), true
}

func (m *Map) Unset(k string) int64 {
	prevSize := atomic.LoadInt64(&m.nSize)
	unsetN := m.unset(k)
	atomic.AddInt64(&m.nSize, -unsetN)

	return prevSize - atomic.LoadInt64(&m.nSize)
}

func (m *Map) unset(k string) int64 {
	_, loaded := m.items.LoadAndDelete(k)
	if !loaded {
		return 0
	}

	return int64(1)

}

func (m *Map) NumEqualTo(v any) int64 {
	count := 0
	m.items.Range(func(_, value any) bool {
		item := value.(*Item)
		if item.Data == v {
			count++
		}
		return true
	})
	return int64(count)
}
