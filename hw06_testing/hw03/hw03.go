package hw03

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
