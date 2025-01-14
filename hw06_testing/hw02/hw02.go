package hw02

import (
	"fmt"

	"github.com/srg77global/home_work_basic/hw02_fix_app/printer"
	"github.com/srg77global/home_work_basic/hw02_fix_app/reader"
	"github.com/srg77global/home_work_basic/hw02_fix_app/types"
)

func main() {
	path := "data.json"

	fmt.Printf("Enter data file path: ")

	var err error

	ret, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret)

	var staff []types.Employee

	if len(path) == 0 {
		path = "data.json"
	}

	staff, err = reader.ReadJSON(path)

	fmt.Print(err)

	printer.PrintStaff(staff)
}
