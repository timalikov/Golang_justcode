package main

import (
	"fmt"
	"sync"
)

//Эта реализация обеспечивает безопасный доступ к Мапу из разных горутин
//благодаря использованию мьютексов.

// SafeMap - thread-safe map based on mutex
type SafeMap struct {
	m    map[interface{}]interface{}
	lock sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[interface{}]interface{}),
	}
}

// Load - retrieves the value for a key
func (sm *SafeMap) Load(key interface{}) (value interface{}, ok bool) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	value, ok = sm.m[key]
	return
}

// Store - sets the value for a key
func (sm *SafeMap) Store(key, value interface{}) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	sm.m[key] = value
}

// Delete - removes the key from the map
func (sm *SafeMap) Delete(key interface{}) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	delete(sm.m, key)
}

func main() {
	sm := NewSafeMap()

	sm.Store("hello", "world")

	// Load value
	if v, ok := sm.Load("hello"); ok {
		fmt.Println("Loaded value:", v)
	}

	sm.Delete("hello")

	// Try loading again
	if _, ok := sm.Load("hello"); !ok {
		fmt.Println("Key 'hello' not found!")
	}
}
