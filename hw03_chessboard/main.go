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

	for row := 0; row < size; row++ {
		if row%2 == 0 {
			fmt.Print("#")
		}
		for cell := 1; cell <= size; cell++ {
			switch {
			case cell%2 == 0 && cell != size:
				fmt.Print("#")
			case cell%2 != 0 && cell != size:
				fmt.Print(" ")
			case cell%2 == 0 && cell == size && row%2 != 0:
				fmt.Printf("#\n")
			case cell%2 == 0 && cell == size && row%2 == 0:
				fmt.Printf("\n")
			case cell%2 != 0 && cell == size:
				fmt.Printf("\n")
			}
		}
	}
}
