// Корчагин Евгений 363
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// структура для хранения метрик
type Metrics struct {
	mu                sync.RWMutex
	successCount      int
	errorCount        int
	totalResponseTime time.Duration
	requestCount      int
}

//  создаем новый объект метрик , почему метрики ? а не матрики или еще ччто то

func NewMetrics() *Metrics {
	return &Metrics{}
}

// записываем успешный запрос
func (m *Metrics) RecordSuccess(responseTime time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.successCount++
	m.requestCount++
	m.totalResponseTime += responseTime
}

// RecordError записывает неуспешный запрос
func (m *Metrics) RecordError(responseTime time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.errorCount++
	m.requestCount++
	m.totalResponseTime += responseTime
}

// возвращает красивый отчет , я бы сказал даже лучший отчет
func (m *Metrics) GetReport() string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.requestCount == 0 {
		return "Нет данных для отчета"
	}

	avgResponseTime := m.totalResponseTime / time.Duration(m.requestCount)
	successRate := float64(m.successCount) / float64(m.requestCount) * 100

	return fmt.Sprintf(`
=== ОТЧЕТ ПО МЕТРИКАМ ===
Всего запросов: %d
Успешных: %d
Ошибок: %d
Среднее время ответа: %v
Процент успеха: %.2f%%
========================
`, m.requestCount, m.successCount, m.errorCount, avgResponseTime, successRate)
}

func main() {
	metrics := NewMetrics()
	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	// Запускаем горутины, имитирующие запросы
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				// Имитация времени ответа
				responseTime := time.Duration(rand.Intn(100)+50) * time.Millisecond

				// Случайным образом определяем успех или ошибку
				if rand.Float32() < 0.8 { // 80% успешных запросов
					metrics.RecordSuccess(responseTime)
					fmt.Printf("Запрос %d-%d: УСПЕХ (%v)\n", id, j, responseTime)
				} else {
					metrics.RecordError(responseTime)
					fmt.Printf("Запрос %d-%d: ОШИБКА (%v)\n", id, j, responseTime)
				}

				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println(metrics.GetReport())
}
