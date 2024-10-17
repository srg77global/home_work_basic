package hw05

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
		if ss.CalcArea() >= 0 {
			return ss.CalcArea(), nil
		}
		return 0, fmt.Errorf("ошибка: CalcArea() = %f/n", ss.CalcArea())
	}
	return 0, fmt.Errorf("ошибка: переданный объект не является фигурой")
}
