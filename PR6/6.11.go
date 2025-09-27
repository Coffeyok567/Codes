//Корчагин Евгений 363

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// кинотеатр с местами
type Cinema struct {
	mu    sync.RWMutex
	seats []bool // true - забронировано, false - свободно
}

// новый кинотеатр с 38 местами
func NewCinema() *Cinema {
	return &Cinema{
		seats: make([]bool, 38),
	}
}

// забронировать место
func (c *Cinema) BookSeat(seatNumber int, userID string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Проверяем валидность номера места
	if seatNumber < 0 || seatNumber >= len(c.seats) {
		fmt.Printf("Пользователь %s: место %d не существует\n", userID, seatNumber)
		return false
	}

	// Проверяем, свободно ли место
	if c.seats[seatNumber] {
		fmt.Printf("Пользователь %s: место %d уже занято\n", userID, seatNumber)
		return false
	}

	// Имитация времени обработки бронирования
	time.Sleep(50 * time.Millisecond)

	// Бронируем место
	c.seats[seatNumber] = true
	fmt.Printf("Пользователь %s: место %d успешно забронировано\n", userID, seatNumber)
	return true
}

// количество свободных мест
func (c *Cinema) GetAvailableSeats() int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	count := 0
	for _, booked := range c.seats {
		if !booked {
			count++
		}
	}
	return count
}

// отображает текущее состояние мест
func (c *Cinema) DisplaySeats() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	fmt.Println("\nТекущее состояние мест:")
	for i, booked := range c.seats {
		status := "Свободно"
		if booked {
			status = "Занято"
		}
		fmt.Printf("Место %2d: %s\n", i+1, status)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cinema := NewCinema()
	var wg sync.WaitGroup

	// Количество пользователей, пытающихся забронировать места
	numUsers := 50

	// Запускаем горутины пользователей
	for i := 0; i < numUsers; i++ {
		wg.Add(1)
		go func(userID int) {
			defer wg.Done()

			// Каждый пользователь пытается забронировать 1-3 места
			attempts := rand.Intn(3) + 1
			for j := 0; j < attempts; j++ {
				// Случайный выбор места
				seat := rand.Intn(38)
				cinema.BookSeat(seat, fmt.Sprintf("User%d", userID))

				// Небольшая задержка между попытками
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}(i)
	}

	// Горутина для отображения прогресса
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			time.Sleep(200 * time.Millisecond)
			available := cinema.GetAvailableSeats()
			fmt.Printf("Осталось свободных мест: %d/%d\n", available, 38)
		}
	}()

	wg.Wait()

	// Финальный отчет
	fmt.Println("\n===   ОТЧЕТ ===")
	cinema.DisplaySeats()
	available := cinema.GetAvailableSeats()
	fmt.Printf("Всего забронировано мест: %d\n", 38-available)
	fmt.Printf("Свободных мест: %d\n", available)
}
