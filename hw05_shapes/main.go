package main

import (
	"fmt"
	"math"
)

type Shape interface {
	CalcArea() float64
}

type Circle struct {
	R float64
}

func (c *Circle) CalcArea() float64 {
	return math.Pi * (c.R * c.R)
}

type Rectangle struct {
	A, B float64
}

func (r *Rectangle) CalcArea() float64 {
	return r.A * r.B
}

type Triangle struct {
	A, B float64
}

func (t *Triangle) CalcArea() float64 {
	return t.A * t.B / 2
}

func calculateArea(s any) (float64, error) {
	if ss, ok := s.(Shape); ok {
		return ss.CalcArea(), nil
	}
	return 0, fmt.Errorf("ошибка: переданный объект не является фигурой")
}

func main() {
	a := &Circle{R: 5}
	b := &Rectangle{A: 10, B: 5}
	c := &Triangle{A: 8, B: 6}
	d := 10

	if s1, err := calculateArea(a); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Круг: радиус %.0f\nПлощадь: %.14f\n", a.R, s1)
	}

	if s2, err := calculateArea(b); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Прямоугольник: ширина %.0f, высота %.0f\nПлощадь: %.0f\n", b.A, b.B, s2)
	}

	if s3, err := calculateArea(c); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Треугольник: основание %.0f, высота %.0f\nПлощадь: %.0f\n", c.A, c.B, s3)
	}

	if s4, err := calculateArea(d); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Println(s4)
	}
}
