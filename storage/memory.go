package storage

import (
	"fmt"
	"sync"
)

// Memory Key:Value Storage
type MemoryStore[K comparable, V any] struct {
	mu     sync.RWMutex
	memory map[K]V
}

func NewMemoryStore[K comparable, V any]() *MemoryStore[K, V] {
	return &MemoryStore[K, V]{
		memory: make(map[K]V),
		mu:     sync.RWMutex{},
	}
}

func (m *MemoryStore[K, V]) Put(key K, value V) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.memory[key] = value
	return nil
}

func (m *MemoryStore[K, V]) List(search func(k K) bool) (keys []K, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for key := range m.memory {
		if search(key) {
			keys = append(keys, key)
		}
	}

	return keys, nil
}

func (m *MemoryStore[K, V]) Get(key K) (value V, err error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	value, ok := m.memory[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}
	return value, nil
}

func (m *MemoryStore[K, V]) Has(key K) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	_, ok := m.memory[key]
	if !ok{
		return false
	} else {
		return true
	}
}
 
/*func (m *MemoryStore) Update(key string, value string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, ok := m.memory[key]
	if !ok {
		return fmt.Errorf("the key (%v) does not exists", key)
	}

	m.memory[key] = value
	return nil
}*/

func (m *MemoryStore[K, V]) Delete(key K) (value V, err error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	value, ok := m.memory[key]
	if !ok {
		return value, fmt.Errorf("the key (%v) does not exists", key)
	}

	delete(m.memory, key)
	return value, nil
}
