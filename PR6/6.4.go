// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

var votes []string
var wg sync.WaitGroup
var mutex sync.RWMutex

func reader() {
	mutex.RLock()
	fmt.Print("Количество голосов: ")
	fmt.Println(len(votes))
	mutex.RUnlock()
}

func main() {

	wg.Add(5)
	worker := func(name string) {
		mutex.Lock()
		votes = append(votes, name)
		time.Sleep(2 * time.Second)
		fmt.Println("Голос учтен")
		mutex.Unlock()
		wg.Done()
	}
	go worker("Ivan")
	go worker("Georgy")
	go worker("Svetlana")
	go reader()
	go worker("Leonid")
	go worker("Evgeniy")

	wg.Wait()
	reader()
	fmt.Printf("Le Fin.")
}
