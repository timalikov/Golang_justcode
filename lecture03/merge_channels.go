package main

import (
	"fmt"
	"sync"
)

// Функция для объединения нескольких каналов в один
func merge(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup // Инициализация группы ожидания
	out := make(chan int) // Создаем выходной канал

	// Запускаем горутину для каждого входного канала
	for _, ch := range channels {
		wg.Add(1) // Увеличиваем счетчик группы ожидания
		go func(c <-chan int) {
			defer wg.Done()      // Уменьшаем счетчик по завершении
			for val := range c { // Читаем из входного канала
				out <- val // Пишем в выходной канал
			}
		}(ch)
	}

	// Запускаем горутину, которая закроет выходной канал после завершения всех входных каналов
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	// Создаем три канала
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	// Заполняем каналы в отдельных горутинах
	go func() { ch1 <- 1; ch1 <- 11; close(ch1) }()
	go func() { ch2 <- 2; ch2 <- 22; close(ch2) }()
	go func() { ch3 <- 3; ch3 <- 33; close(ch3) }()

	// Объединяем каналы в один
	merged := merge(ch1, ch2, ch3)

	// Читаем и выводим значения из объединенного канала
	for val := range merged {
		fmt.Println(val)
	}
}
