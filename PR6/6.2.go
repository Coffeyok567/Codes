// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

// я вчера забыл дочитать лекцию(    я сейчас на ходу разбираюсь , вроде понял как работает   вы не подумайте я не бездельник <3
var mu sync.RWMutex
var cache = map[int]string{
	0: "чай",
	1: "Айзен",
	2: "<3",
	3: "Изначально никто не живет на небесах: ни ты, ни я, ни даже Бог. Но скоро пустое место на небесном троне, которое так грустно видеть, будет занято. Отныне я буду жить на небесах",
}

func Reader(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.RLock()
	defer mu.RUnlock()
	time.Sleep(2 * time.Second)
	fmt.Println(cache[id])
	fmt.Printf("кэш прочитан воркером с id: %d \n", id)
}

func Writer(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.Lock()
	time.Sleep(2 * time.Second)
	cache[id] = "Айзен"
	fmt.Println("Перезаписанный кэш:" + cache[id])
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	Worker := 3
	for i := 0; i <= Worker; i++ {
		wg.Add(1)
		go Reader(&wg, i)
	}
	go Writer(&wg, 1)
	wg.Wait()
}
