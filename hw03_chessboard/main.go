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
	for b := 1; b <= size; b++ {
		for a := 1; a <= size; a++ {
			switch {
			case a%2 == 0 && b%2 == 0:
				fmt.Print("#")
			case a%2 == 0 && b%2 != 0:
				fmt.Print(" ")
			case a%2 != 0 && b%2 != 0:
				fmt.Print("#")
			case a%2 != 0 && b%2 == 0:
				fmt.Print(" ")
			}
		}
		fmt.Printf("\n")
	}
}
