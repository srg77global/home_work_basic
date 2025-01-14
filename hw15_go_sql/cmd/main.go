package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/srg77global/home_work_basic/hw15_go_sql/internal/handlers"
	"github.com/srg77global/home_work_basic/hw15_go_sql/pkg/pgdb"
	v1 "github.com/srg77global/home_work_basic/hw15_go_sql/v1"
)

func main() {
	dbDSN, err := pgdb.LoadENV()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, db, err := pgdb.New(dbDSN, 5)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	h, err := handlers.NewAPIHandler(ctx, db)
	if err != nil {
		log.Println(err)
	}

	http.HandleFunc("GET /get/users/byUsername", v1.HandleGetUsersByUsernameV1) // function created without sqlc
	// http.HandleFunc("GET /get/users/byUsername", h.HandleGetUsersByUsername)
	http.HandleFunc("GET /get/users/byOrders", h.HandleGetUsersByOrders)
	http.HandleFunc("GET /get/products", h.HandleGetProducts)
	http.HandleFunc("GET /get/order", h.HandleGetOrder)

	http.HandleFunc("POST /post/user", h.HandlePostUser)
	http.HandleFunc("POST /post/product", h.HandlePostProduct)
	http.HandleFunc("POST /post/order", h.HandlePostOrder)

	http.HandleFunc("PUT /put/username", h.HandlePutUsername)
	http.HandleFunc("PUT /put/productprice", h.HandlePutProductprice)

	http.HandleFunc("DELETE /delete/user", h.HandleDeleteUser)
	http.HandleFunc("DELETE /delete/product", h.HandleDeleteProduct)
	http.HandleFunc("DELETE /delete/order", h.HandleDeleteOrder)

	url := flag.String("url", "127.0.0.1", "URL")
	port := flag.String("port", "8080", "Port")
	flag.Parse()

	log.Println("server is running")

	err = http.ListenAndServe(*url+":"+*port, nil) //nolint
	if err != nil {
		log.Println("Error of starting Server: ", err)
	}
}
