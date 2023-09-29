package race_condition

import (
	"fmt"
	"sync"
)

//Concurrently writes to the map m can be to unpredictable behavior or even a panic

var m = make(map[int]int)

func Map_example() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			m[i] = i
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("Map length:", len(m))
}
