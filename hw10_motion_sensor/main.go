package main

import (
	"fmt"
	"time"
)

func Input(chInput chan<- int, data []int, chStop chan<- struct{}) {
	ticker := time.NewTicker(10 * time.Second) // tested using 5 sec
	defer ticker.Stop()
	timeout := time.After(time.Minute) // tested using 6 sec

	for {
		select {
		case <-ticker.C:
			for _, dataInt := range data {
				chInput <- dataInt
			}
			fmt.Println("[goroutine 1] first case executed")
		case <-timeout:
			fmt.Println("[goroutine 1] second case executed")
			chStop <- struct{}{}
			return
		}
	}
}

func Output(chInput <-chan int, chOutput chan<- int, chStop1 <-chan struct{}, chStop2 chan<- struct{}) {
FOR:
	for {
		select {
		case data := <-chInput:
			for i := 0; i < 9; i++ {
				data += <-chInput
			}
			fmt.Println("[goroutine 2] first case executed")
			chOutput <- data / 10
		case <-chStop1:
			fmt.Println("[goroutine 2] second case executed")
			break FOR
		case <-time.After(time.Minute): // tested using 7 sec
			fmt.Println("[goroutine 2] third case executed")
			break FOR
		}
	}
	fmt.Println("[goroutine 2] sent struct{}")
	chStop2 <- struct{}{}
}

func main() {
	chMainInput := make(chan int, 100)
	chMainOutput := make(chan int, 100)
	chMainStop1 := make(chan struct{})
	chMainStop2 := make(chan struct{})
	dataMain := []int{}

	fmt.Println("[main] variables created")

	dataMain = append(dataMain, 11, 21, 31, 41, 51, 61, 71, 81, 91, 101)

	go Input(chMainInput, dataMain, chMainStop1)

	go Output(chMainInput, chMainOutput, chMainStop1, chMainStop2)

	for {
		select {
		case dataOutput := <-chMainOutput:
			println(dataOutput)
		case <-chMainStop2:
			fmt.Println("[main] return")
			return
		}
	}
}
