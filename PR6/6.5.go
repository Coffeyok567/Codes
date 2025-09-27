//Корчагин Евгений 363

package main

import (
	"fmt"
	"sync"
	"time"
)

// CacheItem представляет элемент кэша
type CacheItem struct {
	Value      interface{}
	Expiration time.Time
}

// CacheWithTTL кэш с временем жизни записей
type CacheWithTTL struct {
	mu    sync.RWMutex
	items map[string]CacheItem
	ttl   time.Duration
}

// NewCacheWithTTL создает новый кэш с TTL
func NewCacheWithTTL(ttl time.Duration) *CacheWithTTL {
	cache := &CacheWithTTL{
		items: make(map[string]CacheItem),
		ttl:   ttl,
	}

	// Запускаем фоновую горутину для очистки устаревших записей
	go cache.cleanupWorker()

	return cache
}

// Set добавляет или обновляет значение в кэше
func (c *CacheWithTTL) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = CacheItem{
		Value:      value,
		Expiration: time.Now().Add(c.ttl),
	}

	fmt.Printf("Добавлено в кэш: %s -> %v (TTL: %v)\n", key, value, c.ttl)
}

// Get получает значение из кэша (не обновляет TTL)
func (c *CacheWithTTL) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	// Проверяем, не истекло ли время жизни
	if time.Now().After(item.Expiration) {
		return nil, false
	}

	return item.Value, true
}

// Delete удаляет элемент из кэша
func (c *CacheWithTTL) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.items, key)
	fmt.Printf("Удалено из кэша: %s\n", key)
}

// cleanupWorker периодически очищает устаревшие записи
func (c *CacheWithTTL) cleanupWorker() {
	ticker := time.NewTicker(c.ttl / 2) // Очистка каждые TTL/2
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		now := time.Now()
		count := 0

		for key, item := range c.items {
			if now.After(item.Expiration) {
				delete(c.items, key)
				count++
			}
		}

		if count > 0 {
			fmt.Printf("Автоочистка: удалено %d устаревших записей\n", count)
		}
		c.mu.Unlock()
	}
}

// GetStats возвращает статистику кэша
func (c *CacheWithTTL) GetStats() (int, int) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	total := len(c.items)
	valid := 0
	now := time.Now()

	for _, item := range c.items {
		if now.Before(item.Expiration) {
			valid++
		}
	}

	return total, valid
}

// PrintStatus печатает текущее состояние кэша
func (c *CacheWithTTL) PrintStatus() {
	c.mu.RLock()
	defer c.mu.RUnlock()

	now := time.Now()
	fmt.Println("\n=== Состояние кэша ===")
	fmt.Printf("Всего записей: %d\n", len(c.items))

	for key, item := range c.items {
		remaining := item.Expiration.Sub(now)
		if remaining > 0 {
			fmt.Printf("  %s: %v (осталось: %v)\n", key, item.Value, remaining)
		} else {
			fmt.Printf("  %s: %v (ИСТЕК)\n", key, item.Value)
		}
	}
	fmt.Println("=====================")
}

func main() {
	// Создаем кэш с TTL = 3 секунды
	cache := NewCacheWithTTL(3 * time.Second)

	// Добавляем несколько значений
	cache.Set("user:1", "Алексей")
	cache.Set("user:2", "Мария")
	cache.Set("config:timeout", 30)

	// Проверяем сразу после добавления
	fmt.Println("\n--- Проверка сразу после добавления ---")
	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Найден user:1 -> %v\n", val)
	}

	// Ждем 2 секунды и проверяем снова
	time.Sleep(2 * time.Second)
	fmt.Println("\n--- Проверка через 2 секунды ---")
	cache.PrintStatus()

	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Найден user:1 -> %v\n", val)
	} else {
		fmt.Println("user:1 не найден или истек")
	}

	// Ждем еще 2 секунды (всего 4 секунды - больше чем TTL)
	time.Sleep(2 * time.Second)
	fmt.Println("\n--- Проверка через 4 секунды ---")
	cache.PrintStatus()

	if val, ok := cache.Get("user:1"); ok {
		fmt.Printf("Найден user:1 -> %v\n", val)
	} else {
		fmt.Println("user:1 не найден или истек")
	}

	// Добавляем новое значение и проверяем
	cache.Set("user:3", "Иван")
	time.Sleep(1 * time.Second)
	fmt.Println("\n--- После добавления нового пользователя ---")
	cache.PrintStatus()

	//  статистикa
	total, valid := cache.GetStats()
	fmt.Printf("\nСтатистика: всего %d записей, валидных %d\n", total, valid)

	time.Sleep(5 * time.Second)
	fmt.Println("\n--- Финальное состояние ---")
	cache.PrintStatus()
}
