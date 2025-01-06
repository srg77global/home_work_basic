package main

import (
	"flag"
	"hw15_go_sql/internal/handlers"
	"log"
	"net/http"
)

func main() {
	url := flag.String("url", "127.0.0.1", "URL")
	port := flag.String("port", "8080", "Port")
	flag.Parse()

	log.Println("server is running")

	//http.HandleFunc("/get/users/byUsername", v1.HandleGetUsersByUsernameV1) // function created without sqlc: ./v1/handlersV1.go
	http.HandleFunc("/get/users/byUsername", handlers.HandleGetUsersByUsername)
	http.HandleFunc("/get/users/byOrders", handlers.HandleGetUsersByOrders)
	http.HandleFunc("/get/products", handlers.HandleGetProducts)
	http.HandleFunc("/get/order", handlers.HandleGetOrder)

	http.HandleFunc("/post/user", handlers.HandlePostUser)
	http.HandleFunc("/post/product", handlers.HandlePostProduct)
	http.HandleFunc("/post/order", handlers.HandlePostOrder)

	http.HandleFunc("/put/username", handlers.HandlePutUsername)
	http.HandleFunc("/put/productprice", handlers.HandlePutProductprice)

	http.HandleFunc("/delete/user", handlers.HandleDeleteUser)
	http.HandleFunc("/delete/product", handlers.HandleDeleteProduct)
	http.HandleFunc("/delete/order", handlers.HandleDeleteOrder)

	err := http.ListenAndServe(*url+":"+*port, nil) //nolint

	if err != nil {
		log.Println("Error of starting Server: ", err)
		return
	}
}

// Server:
// 	- GET
// 		- users
// 			- byUsername
// 			- byOrders
// 		- products
// 		- order
// 	- POST
// 		- user
// 		- product
// 		- order
// 	- PUT
// 		- username
// 		- productprice
// 	- DELETE
// 		- user
// 		- product
// 		- order
