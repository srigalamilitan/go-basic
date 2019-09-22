package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{X: 1, Y: 1, Radius: 5}
	retangle := Retangle{4, 6}
	fmt.Printf("Area Dari Circle : %f\n", getArea(circle))
	fmt.Printf("Area Dari Retangle : %f\n", getArea(retangle))
}

type Shape interface {
	Area() float64
}

type Circle struct {
	X      float64
	Y      float64
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Retangle struct {
	Width  float64
	Height float64
}

func (r Retangle) Area() float64 {
	return r.Height * r.Width
}

func getArea(s Shape) float64 {
	return s.Area()
}
