package main

import "fmt"

// Deadlock with Conditional Wait
func conditionalWait(ch chan int, condition *bool) {
	if !*condition {
		fmt.Println("Горутина ожидает условия.")
		<-ch // Зависание здесь
	}
}

func main() {
	dataChannel := make(chan int)
	var conditionFlag bool = false

	go conditionalWait(dataChannel, &conditionFlag)

	// Основная горутина также ждет, приводя к зависанию
	fmt.Println("Основная горутина ожидает.")
	<-dataChannel
}
