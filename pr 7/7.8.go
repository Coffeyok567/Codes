// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
	"time"
)

type Room struct {
	Number int
	Type   string
	Price  float64
}

type Hotel struct {
	Name  string
	Rooms map[int]*Room
}

type Reservation struct {
	ID      int
	Hotel   string
	Room    int
	Guest   string
	CheckIn time.Time
	Nights  int
	Total   float64
}

type BookingSystem struct {
	hotels       map[string]*Hotel
	reservations map[int]*Reservation
	nextID       int
}

func NewBookingSystem() *BookingSystem {
	return &BookingSystem{
		hotels:       make(map[string]*Hotel),
		reservations: make(map[int]*Reservation),
		nextID:       1,
	}
}

func (bs *BookingSystem) AddHotel(name string) {
	bs.hotels[name] = &Hotel{
		Name:  name,
		Rooms: make(map[int]*Room),
	}
}

func (bs *BookingSystem) AddRoom(hotel string, number int, roomType string, price float64) error {
	h, exists := bs.hotels[hotel]
	if !exists {
		return errors.New("отель не найден")
	}
	h.Rooms[number] = &Room{Number: number, Type: roomType, Price: price}
	return nil
}

func (bs *BookingSystem) CheckAvailability(hotel string, roomType string) []*Room {
	h, exists := bs.hotels[hotel]
	if !exists {
		return nil
	}
	var available []*Room
	for _, room := range h.Rooms {
		if room.Type == roomType {
			available = append(available, room)
		}
	}
	return available
}

func (bs *BookingSystem) BookRoom(hotel, guest string, roomNumber, nights int) (*Reservation, error) {
	h, exists := bs.hotels[hotel]
	if !exists {
		return nil, errors.New("отель не найден")
	}
	room, exists := h.Rooms[roomNumber]
	if !exists {
		return nil, errors.New("номер не найден")
	}
	reservation := &Reservation{
		ID:      bs.nextID,
		Hotel:   hotel,
		Room:    roomNumber,
		Guest:   guest,
		CheckIn: time.Now(),
		Nights:  nights,
		Total:   room.Price * float64(nights),
	}
	bs.reservations[bs.nextID] = reservation
	bs.nextID++
	return reservation, nil
}

func main() {
	system := NewBookingSystem()

	// Добавляем отели и номера
	system.AddHotel("Гранд Отель")
	system.AddRoom("Гранд Отель", 101, "стандарт", 5000)
	system.AddRoom("Гранд Отель", 102, "люкс", 10000)

	system.AddHotel("Плаза")
	system.AddRoom("Плаза", 201, "стандарт", 4500)
	system.AddRoom("Плаза", 202, "люкс", 9000)

	// Проверяем доступность
	fmt.Println("Доступные стандартные номера в Гранд Отеле:")
	for _, room := range system.CheckAvailability("Гранд Отель", "стандарт") {
		fmt.Printf("- номер %d: %.2f руб/ночь\n", room.Number, room.Price)
	}

	// Бронируем номер
	res, err := system.BookRoom("Гранд Отель", "Иван Иванов", 101, 3)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("\nБронь #%d создана: %s, номер %d, %d ночей, итого: %.2f руб\n",
			res.ID, res.Guest, res.Room, res.Nights, res.Total)
	}

	// Ошибки
	_, err = system.BookRoom("Несуществующий", "Петр", 101, 2)
	fmt.Println("\nОшибка при брони:", err)
}
