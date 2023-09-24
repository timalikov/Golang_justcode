package main

import (
	"fmt"
)

// Deadlock with Buffered Channel
func bufferedRoutine(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	fmt.Println("Буфер канала заполнен, ожидание возможности отправить еще...")
	ch <- 5 // Зависание здесь
}

func main() {
	bufferedChannel := make(chan int, 4) // Размер буфера 4

	go bufferedRoutine(bufferedChannel)

	// Намеренно не читаем из канала
	select {}
}
