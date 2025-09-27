// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mu    sync.RWMutex
	tasks = map[int]string{
		0: "проснутся",
		1: "Прес качат",
		2: "бегит",
		3: "аджуманя",
	}
	nextID int = 4
)

func Reader(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	mu.Lock()
	task, exists := tasks[id]
	if exists {
		delete(tasks, id)
	}
	mu.Unlock()
	if exists {
		time.Sleep(2 * time.Second)
		fmt.Printf("Задача для  (id: %d):%s\n", id, task)
		fmt.Printf("читает задачи воркер %d\n", id)
	} else {
		fmt.Printf("Задача с ID : %d нет или уж выполняется", id)
	}
}

func Writer(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	newID := nextID
	nextID++
	tasks[newID] = "ест"
	mu.Unlock()
	fmt.Println("Задача создана")
}

func main() {
	var wg sync.WaitGroup
	CountWorker := 4
	for i := 0; i < CountWorker; i++ {
		wg.Add(1)
		go Reader(&wg, i)
	}
	wg.Add(1)
	go Writer(&wg)
	time.Sleep(2 * time.Second)

	wg.Add(1)
	go Reader(&wg, 4)

	wg.Wait()
}
