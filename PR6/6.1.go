// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Mutex
var mutex sync.RWMutex
var pages int

func reader() {
	mutex.RLock()
	fmt.Println("Количество подсчитанных страниц:", pages)
	mutex.RUnlock()
}

func Pages(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	time.Sleep(1 * time.Second)
	fmt.Println("Страница подсчитана")
	pages++
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	Worker := 3
	for i := 0; i <= Worker; i++ {
		wg.Add(1)
		go Pages(&wg)
	}
	wg.Wait()
	reader()
}
