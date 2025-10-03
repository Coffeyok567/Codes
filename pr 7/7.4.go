// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
)

type Order struct {
	Customer
	items  map[int]*OrderItem
	Status OrderStatus
}

type OrderItem struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

type Customer struct {
	ID   int
	Name string
}

type OrderStatus string

const (
	OrderStatusCreated   OrderStatus = "created"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusShipped   OrderStatus = "shipped"
	OrderStatusDelivered OrderStatus = "delivered"
	OrderStatusCancelled OrderStatus = "cancelled"
)

func NewOrder(customer Customer) *Order {
	return &Order{
		Customer: customer,
		items:    make(map[int]*OrderItem),
		Status:   OrderStatusCreated,
	}
}

func (o *Order) AddItem(item *OrderItem) {
	if o.items == nil {
		o.items = make(map[int]*OrderItem)
	}

	// Если товар уже есть в заказе, увеличиваем количество
	if existingItem, exists := o.items[item.ID]; exists {
		existingItem.Quantity += item.Quantity
	} else {
		o.items[item.ID] = item
	}
	fmt.Printf("Товар '%s' добавлен в заказ\n", item.Name)
}

func (o *Order) RemoveItem(itemID int) error {
	item, exists := o.items[itemID]
	if !exists {
		return errors.New("товар не найден в заказе")
	}

	delete(o.items, itemID)
	fmt.Printf("Товар '%s' удален из заказа\n", item.Name)
	return nil
}

func (o *Order) UpdateItemQuantity(itemID int, newQuantity int) error {
	item, exists := o.items[itemID]
	if !exists {
		return errors.New("товар не найден в заказе")
	}

	if newQuantity <= 0 {
		return o.RemoveItem(itemID)
	}

	item.Quantity = newQuantity
	fmt.Printf("Количество товара '%s' изменено на %d\n", item.Name, newQuantity)
	return nil
}

func (o *Order) UpdateStatus(newStatus OrderStatus) {
	oldStatus := o.Status
	o.Status = newStatus
	fmt.Printf("Статус заказа изменен: %s -> %s\n", o.GetStatusText(oldStatus), o.GetStatusText(newStatus))
}

func (o *Order) GetStatusText(status OrderStatus) string {
	statusMap := map[OrderStatus]string{
		OrderStatusCreated:   "Создан",
		OrderStatusPaid:      "Оплачен",
		OrderStatusShipped:   "Отправлен",
		OrderStatusDelivered: "Доставлен",
		OrderStatusCancelled: "Отменен",
	}
	return statusMap[status]
}

func (o *Order) GetCurrentStatus() string {
	return o.GetStatusText(o.Status)
}

func (o *Order) TotalCost() float64 {
	var total float64
	for _, item := range o.items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

func (o *Order) GetItems() []*OrderItem {
	items := make([]*OrderItem, 0, len(o.items))
	for _, item := range o.items {
		items = append(items, item)
	}
	return items
}

func (o *Order) DisplayOrderInfo() {
	fmt.Printf("\n=== Информация о заказе ===\n")
	fmt.Printf("Клиент: %s\n", o.Name)
	fmt.Printf("Статус: %s\n", o.GetCurrentStatus())
	fmt.Printf("Товары в заказе:\n")

	if len(o.items) == 0 {
		fmt.Println("  Заказ пуст")
	} else {
		for _, item := range o.items {
			fmt.Printf("  - %s: %d × %.2f ₽ = %.2f ₽\n",
				item.Name, item.Quantity, item.Price,
				float64(item.Quantity)*item.Price)
		}
	}
	fmt.Printf("Общая стоимость: %.2f ₽\n", o.TotalCost())
	fmt.Printf("===========================\n\n")
}

func NewOrderItem(id int, name string, quantity int, price float64) *OrderItem {
	return &OrderItem{
		ID:       id,
		Name:     name,
		Quantity: quantity,
		Price:    price,
	}
}

func main() {
	// Создаем клиента
	customer := Customer{ID: 1, Name: "Иван Иванов"}

	// Создаем заказ
	order := NewOrder(customer)
	fmt.Printf("Создан новый заказ для клиента: %s\n", customer.Name)

	// Добавляем товары в заказ
	item1 := NewOrderItem(1, "Ноутбук", 1, 50000.0)
	item2 := NewOrderItem(2, "Мышь", 2, 1500.0)
	item3 := NewOrderItem(3, "Клавиатура", 1, 3000.0)

	order.AddItem(item1)
	order.AddItem(item2)
	order.AddItem(item3)

	order.DisplayOrderInfo()

	order.UpdateStatus(OrderStatusPaid)

	order.UpdateItemQuantity(2, 3)

	order.RemoveItem(3)

	order.DisplayOrderInfo()

	order.UpdateStatus(OrderStatusShipped)
	order.UpdateStatus(OrderStatusDelivered)

	order.DisplayOrderInfo()

	fmt.Println("\n--- Демонстрация обработки ошибок ---")
	err := order.RemoveItem(999)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}

	err = order.UpdateItemQuantity(999, 5)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
	}
}
