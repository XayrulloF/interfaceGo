package main

import (
	"fmt"
)

type iFace interface {
	Area() float64
	Perimeter() float64
}

// Shape struct for geometrical shapes
type Shape struct {
	width  float64
	height float64
}

// Area calculating
func (u *Shape) Area() float64 {
	return u.height * u.width
}

// Perimeter calculate
func (u *Shape) Perimeter() float64 {
	return (u.height + u.width) * 2
}
func main() {
	var first Shape
	fmt.Scanln(&first.width)
	fmt.Scanln(&first.height)
	//first := Shape{12, 10.4}
	fmt.Printf("Area: %.2f\n", first.Area())
	fmt.Printf("Perimeter: %.2f", first.Perimeter())
}
