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

const (
	year = iota
	size
	rate
)

type CBooks struct {
	enum int
}

func (s CBooks) Comparator(a, b Book) bool {
	switch s.enum {
	case 0:
		return a.year > b.year
	case 1:
		return a.size > b.size
	case 2:
		return a.rate > b.rate
	default:
		return false
	}
}

func createStruct(x int) *CBooks {
	return &CBooks{enum: x}
}

func main() {
	Book1 := Book{
		id:     1,
		title:  "Title1",
		author: "Author1",
		year:   1990,
		size:   777,
		rate:   4.5,
	}
	Book2 := Book{
		id:     2,
		title:  "Title1",
		author: "Author2",
		year:   2014,
		size:   555,
		rate:   5.3,
	}
	BookS := createStruct(rate)
	fmt.Println(BookS.Comparator(Book1, Book2))
}
