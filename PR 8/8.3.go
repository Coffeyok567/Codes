// Корчагин Евгений 363
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func main() {
	shapes := []Shape{
		Rectangle{Width: 5, Height: 3},
		Circle{Radius: 4},
		Triangle{A: 3, B: 4, C: 5},
	}

	for i, shape := range shapes {
		fmt.Printf("Фигура %d:\n", i+1)
		fmt.Printf("  Площадь: %.2f\n", shape.Area())
		fmt.Printf("  Периметр: %.2f\n", shape.Perimeter())
		fmt.Println()
	}
}