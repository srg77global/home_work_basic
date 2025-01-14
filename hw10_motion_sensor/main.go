package main

import (
	"fmt"
	"time"
)

func ReadSens(clock time.Duration) func(cInput chan<- int, data int) {
	return func(cInput chan<- int, data int) {
		timeout := time.After(clock)

		for {
			select {
			case cInput <- data:
				fmt.Println("[goroutine 1]: data sent")
			case <-timeout:
				fmt.Println("[goroutine 1]: timeout")
				close(cInput)
				return
			}
		}
	}
}

func Output(cInput <-chan int, cOutput chan<- int) {
	val := 0
	dataSlice := []int{}

	for valS1 := range cInput {
		dataSlice = append(dataSlice, valS1)
		fmt.Println("[goroutine 2]: for range executed")

		if len(dataSlice) == 10 {
			for _, valS2 := range dataSlice {
				val += valS2
			}

			cOutput <- (val / 10)
			val = 0
			dataSlice = []int{}
		}
	}

	fmt.Println("[goroutine 2]: finished")
	close(cOutput)
}

func main() {
	cInput := make(chan int, 100)
	cOutput := make(chan int, 100)
	dataV := 523
	fmt.Println("[main]: variables created")

	Input := ReadSens(130 * time.Microsecond)

	go Input(cInput, dataV)
	go Output(cInput, cOutput)

	for dataOutput := range cOutput {
		fmt.Println(dataOutput)
	}

	fmt.Println("[main]: finished")
}
