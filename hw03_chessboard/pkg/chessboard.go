package pkg

import "fmt"

var Row = 0

func CreatingBoard(size int) {
	for Row < size {
		if Row%2 == 0 {
			fmt.Print("#")
		}
		Cells(size)
		Row++
	}
}

func Cells(size int) {
	for cell := 1; cell <= size; cell++ {
		switch {
		case cell%2 == 0 && cell != size:
			fmt.Print("#")
		case cell%2 != 0 && cell != size:
			fmt.Print(" ")
		case cell%2 == 0 && cell == size && Row%2 != 0:
			fmt.Printf("#\n")
		case cell%2 == 0 && cell == size && Row%2 == 0:
			fmt.Printf("\n")
		case cell%2 != 0 && cell == size:
			fmt.Printf(" \n")
		}
	}
}
