// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"time"
)

type Cache struct {
	items map[string]*CacheItem
}

type CacheItem struct {
	value      interface{}
	expiration int64
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]*CacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	expiration := time.Now().Add(ttl).UnixNano()
	c.items[key] = &CacheItem{
		value:      value,
		expiration: expiration,
	}
	fmt.Printf("Добавлено в кэш: ключ '%s', значение '%v', TTL: %v\n", key, value, ttl)
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if time.Now().UnixNano() > item.expiration {
		delete(c.items, key)
		fmt.Printf("Запись с ключом '%s' устарела и удалена\n", key)
		return nil, false
	}

	fmt.Printf("Получено из кэша: ключ '%s', значение '%v'\n", key, item.value)
	return item.value, true
}

// туда их
func (c *Cache) Delete(key string) error {
	_, exists := c.items[key]
	if !exists {
		return errors.New("ключ не найден в кэше")
	}

	delete(c.items, key)
	fmt.Printf("Запись с ключом '%s' удалена из кэша\n", key)
	return nil
}

// уборка
func (c *Cache) CleanupExpired() {
	now := time.Now().UnixNano()
	count := 0
	for key, item := range c.items {
		if now > item.expiration {
			delete(c.items, key)
			count++
		}
	}
	fmt.Printf("Очистка кэша: удалено %d устаревших записей\n", count)
}

func (c *Cache) Size() int {
	return len(c.items)
}

func (c *Cache) Display() {
	fmt.Printf("\n=== Состояние кэша ===\n")
	fmt.Printf("Всего записей: %d\n", c.Size())

	if len(c.items) == 0 {
		fmt.Println("Кэш пуст")
	} else {
		now := time.Now().UnixNano()
		for key, item := range c.items {
			remaining := time.Duration(item.expiration - now)
			status := "активна"
			if remaining <= 0 {
				status = "устарела"
			}
			fmt.Printf("  - %s: %v (TTL осталось: %v, статус: %s)\n",
				key, item.value, remaining, status)
		}
	}
	fmt.Printf("======================\n\n") // красиво
}

func main() {
	cache := NewCache()

	// Добавляем записи с разным TTL
	cache.Set("user:1", "JDH", 2*time.Second)
	cache.Set("user:2", "lololowka", 5*time.Second)
	cache.Set("config:timeout", 30, 10*time.Second)

	cache.Display()

	// Получаем существующие записи
	if value, found := cache.Get("user:1"); found {
		fmt.Printf("Найден пользователь: %v\n", value)
	}

	//   типо ошиблись
	if value, found := cache.Get("user:999"); !found {
		fmt.Println("Пользователь с ID 999 не найден в кэше")
	} else {
		fmt.Printf("Найден пользователь: %v\n", value)
	}

	// Удаляем запись
	err := cache.Delete("user:2")
	if err != nil {
		fmt.Printf("Ошибка при удалении: %v\n", err)
	}

	cache.Display()

	fmt.Println("Ждем 3 секунды...")
	time.Sleep(3 * time.Second)

	// Проверяем записи после ожидания                ждемс
	fmt.Println("\nПосле ожидания 3 секунды:")
	cache.Display()

	// user:1 должен был устареть     ну старый уже
	if value, found := cache.Get("user:1"); !found {
		fmt.Println("user:1 больше не в кэше (TTL истек)")
	} else {
		fmt.Printf("user:1 все еще в кэше: %v\n", value)
	}

	// Очищаем устаревшие записи  p.s  у а кому она нужна в бд да нужна а тут нет
	cache.CleanupExpired()
	cache.Display()

	// Демонстрация обработки ошибок
	fmt.Println("\n--- Демонстрация обработки ошибок ---")
	err = cache.Delete("nonexistent")
	if err != nil {
		fmt.Printf("Ошибка при удалении: %v\n", err)
	}
}
