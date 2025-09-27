// Корчагин Евгений 363
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

// Структура для хранения результатов
type Result struct {
	WorkerID int
	JobID    int
	Output   int
}

var results []Result

func producer(id int, jobs chan<- int, numJobs int) {
	defer wg.Done()

	for i := 1; i <= numJobs; i++ {
		jobID := id*100 + i // Создаем уникальный ID задачи
		fmt.Printf("Продюсер %d отправил задачу: %d\n", id, jobID)
		jobs <- jobID
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Printf("Продюсер %d завершил работу\n", id)
}

func worker(id int, jobs <-chan int) {
	defer wg.Done()

	for job := range jobs {
		// Обрабатываем задачу
		time.Sleep(1 * time.Second)
		result := job * 2 // Пример обработки: умножаем на 2

		// Записываем результат с использованием мьютекса
		mutex.Lock()
		results = append(results, Result{
			WorkerID: id,
			JobID:    job,
			Output:   result,
		})
		mutex.Unlock()

		fmt.Printf("Воркер %d обработал: %d -> %d\n", id, job, result)
	}
	fmt.Printf("Воркер %d завершил работу\n", id)
}

func main() {
	// Создаем буферизированный канал
	jobs := make(chan int, 10)

	// Запускаем воркеров
	workers := 3
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs)
	}

	// Запускаем продюсеров
	producers := 2
	numJobsPerProducer := 5

	for i := 1; i <= producers; i++ {
		wg.Add(1)
		go producer(i, jobs, numJobsPerProducer)
	}

	// Ждем завершения всех продюсеров в отдельной горутине
	go func() {
		// Даем время продюсерам отправить задачи
		time.Sleep(3 * time.Second)
		close(jobs) // Закрываем канал только после отправки всех задач
	}()

	// Ждем завершения всех воркеров
	wg.Wait()

	// Выводим статистику
	fmt.Println("\n=== Статистика обработки ===")
	mutex.RLock()
	fmt.Printf("Всего обработано задач: %d\n", len(results))
	for _, result := range results {
		fmt.Printf("Задача %d обработана воркером %d, результат: %d\n",
			result.JobID, result.WorkerID, result.Output)
	}
	mutex.RUnlock()
}
