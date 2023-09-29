package race_condition

import (
	"fmt"
)

//Here, we don't know which goroutine will send its value first.
//The output could be either '1 2' or '2 1'.

var ch = make(chan int)

func Unbuffered_chanel_example() {
	go func() {
		ch <- 1
	}()
	go func() {
		ch <- 2
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
