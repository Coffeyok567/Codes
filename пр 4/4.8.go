//Корчагин Евгений 363

package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id   int
	jobs <-chan int
	wg   *sync.WaitGroup
}

type LoadBalancer struct {
	workers []*Worker
	jobs    chan int
	next    int
	mu      sync.Mutex
	wg      sync.WaitGroup
}

func (w *Worker) Start() {
	defer w.wg.Done()

	for job := range w.jobs {
		fmt.Printf("Воркер %d начал обработку запроса %d\n", w.id, job)
		time.Sleep(time.Duration(100) * time.Millisecond) // делает вид работы
		fmt.Printf("Воркер %d завершил обработку запроса %d\n", w.id, job)
	}
}

func NewLoadBalancer(workerCount int) *LoadBalancer {
	lb := &LoadBalancer{
		jobs: make(chan int, 10),
		next: 0,
	}

	for i := 1; i <= workerCount; i++ {
		worker := &Worker{
			id:   i,
			jobs: lb.jobs,
			wg:   &lb.wg,
		}
		lb.workers = append(lb.workers, worker)
		lb.wg.Add(1)
		go worker.Start()
	}

	return lb
}

func (lb *LoadBalancer) Dispatch(request int) {
	lb.jobs <- request
	lb.mu.Lock()
	lb.next = (lb.next + 1) % len(lb.workers)
	lb.mu.Unlock()
}

func (lb *LoadBalancer) Stop() {
	close(lb.jobs)
	lb.wg.Wait()
}

func main() {
	lb := NewLoadBalancer(3)

	for i := 1; i <= 10; i++ {
		lb.Dispatch(i)
		time.Sleep(50 * time.Millisecond)

		time.Sleep(1 * time.Second)

		lb.Stop()
		fmt.Println("Все запросы обработаны")
	}
}
