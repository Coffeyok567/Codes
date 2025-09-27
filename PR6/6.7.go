// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

var productCount int
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
	provider := func(count int) {
		mutex.Lock()
		productCount = productCount + count
		time.Sleep(2 * time.Second)

		mutex.Unlock()
		wg.Done()
	}
	salesman := func(count int) {
		mutex.Lock()
		productCount = productCount - count
		time.Sleep(3 * time.Second)
		fmt.Printf("Поступило")
		mutex.Unlock()
		wg.Done()
	}

	go worker("Ivan")
	go worker("Georgy")
	go worker("Svetlana")
	go reader()
	go worker("Leonid")
	go worker("Evgeni")

	wg.Wait()
	reader()
	fmt.Printf("Le Fin.")
}
