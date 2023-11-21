package main

import (
	"fmt"
	"sync"
	"time"
)

//В этом примере у нас есть безопасный счетчик, который использует,
//RWMutex для синхронизации доступа к внутреннему словарю.
//Мы одновременно инкрементируем и читаем значение счетчика из разных горутин.

type SafeCounter struct {
	m    map[string]int
	rwmu sync.RWMutex
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		m: make(map[string]int),
	}
}

func (sc *SafeCounter) Increment(key string) {
	sc.rwmu.Lock()
	defer sc.rwmu.Unlock()
	sc.m[key]++
}

func (sc *SafeCounter) Value(key string) int {
	sc.rwmu.RLock()
	defer sc.rwmu.RUnlock()
	return sc.m[key]
}

func main() {
	counter := NewSafeCounter()
	var wg sync.WaitGroup

	// Increment the counter in multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			counter.Increment("key")
			time.Sleep(10 * time.Millisecond)
			wg.Done()
		}()
	}

	// Read the counter value in multiple goroutines
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(counter.Value("key"))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter.Value("key"))
}
