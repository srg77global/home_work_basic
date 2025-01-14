package main

import "fmt"

func createChessboard(size int) string {
	var str string
	for b := 1; b <= size; b++ {
		for a := 1; a <= size; a++ {
			switch {
			case a%2 == 0 && b%2 == 0:
				str += "#"
			case a%2 == 0 && b%2 != 0:
				str += " "
			case a%2 != 0 && b%2 != 0:
				str += "#"
			case a%2 != 0 && b%2 == 0:
				str += " "
			}
		}
		str += "\n"
	}
	return str
}

func main() {
	var size int
	fmt.Printf("Enter the size for the chessboard:")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(createChessboard(size))
}
