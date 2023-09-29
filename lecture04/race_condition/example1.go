package race_condition

import (
	"fmt"
	"sync"
)

//Simple Increment: The output is unpredictable and won't necessarily be 1000

var counter int

func Increment_example() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			counter++
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
