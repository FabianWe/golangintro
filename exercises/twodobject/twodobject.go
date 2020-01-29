package main

import (
	"fmt"
	"math"
)

type TwoDObject interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

func NewRectangle(width, height float64) *Rectangle {
	return &Rectangle{Width: width, Height: height}
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2.0 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{Radius: radius}
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.Radius
}

func PrintInfo(object TwoDObject) {
	fmt.Printf("TwoDObject: Area = %.3f, Perimeter = %.3f\n",
		object.Area(), object.Perimeter())
}

func main() {
	r := NewRectangle(5, 10)
	c := NewCircle(5)
	PrintInfo(r)
	PrintInfo(c)
	PrintInfo(t)
}
