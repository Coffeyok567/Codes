// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"sync"
)

type EventBus struct {
	subscribers map[string][]func(interface{})
	mu          sync.RWMutex
}

type Subscription struct {
	Event   string
	Handler func(interface{})
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]func(interface{})),
	}
}

func (eb *EventBus) Subscribe(event string, handler func(interface{})) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if eb.subscribers == nil {
		eb.subscribers = make(map[string][]func(interface{}))
	}

	eb.subscribers[event] = append(eb.subscribers[event], handler)
	fmt.Printf("Добавлен подписчик на событие: '%s'. Всего подписчиков: %d\n",
		event, len(eb.subscribers[event]))
}

func (eb *EventBus) Publish(event string, data interface{}) error {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	handlers, exists := eb.subscribers[event]
	if !exists || len(handlers) == 0 {
		return errors.New("нет подписчиков для данного события")
	}

	fmt.Printf("Публикация события: '%s' с данными: %v\n", event, data)
	fmt.Printf("Уведомление %d подписчиков...\n", len(handlers)) // в тгк

	for i, handler := range handlers {
		fmt.Printf("  [%d] ", i+1)
		handler(data)
	}

	fmt.Printf("Событие '%s' успешно обработано всеми подписчиками\n", event) // съели
	return nil
}

func (eb *EventBus) Unsubscribe(event string, handler func(interface{})) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	handlers, exists := eb.subscribers[event]
	if !exists {
		return errors.New("событие не найдено")
	}

	//встал вышел
	for i, h := range handlers {
		if fmt.Sprintf("%p", h) == fmt.Sprintf("%p", handler) {
			eb.subscribers[event] = append(handlers[:i], handlers[i+1:]...)
			fmt.Printf("Удален подписчик из события: '%s'. Осталось подписчиков: %d\n",
				event, len(eb.subscribers[event]))
			return nil
		}
	}

	return errors.New("обработчик не найден")
}

func (eb *EventBus) GetSubscribersCount(event string) int {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	return len(eb.subscribers[event])
}

func (eb *EventBus) GetEvents() []string {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	events := make([]string, 0, len(eb.subscribers))
	for event := range eb.subscribers {
		events = append(events, event)
	}
	return events
}

// статистик еее
func (eb *EventBus) DisplayStats() {
	fmt.Printf("\n=== Статистика EventBus ===\n")

	events := eb.GetEvents()
	if len(events) == 0 {
		fmt.Println("Нет зарегистрированных событий")
	} else {
		fmt.Printf("Всего событий: %d\n", len(events))
		for _, event := range events {
			count := eb.GetSubscribersCount(event)
			fmt.Printf("  - '%s': %d подписчиков\n", event, count)
		}
	}
	fmt.Printf("===========================\n\n")
}

func main() {
	eventBus := NewEventBus()

	// Создаем обработчики для события "user.created"
	userCreatedHandler1 := func(data interface{}) {
		user, ok := data.(map[string]string)
		if ok {
			fmt.Printf("Обработчик 1: Создан пользователь %s (email: %s)\n",
				user["name"], user["email"])
		} else {
			fmt.Printf("Обработчик 1: Получены данные: %v\n", data)
		}
	}

	userCreatedHandler2 := func(data interface{}) {
		user, ok := data.(map[string]string)
		if ok {
			fmt.Printf("Обработчик 2: Отправка приветственного email для %s\n",
				user["email"])
		}
	}

	// Создаем обработчики для события "order.placed"
	orderHandler1 := func(data interface{}) {
		order, ok := data.(map[string]interface{})
		if ok {
			fmt.Printf("Обработчик заказов: Обработка заказа #%d на сумму %.2f ₽\n",
				order["id"], order["amount"])
		}
	}

	orderHandler2 := func(data interface{}) {
		order, ok := data.(map[string]interface{})
		if ok {
			fmt.Printf("Склад: Резервирование товаров для заказа #%d\n", order["id"])
		}
	}

	// Подписываем обработчики на события
	fmt.Println("--- Регистрация подписчиков ---")
	eventBus.Subscribe("user.created", userCreatedHandler1)
	eventBus.Subscribe("user.created", userCreatedHandler2)
	eventBus.Subscribe("order.placed", orderHandler1)
	eventBus.Subscribe("order.placed", orderHandler2)
	eventBus.Subscribe("system.alert", func(data interface{}) {
		fmt.Printf("СИСТЕМНЫЙ АЛЕРТ: %v\n", data)
	})

	eventBus.DisplayStats()

	// Публикуем события
	fmt.Println("--- Публикация событий ---")

	// Событие user.created
	userData := map[string]string{
		"name":  "Иван Иванов",
		"email": "ivan@example.com",
	}
	err := eventBus.Publish("user.created", userData)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println()

	// Событие order.placed
	orderData := map[string]interface{}{
		"id":     1001,
		"amount": 15000.50,
		"items":  []string{"Ноутбук", "Мышь"},
	}
	err = eventBus.Publish("order.placed", orderData)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println()

	// Событие system.alert
	err = eventBus.Publish("system.alert", "Обнаружена подозрительная активность")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println()

	// Демонстрация отписки
	fmt.Println("--- Демонстрация отписки ---")
	err = eventBus.Unsubscribe("user.created", userCreatedHandler1)
	if err != nil {
		fmt.Printf("Ошибка отписки: %v\n", err)
	}

	eventBus.DisplayStats()

	// Публикуем событие после отписки
	fmt.Println("--- Публикация после отписки ---")
	err = eventBus.Publish("user.created", map[string]string{
		"name":  "Петр Петров",
		"email": "petr@example.com",
	})
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	fmt.Println()

	// Демонстрация обработки ошибок
	fmt.Println("--- Демонстрация обработки ошибок ---")
	err = eventBus.Publish("nonexistent.event", "данные")
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	err = eventBus.Unsubscribe("nonexistent.event", userCreatedHandler1)
	if err != nil {
		fmt.Printf("Ошибка отписки: %v\n", err)
	}
}
