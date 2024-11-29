package main

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
	Size   int
	Rate   float32
}

func (b *Book) MarshalJSON() ([]byte, error) {
	MByte := []byte("1")
	return MByte, nil
}

func (b *Book) UnmarshalJSON(data []byte) error {
	return nil
}

func main() {
	BookM := &SBook{
		ID:     1,
		Title:  "Title1",
		Author: "Author1",
		Year:   1999,
		Size:   123,
		Rate:   5.8,
	}

	j, err := proto.Marshal(BookM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", j)

	var BookU SBook

	err = proto.Unmarshal(j, &BookU)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", j)
}
