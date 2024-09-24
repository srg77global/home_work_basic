package main

import (
	"fmt"
)

type Book struct {
	id     []int
	title  []string
	author []string
	year   []int
	size   []int
	rate   []float32
}

type Sets interface {
	writeID(x int)
	writeTitle(x string)
	writeAuthor(x string)
	writeYear(x int)
	writeSize(x int)
	writeRate(x float32)
}

type Gets interface {
	getID()
	getTitle()
	getAuthor()
	getYear()
	getSize()
	getRate()
}

func (s *Book) writeID(x int) {
	s.id = append(s.id, x)
}

func (s *Book) writeTitle(x string) {
	s.title = append(s.title, x)
}

func (s *Book) writeAuthor(x string) {
	s.author = append(s.author, x)
}

func (s *Book) writeYear(x int) {
	s.year = append(s.year, x)
}

func (s *Book) writeSize(x int) {
	s.size = append(s.size, x)
}

func (s *Book) writeRate(x float32) {
	s.rate = append(s.rate, x)
}

func (s *Book) getID() []int {
	return s.id
}

func (s *Book) getTitle() []string {
	return s.title
}

func (s *Book) getAuthor() []string {
	return s.author
}

func (s *Book) getYear() []int {
	return s.year
}

func (s *Book) getSize() []int {
	return s.size
}

func (s *Book) getRate() []float32 {
	return s.rate
}

func (s Book) Comparator(enum int) func(Book) bool {
	switch enum {
	case 1:
		return func(a Book) bool {
			return a.year[0] > a.year[1]
		}
	case 2:
		return func(a Book) bool {
			return a.size[0] > a.size[1]
		}
	default:
		return func(a Book) bool {
			return a.rate[0] > a.rate[1]
		}
	}
}

func main() {
	var choose int
	var r float32
	var strings string
	books := Book{
		id:     []int{1},
		title:  []string{"Stories of Columb"},
		author: []string{"John Block"},
		year:   []int{2020},
		size:   []int{293},
		rate:   []float32{6.7},
	}

START:
	fmt.Print(`
		1 - Store the book to compare
		2 - Show the field
		3 - Compare books
		4 - Exit
		5 - Show the library
		`)
	valc, errc := fmt.Scan(&choose)
	if errc != nil {
		fmt.Println(valc, "Error: ", errc)
	}

	switch choose {
	case 1:
		fmt.Println("Write ID:")
		val, err := fmt.Scan(&choose)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeID(choose)

		fmt.Println("Write Title:")
		val, err = fmt.Scan(&strings)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeTitle(strings)

		fmt.Println("Write Author:")
		val, err = fmt.Scan(&strings)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeAuthor(strings)

		fmt.Println("Write Year:")
		val, err = fmt.Scan(&choose)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeYear(choose)

		fmt.Println("Write Size:")
		val, err = fmt.Scan(&choose)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeSize(choose)

		fmt.Println("Write Rate:")
		val, err = fmt.Scan(&r)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		books.writeRate(r)

		goto START

	case 2:

		fmt.Print(`What do you want to get?
			1. ID
			2. Title
			3. Author
			4. Year
			5. Size
			6. Rate
			`)
		val, err := fmt.Scan(&choose)
		if err != nil {
			fmt.Println(val, "Error: ", err)
		}
		switch choose {
		case 1:
			fmt.Println(books.getID())
		case 2:
			fmt.Println(books.getTitle())
		case 3:
			fmt.Println(books.getAuthor())
		case 4:
			fmt.Println(books.getYear())
		case 5:
			fmt.Println(books.getSize())
		case 6:
			fmt.Println(books.getRate())
		}
		goto START

	case 3:
		fmt.Println(`Choose the comparable field:
		1 - Year
		2 - Size
		3 - Rate
		`)
		val5, err := fmt.Scanln(&choose)
		if err != nil {
			fmt.Println(val5, "Error: ", err)
		}
		compare := books.Comparator(choose)
		fmt.Println(compare(books))

		goto START
	case 4:
		break
	case 5:
		fmt.Println(books)

		goto START
	}
}
