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
