//Корчагин Евгений 363

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		result := job * job
		fmt.Printf("воркер %d обработал: %d -> %d", id, job, result)
		results <- result
	}
}

func main() {

	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	fmt.Println("вывод результатов:")
	for result := range results {
		fmt.Println(result)
	}
}
