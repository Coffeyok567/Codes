// Корчагин Евгений 363
package main

import (
	"fmt"
	"time"
)

// логи
func main() {
	logCh := make(chan string, 100)

	go func() {
		for msg := range logCh {
			fmt.Print(msg)
			time.Sleep(10 * time.Millisecond)
		}
	}()

	// Рабы которые заполняют логи
	for i := 1; i <= 5; i++ {
		go func(id int) {
			for j := 1; j <= 5; j++ {
				msg := fmt.Sprintf("[Горутина %d]: Сообщение %d\n", id, j)
				logCh <- msg
				time.Sleep(50 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
	close(logCh)
	fmt.Println("Программа завершена")
}
