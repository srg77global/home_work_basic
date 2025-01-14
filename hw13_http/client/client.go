package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func (u *User) String() string {
	uMarshalled, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
	}
	return string(uMarshalled)
}

func main() {
	url := flag.String("url", "127.0.0.1:9090", "URL for request.")
	path := flag.String("path", "get_user", `path of resource: "get_user" or "post_user"`)
	flag.Parse()

	urlStr := "http://" + *url + "/" + *path

	switch *path {
	case "get_user":

		conn, err := http.Get(urlStr) //nolint
		if err != nil {
			fmt.Println("Error connection with Method GET: ", err)
			return
		}
		defer conn.Body.Close()

		if conn.StatusCode != http.StatusOK {
			fmt.Println("Incorrect status: ", conn.StatusCode)
			return
		}

		var getUser User

		body, err := io.ReadAll(conn.Body)
		if err != nil {
			fmt.Println("Error of reading: ", err)
			return
		}

		err = json.Unmarshal(body, &getUser)
		if err != nil {
			fmt.Println("Error Unmarshalling: ", err)
			return
		}

		fmt.Printf("%+v\n", getUser)

	case "post_user":

		newUser := User{
			Name:    "Anat",
			Surname: "Volov",
			Age:     22,
		}

		conn, err := http.Post(urlStr, "encording/json", bytes.NewBufferString(newUser.String())) //nolint
		if err != nil {
			fmt.Println("Error connection with Method POST: ", err)
			return
		}
		defer conn.Body.Close()

		if conn.StatusCode != http.StatusCreated {
			fmt.Println("Error status: ", conn.StatusCode)
			return
		}

		body, err := io.ReadAll(conn.Body)
		if err != nil {
			fmt.Println("Error of reading: ", err)
			return
		}

		var postedUser User

		err = json.Unmarshal(body, &postedUser)
		if err != nil {
			fmt.Println("Error Unmarshalling: ", err)
			return
		}

		fmt.Printf("%+v\n", postedUser)

	default:
		fmt.Println("Error Method.")
	}
}
