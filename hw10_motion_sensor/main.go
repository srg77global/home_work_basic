package main

import (
	"fmt"
	"time"
)

func ConstrInput(clock time.Duration) func(cInput chan<- int, data int) {
	return func(cInput chan<- int, data int) {
		timeout := time.After(clock)
		for {
			select {
			case cInput <- data:
				fmt.Println("[goroutine 1] first case executed")
			case <-timeout:
				fmt.Println("[goroutine 1] second case executed")
				close(cInput)
				return
			}
		}
	}
}

func Output(cInput <-chan int, cOutput chan<- int, cStop chan<- struct{}) {
	var (
		val       int
		dataSlice []int
	)

	for {
		if data, ok := <-cInput; ok {
			fmt.Println("[goroutine 2] first case executed")
			dataSlice = append(dataSlice, data)
			if len(dataSlice) == 10 {
				for _, valS := range dataSlice {
					val += valS
				}
				cOutput <- val / 10
				val = 0
				dataSlice = []int{}
			}
		} else {
			fmt.Println("[goroutine 2] second case executed")
			cStop <- struct{}{}
			return
		}
	}
}

func main() {
	cInput := make(chan int, 100)
	cOutput := make(chan int, 100)
	cStop := make(chan struct{})
	dataV := 523

	fmt.Println("[main] variables created")

	Input := ConstrInput(time.Minute)

	go Input(cInput, dataV)
	go Output(cInput, cOutput, cStop)

	for {
		select {
		case dataOutput := <-cOutput:
			fmt.Println(dataOutput)
		case <-cStop:
			fmt.Println("[main] return")
			return
		}
	}
}


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
			Input := ConstrInput(time.Second)
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
			chStop := make(chan struct{}, 1)

			for _, nums := range tC.data {
				chInput <- nums
			}
			close(chInput)

			Output(chInput, chOutput, chStop)
			res := <-chOutput
			require.Equal(t, tC.exp, res)
		})
	}
}
