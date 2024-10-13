package main

import (
	"fmt"
	"math"
)

type Shape interface {
	CalcArea() any
}

type Circle struct {
	R int
}

func (c *Circle) CalcArea() any {
	return math.Pi * (float64(c.R) * float64(c.R))
}

type Rectangle struct {
	A, B int
}

func (r *Rectangle) CalcArea() any {
	return r.A * r.B
}

type Triangle struct {
	A, B int
}

func (t *Triangle) CalcArea() any {
	return t.A * t.B / 2
}

func calculateArea(s any) (any, error) {
	if ss, ok := s.(Shape); ok {
		return ss.CalcArea(), nil
	}
	return "", fmt.Errorf("ошибка: переданный объект не является фигурой")
}

func main() {
	a := &Circle{R: 5}
	b := &Rectangle{A: 10, B: 5}
	c := &Triangle{A: 8, B: 6}
	d := 10

	if s1, err := calculateArea(a); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Круг: радиус %d\nПлощадь: %.14f\n", a.R, s1)
	}

	if s2, err := calculateArea(b); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Прямоугольник: ширина %d, высота %d\nПлощадь: %d\n", b.A, b.B, s2)
	}

	if s3, err := calculateArea(c); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Printf("Треугольник: основание %d, высота %d\nПлощадь: %d\n", c.A, c.B, s3)
	}

	if s4, err := calculateArea(d); err != nil {
		fmt.Println("Ошибка: переданный объект не является фигурой.")
	} else {
		fmt.Println(s4)
	}
}
