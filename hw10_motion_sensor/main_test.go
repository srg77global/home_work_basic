package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestInput(t *testing.T) {
	testCases := []struct {
		desc string
		data int
		res  int
	}{
		{
			desc: "Positive1",
			data: 2,
			res:  2,
		},
		{
			desc: "Positive2",
			data: 5235236,
			res:  5235236,
		},
		{
			desc: "BorderCase",
			data: 0,
			res:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			cInput := make(chan int, 100)
			Input := ReadSens(time.Second)
			Input(cInput, tC.data)
			require.Equal(t, tC.res, <-cInput)
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

			for _, nums := range tC.data {
				chInput <- nums
			}
			close(chInput)

			Output(chInput, chOutput)
			res := <-chOutput
			require.Equal(t, tC.exp, res)
		})
	}
}
