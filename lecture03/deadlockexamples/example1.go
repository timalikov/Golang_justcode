package main

import "fmt"

// Deadlock with a Function Argument
func specialRoutine(ch chan int, flag chan bool) {
	ch <- 1
	fmt.Println("Ожидание следующего числа...")
	<-ch // Зависание здесь, ожидание второй записи
	flag <- true
}

func main() {
	numChannel := make(chan int)
	done := make(chan bool)

	go specialRoutine(numChannel, done)

	val := <-numChannel
	fmt.Println("Получено:", val)

	// Бесконечное ожидание, приводящее к зависанию
	<-done
}
