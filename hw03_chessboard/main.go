package main

import (
	"fmt"

	"github.com/srg77global/home_work_basic/hw03_chessboard/pkg"
)

func main() {
	var size int
	fmt.Printf("Enter the size for the chessboard:")
	_, err := fmt.Scanf("%d", &size)
	if err != nil {
		fmt.Println(err)
		return
	}
	pkg.CreatingBoard(size)
}
