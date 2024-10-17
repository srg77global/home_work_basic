package main

import "fmt"

type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (s *Book) WriteID(x int) {
	s.id = x
}

func (s *Book) WriteTitle(x string) {
	s.title = x
}

func (s *Book) WriteAuthor(x string) {
	s.author = x
}

func (s *Book) WriteYear(x int) {
	s.year = x
}

func (s *Book) WriteSize(x int) {
	s.size = x
}

func (s *Book) WriteRate(x float32) {
	s.rate = x
}

func (s *Book) GetID() int {
	return s.id
}

func (s *Book) GetTitle() string {
	return s.title
}

func (s *Book) GetAuthor() string {
	return s.author
}

func (s *Book) GetYear() int {
	return s.year
}

func (s *Book) GetSize() int {
	return s.size
}

func (s *Book) GetRate() float32 {
	return s.rate
}

type CBooks struct {
	enum int
}

func (s CBooks) Comparator(a, b Book) bool {
	switch s.enum {
	case 1:
		return a.year > b.year
	case 2:
		return a.size > b.size
	case 3:
		return a.rate > b.rate
	default:
		return false
	}
}

func main() {
	FirstBook := Book{
		id:     1,
		title:  "Title1",
		author: "Author1",
		year:   2024,
		size:   999,
		rate:   9.9,
	}
	SecondBook := Book{
		id:     2,
		title:  "Title2",
		author: "Author2",
		year:   2020,
		size:   111,
		rate:   1.1,
	}
	ComparableBook := CBooks{1}
	fmt.Println(
		ComparableBook.Comparator(FirstBook, SecondBook))
}

/* func (s Comp) Comparator(enum int) func(Comp, Comp) bool {
	switch enum {
	case 1:
		return func(a, b Comp) bool {
			return a.year > b.year
		}
	case 2:
		return func(a, b Comp) bool {
			return a.size > b.size
		}
	default:
		return func(a, b Comp) bool {
			return a.rate > b.rate
		}
	}
}
	var choose int
	var r float32
	var strings string
	books := Book{
		id = 1,
		title = "Stories of Columb",
		author = "John Block",
		year = 2020,
		size = 293,
		rate = 6.7,
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
*/
