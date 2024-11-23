package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInput(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
	}{
		{
			desc: "Positive1",
			data: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
		},
		{
			desc: "Positive2",
			data: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
		},
		{
			desc: "BorderCase",
			data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chInput := make(chan int, 100)
			chStop := make(chan struct{}, 1)
			chInputSlice := make([]int, 0)

			Input(chInput, tC.data, chStop)

			for i := 0; i < len(tC.data); i++ {
				nums := <-chInput
				chInputSlice = append(chInputSlice, nums)
			}

			require.Equal(t, tC.data, chInputSlice)
		})
	}
}

func TestOutput(t *testing.T) {
	testCases := []struct {
		desc string
		data []int
		exp  int
	}{
		{
			desc: "Positive1",
			data: []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2},
			exp:  2,
		},
		{
			desc: "Positive2",
			data: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			exp:  55,
		},
		{
			desc: "BorderCase",
			data: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			exp:  0,
		},
		{
			desc: "NegativeNums",
			data: []int{-1, -2, -3, 4, -5, 6, -7, 8, -9, -10},
			exp:  -1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			chInput := make(chan int, 100)
			chOutput := make(chan int, 100)
			chStop1 := make(chan struct{}, 1)
			chStop2 := make(chan struct{}, 1)

			for _, nums := range tC.data {
				chInput <- nums
			}

			Output(chInput, chOutput, chStop1, chStop2)

			res := <-chOutput

			require.Equal(t, tC.exp, res)
		})
	}
}
