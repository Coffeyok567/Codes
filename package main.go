package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	RoomTypeSingle = "Одноместный"
	RoomTypeDouble = "Двухместнй"
	RoomTypeSuite  = "Люкс"

	RoomStatusFree        = "Свободен"
	RoomStatusBooked      = "Забронирован"
	RoomStatusMaintenance = "Обслуживается"

	bin = 2
	dec = 10
	hex = 16
)

type SensorData struct {
	SensorID    string
	Temperature float64
	Humidity    float64
	Timestamp   time.Time
}

type Movie struct {
	Title  string
	Year   int
	Rating float64
	Genres []string
}

type InventoryItem struct {
	Name        string
	Weight      float64
	IsQuestItem bool
}

type Product struct {
	Name     string
	Category string
	Price    float64
}

type HotelRoom struct {
	Type   string
	Status string
	Price  float64
}

type LogEntry struct {
	IP        string
	HTTPCode  int
	Timestamp time.Time
}

type Employee struct {
	ID       int
	Name     string
	Position string
	Salary   float64
}

type Order struct {
	ID           int
	Items        []int
	Allsum       []int
	Total        float64
	Adress       string
	isCompleated bool
}

func calculateAverageTemperature(data []SensorData) float64 {
	if len(data) == 0 {
		return 0
	}

	total := 0.0
	for _, reading := range data {
		total += reading.Temperature
	}

	return total / float64(len(data))
}

func FindHighestRating(Movies []Movie) (float64, string) {
	if len(Movies) == 0 {
		return 0, ""
	}

	highestRating := Movies[0].Rating
	movieTitle := Movies[0].Title

	for _, movie := range Movies {
		if movie.Rating > highestRating {
			highestRating = movie.Rating
			movieTitle = movie.Title
		}
	}

	return highestRating, movieTitle
}

func SumWeight(inventory []InventoryItem) (SumWeight float64) {
	for _, inventoryweight := range inventory {
		SumWeight += inventoryweight.Weight
	}
	return SumWeight
}

func convertNumber(numberStr string, fromBase, toBase int) (string, error) {
	// Преобразуем строку в число с указанной исходной системой счисления
	num, err := strconv.ParseInt(numberStr, fromBase, 64)
	if err != nil {
		return "", fmt.Errorf("ошибка преобразования: %v", err)
	}

	// Преобразует число в строку в систему счисления котору. указали при вызове функции
	switch toBase {
	case 2:
		return strconv.FormatInt(num, bin), nil
	case 10:
		return strconv.FormatInt(num, dec), nil
	case 16:
		return strconv.FormatInt(num, hex), nil
	default:
		return "", fmt.Errorf("неподдерживаемая система счисления: %d", toBase)
	}
}

func filterProducts(products []Product, maxPrice float64, category string) []Product {
	// объявил пустой срез который ссылается на срез (Вроде)
	var filtered []Product

	//цыкл for  для того что бы перебрать все строки в срезе products
	for _, product := range products {

		if product.Price < maxPrice && product.Category == category {
			filtered = append(filtered, product) //тут у нас записываются отфильтрованыне значения в новый срез(массив) из которого будем брать отфильтрованные значения
		}
	}
	return filtered
}

func textStats(text string) {

	symbolcount := len(text)

	words := strings.Fields(text)
	wordcount := len(words)

	ExclamationMark := strings.Count(text, "!")
	Point := strings.Count(text, ".")
	QuastionMark := strings.Count(text, "?")

	fmt.Println("строка: ", text)
	fmt.Println("Количество символов в строке: ", symbolcount)
	fmt.Println("Количество слов в строке:", wordcount)
	fmt.Println("Количество предложений в строке: ", ExclamationMark+Point+QuastionMark)
}

func bookRoom(rooms map[string]HotelRoom, roomNumber string) error {
	room, hotelroom := rooms[roomNumber]
	if !hotelroom {
		return fmt.Errorf("комната %s не существует", roomNumber)
	}

	if room.Status != RoomStatusFree {
		return fmt.Errorf("комната %s недоступна для бронирования. Текущий статус: %s", roomNumber, room.Status)
	}

	room.Status = RoomStatusFree
	rooms[roomNumber] = room
	fmt.Printf("Комната %s успешно забронирована!\n", roomNumber)
	return nil
}

func ValidUser(name string, age int, email string) error {

	if name == "" || len(name) >= 50 {
		return errors.New("invalid name: Имя не может быть пустым или больше 50 символов")
	}

	if age < 18 || age > 120 {
		return errors.New("invalid age: Возраст не подходит")
	}

	if strings.Index(email, "@") == -1 {
		return errors.New("invalid mail: Почта не корректна")
	}

	return nil

}

func filterErrorLogs(entries []LogEntry) []LogEntry {
	var errorEntries []LogEntry

	for _, entry := range entries {
		// это накрутейшая проверка на ошибки
		if (entry.HTTPCode >= 400 && entry.HTTPCode <= 499) ||
			(entry.HTTPCode >= 500 && entry.HTTPCode <= 599) {
			errorEntries = append(errorEntries, entry)
		}
	}

	return errorEntries
}

func SalaryСalculation(employees []Employee) (TotalSalary float64, AvgSalary float64) {
	for _, employee := range employees {
		TotalSalary += employee.Salary
	}

	AvgSalary = TotalSalary / float64(len(employees))

	return TotalSalary, AvgSalary
}

func CountVote(Votes []string) {
	var VoteBoris int
	var VoteAnna int
	var VoteVictor int
	var AllVotes float64
	var j string
	for _, j = range Votes {
		if j == "Борис" {
			VoteBoris++
		}
		if j == "Анна" {
			VoteAnna++
		}
		if j == "Виктор" {
			VoteVictor++
		}
	}
	fmt.Printf("Проголосовли за Анну: %d\n", VoteAnna)
	fmt.Printf("Проголосовли за Бориса: %d\n", VoteBoris)
	fmt.Printf("Проголосовли за Виктор: %d\n", VoteVictor)
	AllVotes = float64(VoteBoris+VoteAnna+VoteVictor) / 100
	fmt.Printf("Процент проголосовавших за Анну: %F\n", AllVotes*float64(VoteAnna))
	fmt.Printf("Процент проголосовавших за Бориса: %F\n", AllVotes*float64(VoteBoris))
	fmt.Printf("Процент проголосовавших за Виктора: %F\n", AllVotes*float64(VoteVictor))
}

func NewOrder(order *map[int]Order) {
	var LastKey int
	var Items []int
	var Total float64
	var Adress string
	for k := range *order {
		if k > LastKey {
			LastKey = k + 1
		}
	}
	for {
		fmt.Println("Введите id предмета ")
		var Item int
		fmt.Scan(&Item)
		if Item == -1 {
			break
		} else if Item <= 0 {
			fmt.Println("Нет такого id")
		} else {
			Items = append(Items, Item)
		}
	}
	for {
		fmt.Println("введите цену ")
		fmt.Scan(&Total)
		if Total > 0 {
			break
		}
	}
	for {
		fmt.Println("введите адрес ")
		fmt.Scan(&Adress)
		if Adress != "" {
			break
		}
	}

	OrderValue := Order{
		ID:           LastKey,
		Items:        Items,
		Total:        Total,
		Adress:       Adress,
		isCompleated: false,
	}

	(*order)[LastKey] = OrderValue
	fmt.Println(" Заказ добавили ")

}

func collectUniqueTags(posts [][]string) map[string]bool {
	uniqueTags := make(map[string]bool)

	for _, post := range posts {
		for _, tag := range post {
			uniqueTags[tag] = true
		}
	}

	return uniqueTags
}

func main() {
	//Помогите я устал(   я сделал только 7 задач а мне еще столько узнать как что делать

	fmt.Println("Введите номер задания")
	var TaskNumber int
	fmt.Scan(&TaskNumber)
	switch TaskNumber {

	case 1:
		PriceDay := map[string]int{
			"ПН": 2100,
			"ВТ": 2100,
			"СР": 2100,
			"ЧТ": 2100,
			"ПТ": 2850,
			"СБ": 2850,
			"ВС": 2850,
		}
		price := PriceDay["ВТ"] + PriceDay["СР"] + PriceDay["ЧТ"] + PriceDay["ПТ"] + PriceDay["СБ"] + PriceDay["ВС"] + PriceDay["ЧТ"] + PriceDay["ПТ"]
		fmt.Println(price)

	case 2:
		var MainWaight float64
		var HandWaight float64
		var BonusWaight float64
		fmt.Println("введите вес основного багажа")
		fmt.Scan(&MainWaight)

		fmt.Println("введите вес ручной клади")
		fmt.Scan(&HandWaight)
		fmt.Println("введите вес дополнительной  клади")
		fmt.Scan(&BonusWaight)

		if MainWaight < 0 || HandWaight < 0 || BonusWaight < 0 {
			fmt.Println("вес не может быть отрицательным")
		} else if MainWaight >= 0 || HandWaight >= 0 || BonusWaight >= 0 {
			sum := MainWaight + HandWaight + BonusWaight
			fmt.Println("Общий вес")
			fmt.Println(sum)
		}

	case 3:
		Orders := map[int]Order{
			1: {
				ID:           1,
				Items:        []int{2, 3, 4},
				Total:        200.5,
				Adress:       "зазик",
				isCompleated: false,
			},
		}
		fmt.Println(Orders)

	case 4:
		var Condidat = []string{"Анна", "Борис", "Виктор", "Анна", "Борис", "Виктор", "Анна", "Борис", "Виктор", "Анна"}
		CountVote(Condidat)

	case 5:
		var name string
		var age int
		var email string

		fmt.Print("Введите имя:")
		fmt.Scan(&name)
		fmt.Print("Введите возраст:")
		fmt.Scan(&age)
		fmt.Print("Введите почту:")
		fmt.Scan(&email)

		user := ValidUser(name, age, email)

		if user != nil {
			fmt.Println("Ошибка:", user)

		} else {
			fmt.Println("Все данные корректны!")
			fmt.Println("Имя:", name)
			fmt.Println("Возраст:", age)
			fmt.Println("Email:", email)
		}

	case 6:
		posts := [][]string{
			{"go", "backend"},
			{"git", "go", "tools"},
			{"backend", "database", "go"},
			{"tools", "git", "version-control"},
			{"web", "frontend", "javascript"},
		}

		uniqueTags := collectUniqueTags(posts)

		fmt.Println("Все уникальные теги:")
		for tag := range uniqueTags {
			fmt.Println("-", tag)
		}

	case 7:
		employees := []Employee{
			{ID: 1, Name: "kiyoshi", Position: "Разработчик", Salary: 100000},
			{ID: 2, Name: "Yuki", Position: "Дизайнер", Salary: 100000},
			{ID: 3, Name: "Takashi", Position: "Тестировщик", Salary: 90000},
		}

		TotalSalary, AvgSalary := SalaryСalculation(employees)

		fmt.Print("Общий фонд: ")
		fmt.Println(TotalSalary)
		fmt.Print("Средняя зп: ")
		fmt.Println(AvgSalary)

	case 8:
		logEntries := []LogEntry{
			{IP: "192.168.1.1", HTTPCode: 200, Timestamp: time.Now()},
			{IP: "192.168.1.2", HTTPCode: 404, Timestamp: time.Now().Add(-time.Minute * 5)},
			{IP: "192.168.1.3", HTTPCode: 500, Timestamp: time.Now().Add(-time.Minute * 10)},
			{IP: "192.168.1.4", HTTPCode: 301, Timestamp: time.Now().Add(-time.Minute * 15)},
			{IP: "192.168.1.5", HTTPCode: 403, Timestamp: time.Now().Add(-time.Minute * 20)},
			{IP: "192.168.1.6", HTTPCode: 200, Timestamp: time.Now().Add(-time.Minute * 25)},
			{IP: "192.168.1.7", HTTPCode: 503, Timestamp: time.Now().Add(-time.Minute * 30)},
		}

		errorLogs := filterErrorLogs(logEntries)
		fmt.Println("Все записи лога:")
		for i, entry := range logEntries {
			fmt.Printf("%d. IP: %s, Код: %d, Время: %s\n",
				i+1, entry.IP, entry.HTTPCode, entry.Timestamp.Format("2006-01-02 15:04:05"))
		}
		// Выводим  записи с ошибками    я устал
		fmt.Println("\nЗаписи с ошибками (4xx и 5xx):")
		for i, entry := range errorLogs {
			fmt.Printf("%d. IP: %s, Код: %d, Время: %s\n",
				i+1, entry.IP, entry.HTTPCode, entry.Timestamp.Format("2006-01-02 15:04:05"))
		}

	case 9:
		rooms := map[string]HotelRoom{
			"101": {Type: RoomTypeSingle, Status: RoomStatusMaintenance, Price: 2500},
			"102": {Type: RoomTypeSingle, Status: RoomStatusFree, Price: 2500},
			"103": {Type: RoomTypeSuite, Status: RoomStatusBooked, Price: 4000},
		}

		fmt.Println("\nИнформация о номерах отеля:")
		for number, room := range rooms {
			fmt.Printf("Номер: %s, Тип: %s, Статус: %s, Цена за ночь: %.2f руб.\n",
				number, room.Type, room.Status, room.Price)
		}

		fmt.Println("\nПопытка бронирования:")

		// Успешное бронирование
		if err := bookRoom(rooms, "102"); err != nil {
			fmt.Println("Ошибка:", err)
		}

		// Попытка забронировать уже занятый номер
		if err := bookRoom(rooms, "201"); err != nil {
			fmt.Println("Ошибка:", err)
		}

		// Попытка забронировать номер на обслуживании
		if err := bookRoom(rooms, "101"); err != nil {
			fmt.Println("Ошибка:", err)
		}

	case 10:
		text := "Hello ,world"
		textStats(text)

	case 11:

		// сделали срез который ссылается на тип ( структуру ) Product
		products := []Product{
			{Name: "Телефон ", Category: "Техника", Price: 20000},
			{Name: "Футболка ", Category: "Одежда", Price: 1500},
			{Name: "Худи ", Category: "Одежда", Price: 2000},
		}

		//объявил переменную product в который вызвал функцию filterProducts(products []Product, maxPrice float64, category string)
		// которая принимает в значения (срез, макс цену , категорию)

		product := filterProducts(products, 25000, "Техника")

		fmt.Println("Отфильтрованные товары:")
		for i, product := range product {
			fmt.Printf("%d Название: %s  Категория: %s  Цена: %.2f", i+1, product.Name, product.Category, product.Price)
		}

	case 12:

		var input string
		//  Принимает десятичное число
		fmt.Print("Введите десятичное число: ")
		fmt.Scan(&input)

		// Преобразуем введенную строку в число
		num, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			fmt.Printf("Ошибка: введено недопустимое число: %v\n", err)
			return
		}

		// в двоичной системе
		binary := strconv.FormatInt(num, bin)
		fmt.Printf("Двоичное представление: %s\n", binary)

		//в шестнадцатеричной системе
		hexadecimal := strconv.FormatInt(num, hex)
		fmt.Printf("Шестнадцатеричное представление: %s\n", hexadecimal)

		// Конвертируем двоичное в десятичное
		decFromBin, err := convertNumber(binary, bin, dec)
		if err != nil {
			fmt.Printf("Ошибка конвертации: %v\n", err)
		} else {
			fmt.Printf("Двоичное %s -> Десятичное %s\n", binary, decFromBin)
		}

	case 13:
		spending := make(map[string]float64)

		spending["Еда"] = 15000
		spending["Транспорт"] = 5000
		spending["Развлечения"] = 3000

		fmt.Println("Начальные траты:")
		for category, amount := range spending {
			fmt.Printf("Траты по категории: %s  составляют: %.2f\n", category, amount)
		}

		fmt.Println("Финальные траты:")
		spending["Еда"] += 2000
		for category, amount := range spending {
			fmt.Printf("Финальные траты по категории: %s составляют: %.2f\n", category, amount)
		}

	case 14:
		inventory := []InventoryItem{
			{Name: "Sword", Weight: 3, IsQuestItem: false},
			{Name: "HealthPoition", Weight: 0.3, IsQuestItem: false},
			{Name: "key", Weight: 0.1, IsQuestItem: true},
			{Name: "bow", Weight: 1, IsQuestItem: false},
			{Name: "ring", Weight: 0.1, IsQuestItem: true},
		}

		weight := SumWeight(inventory)
		fmt.Printf("Суммарный вес инвентаря игрока: %.2f", weight)

	case 15:
		// все данные я брал с кинопоиска  и с  некоторых сайтов с аниме
		Movies := []Movie{
			{Title: "1+1", Year: 2011, Rating: 8.3, Genres: []string{"Драма"}},
			{Title: "Волк с Уолл-стрит", Year: 2013, Rating: 8.1, Genres: []string{"Комедия"}}, // там очень много ,  поэтому что то одно
			{Title: "Необъятный океан 2", Year: 2025, Rating: 9.6, Genres: []string{"Аниме"}},  // тоже очень много поэтому просто аниме
			{Title: "Человек-бензопила", Year: 1995, Rating: 8.4, Genres: []string{"Аниме"}},
			{Title: "Поднятие уровня в одиночку", Year: 2024, Rating: 9.3, Genres: []string{"Фэнтези"}},
		}

		HigestRating, MovieTitle := FindHighestRating(Movies)
		fmt.Printf("Фильм с самым высоким рейтингом: %s  рейтинг: %.2f:", MovieTitle, HigestRating)

	case 16:
		var sensorReadings []SensorData

		now := time.Now()
		sensorReadings = append(sensorReadings,
			SensorData{
				SensorID:    "sensor-001",
				Temperature: 23.5,
				Humidity:    45.0,
				Timestamp:   now.Add(-time.Hour * 6),
			},
			SensorData{
				SensorID:    "sensor-001",
				Temperature: 24.1,
				Humidity:    43.5,
				Timestamp:   now.Add(-time.Hour * 4),
			},
			SensorData{
				SensorID:    "sensor-001",
				Temperature: 25.3,
				Humidity:    42.0,
				Timestamp:   now.Add(-time.Hour * 2),
			},
			SensorData{
				SensorID:    "sensor-001",
				Temperature: 26.0,
				Humidity:    40.5,
				Timestamp:   now,
			},
			SensorData{
				SensorID:    "sensor-001",
				Temperature: 24.8,
				Humidity:    41.5,
				Timestamp:   now.Add(time.Hour * 2),
			},
		)

		fmt.Println("Показания датчиков за сутки:")
		for i, reading := range sensorReadings {
			fmt.Printf("%d. Датчик: %s, Температура: %.1f°C, Влажность: %.1f%%, Время: %s\n",
				i+1, reading.SensorID, reading.Temperature, reading.Humidity,
				reading.Timestamp.Format("2006-01-02 15:04:05"))
		}

		avgTemp := calculateAverageTemperature(sensorReadings)
		fmt.Printf("\nСредняя температура за сутки: %.2f°C\n", avgTemp)

	default:
		fmt.Println("Такого задания нет")
	}
}
