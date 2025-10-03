// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"time"
)

type Event struct {
	ID           int
	Title        string
	Date         time.Time
	Location     string
	MaxAttendees int
	Attendees    []string
}

type EventManager struct {
	events map[int]*Event
	nextID int
}

func NewEventManager() *EventManager {
	return &EventManager{
		events: make(map[int]*Event),
		nextID: 1,
	}
}

func (em *EventManager) CreateEvent(title, location string, date time.Time, maxAttendees int) *Event {
	event := &Event{
		ID:           em.nextID,
		Title:        title,
		Date:         date,
		Location:     location,
		MaxAttendees: maxAttendees,
	}
	em.events[em.nextID] = event
	em.nextID++
	return event
}

func (em *EventManager) Register(eventID int, name string) error {
	event, exists := em.events[eventID]
	if !exists {
		return errors.New("мероприятие не найдено")
	}
	if len(event.Attendees) >= event.MaxAttendees {
		return errors.New("нет мест")
	}
	event.Attendees = append(event.Attendees, name)
	return nil
}

func (em *EventManager) Cancel(eventID int, name string) error {
	event, exists := em.events[eventID]
	if !exists {
		return errors.New("мероприятие не найдено")
	}
	for i, a := range event.Attendees {
		if a == name {
			event.Attendees = append(event.Attendees[:i], event.Attendees[i+1:]...)
			return nil
		}
	}
	return errors.New("участник не найден")
}

func (em *EventManager) GetUpcoming() []*Event {
	var result []*Event
	now := time.Now()
	for _, event := range em.events {
		if event.Date.After(now) {
			result = append(result, event)
		}
	}
	return result
}

func main() {
	em := NewEventManager()

	// Создаем мероприятия
	concert := em.CreateEvent("Стрим", "Твич ' '", time.Now().Add(24*time.Hour), 100)
	exhibition := em.CreateEvent("Концерт", "Площадь 'музыка'", time.Now().Add(48*time.Hour), 50)

	// Регистрация
	em.Register(concert.ID, "Анна")
	em.Register(concert.ID, "Иван")
	em.Register(exhibition.ID, "Мария")

	// Показываем предстоящие
	fmt.Println("Предстоящие мероприятия:")
	for _, event := range em.GetUpcoming() {
		fmt.Printf("- %s: %d/%d участников\n", event.Title, len(event.Attendees), event.MaxAttendees)
	}

	// Отмена
	em.Cancel(concert.ID, "Иван")

	// Ошибки
	fmt.Println("\nОшибки:")
	if err := em.Register(999, "Петр"); err != nil {
		fmt.Println(err)
	}
	if err := em.Register(concert.ID, "Анна"); err != nil {
		fmt.Println(err)
	}
}
