package hw05

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcAreaCircle(t *testing.T) {
	testCases := []struct {
		desc string
		r    float64
		res  float64
	}{
		{
			desc: "Positive",
			r:    10,
			res:  314.1592653589793,
		},
		{
			desc: "NegativeValue",
			r:    -10,
			res:  314.1592653589793,
		},
		{
			desc: "ZeroValue",
			r:    0,
			res:  0,
		},
		{
			desc: "BorderCase",
			r:    1.7976931348623157e+308,
			res:  math.Inf(0),
		},
		{
			desc: "BigNumbersAComma",
			r:    1.15926535897932111111111111111,
			res:  4.221974342805975,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str := &Circle{tC.r}
			got := str.CalcArea()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestCalcAreaRectangle(t *testing.T) {
	testCases := []struct {
		desc string
		A    float64
		B    float64
		res  float64
	}{
		{
			desc: "Positive",
			A:    5,
			B:    10,
			res:  50,
		},
		{
			desc: "NegativeValue",
			A:    -10,
			B:    2,
			res:  -20,
		},
		{
			desc: "ZeroValue",
			A:    0,
			B:    0,
			res:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str := &Rectangle{tC.A, tC.B}
			got := str.CalcArea()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestCalcAreaTriangle(t *testing.T) {
	testCases := []struct {
		desc string
		A    float64
		B    float64
		res  float64
	}{
		{
			desc: "Positive",
			A:    5,
			B:    10,
			res:  25,
		},
		{
			desc: "NegativeValue",
			A:    -10,
			B:    2,
			res:  -10,
		},
		{
			desc: "ZeroValue",
			A:    0,
			B:    0,
			res:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			str := &Triangle{tC.A, tC.B}
			got := str.CalcArea()
			require.Equal(t, tC.res, got)
		})
	}
}

func TestCalculateArea(t *testing.T) {
	testCases := []struct {
		desc  string
		inter interface{}
		res   float64
	}{
		{
			desc:  "PositiveCircle",
			inter: &Circle{R: 10},
			res:   314.1592653589793,
		},
		{
			desc:  "PositiveRectangle",
			inter: &Rectangle{A: 10, B: 20},
			res:   200,
		},
		{
			desc:  "PositiveTriangle",
			inter: &Triangle{A: 10, B: 20},
			res:   100,
		},
		{
			desc:  "NegativeValue",
			inter: &Rectangle{A: -100, B: 10},
			res:   0,
		},
		{
			desc:  "ZeroValue",
			inter: &Rectangle{A: -10, B: 0},
			res:   0,
		},
		{
			desc:  "NotCorrectStruct",
			inter: 10,
			res:   0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, _ := calculateArea(tC.inter)
			require.Equal(t, tC.res, got)
		})
	}
}
