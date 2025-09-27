// Корчагин Евгений 363
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ErrorInfo содержит информацию об ошибке
type ErrorInfo struct {
	Stage     string
	Message   string
	Timestamp time.Time
	Data      interface{}
}

// Pipeline конвейер обработки данных
type Pipeline struct {
	mu     sync.Mutex
	errors []ErrorInfo
}

// NewPipeline создает новый конвейер
func NewPipeline() *Pipeline {
	return &Pipeline{
		errors: make([]ErrorInfo, 0),
	}
}

// RecordError записывает ошибку потокобезопасно
func (p *Pipeline) RecordError(stage, message string, data interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.errors = append(p.errors, ErrorInfo{
		Stage:     stage,
		Message:   message,
		Timestamp: time.Now(),
		Data:      data,
	})
}

// GetErrors возвращает все ошибки
func (p *Pipeline) GetErrors() []ErrorInfo {
	p.mu.Lock()
	defer p.mu.Unlock()
	return append([]ErrorInfo{}, p.errors...)
}

// Stage 1: Валидация данных
func validateData(data int, pipeline *Pipeline, output chan<- int) {
	if data < 0 {
		pipeline.RecordError("validation", "отрицательное число", data)
		return
	}
	if data > 1000 {
		pipeline.RecordError("validation", "число слишком большое", data)
		return
	}
	output <- data
}

// Stage 2: Обработка данных
func processData(data int, pipeline *Pipeline, output chan<- int) {
	// Имитация случайной ошибки обработки
	if rand.Float32() < 0.1 { // 10% вероятность ошибки
		pipeline.RecordError("processing", "ошибка обработки", data)
		return
	}

	// Умножаем данные на 2
	result := data * 2
	output <- result
}

// Stage 3: Сохранение данных
func saveData(data int, pipeline *Pipeline, output chan<- int) {
	// Имитация случайной ошибки сохранения
	if rand.Float32() < 0.05 { // 5% вероятность ошибки
		pipeline.RecordError("saving", "ошибка сохранения", data)
		return
	}

	output <- data
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pipeline := NewPipeline()
	var wg sync.WaitGroup

	// Каналы для передачи данных между стадиями
	validateCh := make(chan int, 10)
	processCh := make(chan int, 10)
	saveCh := make(chan int, 10)

	// Запускаем стадию валидации (3 горутины)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for data := range validateCh {
				validateData(data, pipeline, processCh)
			}
		}(i)
	}

	// Запускаем стадию обработки (2 горутины)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for data := range processCh {
				processData(data, pipeline, saveCh)
			}
		}(i)
	}

	// Запускаем стадию сохранения (2 горутины)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for data := range saveCh {
				saveData(data, pipeline, make(chan int, 1)) // Игнорируем вывод
			}
		}(i)
	}

	// Горутина для закрытия каналов после завершения
	go func() {
		wg.Wait()
		close(processCh)
		close(saveCh)
	}()

	// Отправляем данные в конвейер
	for i := 0; i < 20; i++ {
		data := rand.Intn(1500) - 100 // Генерируем данные от -100 до 1400
		validateCh <- data
	}

	close(validateCh)
	wg.Wait()

	// Выводим отчет об ошибках
	errors := pipeline.GetErrors()
	fmt.Printf("\n=== ОТЧЕТ ОБ ОШИБКАХ ===\n")
	fmt.Printf("Всего ошибок: %d\n", len(errors))
	for i, err := range errors {
		fmt.Printf("%d. [%s] %s: %v (время: %s)\n",
			i+1, err.Stage, err.Message, err.Data, err.Timestamp.Format("15:04:05.000"))
	}
}
