package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/srg77global/home_work_basic/hw09_serialize/proto"
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
	bJSON, err := json.Marshal(b)
	return bJSON, err
}

func (b *Book) UnmarshalJSON(data []byte) error {
	err := json.Unmarshal(data, b)
	return err
}

func main() {
	book := &proto.BookProto{
		ID:     1,
		Title:  "title",
		Author: "author",
		Year:   1990,
		Size:   456,
		Rate:   7.3,
	}

	bookJSON, err := json.Marshal(book)
	if err != nil {
		log.Println("Error of Marshaling: ", err)
		return
	}

	fmt.Println(string(bookJSON))

	newBook := &proto.BookProto{}

	err = json.Unmarshal(bookJSON, &newBook)
	if err != nil {
		log.Println("Error of Unmarshaling: ", err)
		return
	}

	fmt.Printf("%+v\n", newBook)
}
