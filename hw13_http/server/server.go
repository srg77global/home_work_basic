package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("Incorrect Method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	fmt.Printf("%+v", &r)
	fmt.Printf("Method: %v\n", r.Method)

	oldUser := User{
		Name:    "Iva",
		Surname: "Inov",
		Age:     30,
	}

	body, err := json.Marshal(oldUser)
	if err != nil {
		fmt.Println("Error Marshalling: ", err)
		return
	}

	fmt.Printf("%+v\n", oldUser)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Writing error: ", err)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Println("Incorrect Method")
		return
	}

	fmt.Printf("Method: %v\n", r.Method)

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println("Decoding error: ", err)
		return
	}

	fmt.Printf("Created new user:\n%+v\n", newUser)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&newUser)
	if err != nil {
		fmt.Println("Encoding error: ", err)
		return
	}
}

func main() {
	url := flag.String("url", "127.0.0.1", "URL.")
	port := flag.String("port", "9090", "Port.")
	flag.Parse()

	http.HandleFunc("/get_user", handleGet)
	http.HandleFunc("/post_user", handlePost)
	err := http.ListenAndServe(*url+":"+*port, nil) //nolint
	if err != nil {
		fmt.Println("Error of starting Server: ", err)
		return
	}
}
