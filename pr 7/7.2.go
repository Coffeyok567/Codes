// Корчагин Евгений 363
package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

type Inventory struct {
	products map[int]*Product
}

var inv Inventory

func AddProduct(product *Product) {
	if inv.products == nil {
		inv.products = make(map[int]*Product)
	}
	fmt.Printf("%s добавлен в магазин", product.Name)
	inv.products[product.ID] = product
}

func WriteOff(productID int, quantity int) error {
	value, ok := inv.products[productID]
	if ok {
		fmt.Println(quantity, value.Name+"было списанно со склада")
		value.Quantity -= quantity
		return nil
	} else {
		return errors.New("нету")
	}
}

func RemoveProduct(productID int) error {
	value, ok := inv.products[productID]
	if ok {
		fmt.Println(value.Name + "был уничтожен")
		delete(inv.products, productID)
		return nil
	} else {
		return errors.New("нету")
	}
}

func GetTotalValue() float64 {
	var sum float64
	for product := range inv.products {
		sum += float64(inv.products[product].Quantity) * inv.products[product].Price
	}
	return sum
}

func main() {
	var vandal Product
	vandal.ID = 1
	vandal.Name = "Vandal"
	vandal.Price = 2175.0
	vandal.Quantity = 2
	var phantom Product
	phantom.ID = 2
	phantom.Name = "phantom"
	phantom.Price = 1775.0
	phantom.Quantity = 2
	AddProduct(&vandal)
	AddProduct(&phantom)
	RemoveProduct(vandal.ID)
	AddProduct(&vandal)
	WriteOff(vandal.ID, 1)
	AddProduct(&phantom)
	fmt.Println("сумма всего на складе:", GetTotalValue())
}
