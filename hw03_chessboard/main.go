package main

import "fmt"

func main() {
	var size int

	fmt.Printf("Enter the size for the chessboard:")

	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println(err)
		return
	}

	for row := 1; row <= size; row++ {
		for cell := 1; cell <= size; cell++ {
			switch {
			case cell%2 == 0 && cell != size:
				fmt.Print("#")
			case cell%2 != 0 && cell != size:
				fmt.Print(" ")
			case cell%2 == 0 && cell == size:
				fmt.Printf("#\n")
			case cell%2 != 0 && cell == size:
				fmt.Printf(" \n")
			}
		}
	}
}
