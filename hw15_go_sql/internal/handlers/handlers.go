package handlers

import (
	"context"
	"encoding/json"
	"hw15_go_sql/internal/repository/online_shop"
	"hw15_go_sql/internal/repository/transaction"
	"hw15_go_sql/pkg/pgdb"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/joho/godotenv"
)

type txSucceedT struct{ Tx string }

var txSucceed = &txSucceedT{Tx: "transaction succeed"}

func (u *txSucceedT) String() string {
	uMarshalled, err := json.Marshal(u)
	if err != nil {
		log.Println("Error Marshalling: ", err)
	}
	return string(uMarshalled)
}

func HandleGetUsersByUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	names := online_shop.SelectTwoUsersByUsernameParams{}
	err = json.NewDecoder(r.Body).Decode(&names)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	TwoUsersByUsername, err := repoOnlineShop.SelectTwoUsersByUsername(ctx, names)
	if err != nil {
		log.Println("error function SelectTwoUsersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", TwoUsersByUsername)

	body, err := json.Marshal(TwoUsersByUsername)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleGetUsersByOrders(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	UsersByOrders, err := repoOnlineShop.SelectUsersByOrders(ctx)
	if err != nil {
		log.Println("error function SelectUsersByOrders: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", UsersByOrders)

	body, err := json.Marshal(UsersByOrders)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	prices := online_shop.SelectProductsByPricesParams{}
	err = json.NewDecoder(r.Body).Decode(&prices)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ProductsByPrices, err := repoOnlineShop.SelectProductsByPrices(ctx, prices)
	if err != nil {
		log.Println("error function SelectProductsByPrices: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", ProductsByPrices)

	body, err := json.Marshal(ProductsByPrices)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameS := struct {
		Name string `db:"name" json:"name"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&nameS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	OrdersByUsername, err := repoOnlineShop.SelectOrdersByUsername(ctx, nameS.Name)
	if err != nil {
		log.Println("error function SelectOrdersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", OrdersByUsername)

	body, err := json.Marshal(OrdersByUsername)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	data := online_shop.InsertUserParams{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	UserID, err := repoOnlineShop.InsertUser(ctx, data)
	if err != nil {
		log.Println("error function InsertUser: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", UserID)

	body, err := json.Marshal(UserID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	data := online_shop.InsertProductParams{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ProductID, err := repoOnlineShop.InsertProduct(ctx, data)
	if err != nil {
		log.Println("error function InsertProduct: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", ProductID)

	body, err := json.Marshal(ProductID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	createOrderJSON := struct {
		Username    string         `db:"username" json:"username"`
		Productname string         `db:"productname" json:"productname"`
		Quantity    pgtype.Numeric `db:"quantity" json:"quantity"`
	}{}

	createOrder := struct {
		Username    *string
		Productname *string
	}{
		Username:    &createOrderJSON.Username,
		Productname: &createOrderJSON.Productname,
	}
	err = json.NewDecoder(r.Body).Decode(&createOrderJSON)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = transaction.CreateOrderByUsernameProductnameAndQuantity(
		ctx, db, repoOnlineShop, *createOrder.Username, *createOrder.Productname, createOrderJSON.Quantity,
	)
	if err != nil {
		log.Println("error function CreateOrderByUsernameProductnameAndQuantity: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Println(txSucceed)

	body, err := json.Marshal(txSucceed.String())
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandlePutUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	data := online_shop.UpdateUsernameByNameParams{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	UserID, err := repoOnlineShop.UpdateUsernameByName(ctx, data)
	if err != nil {
		log.Println("error function UpdateUsernameByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", UserID)

	body, err := json.Marshal(UserID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandlePutProductprice(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	data := online_shop.UpdateProductpriceByNameParams{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ProductID, err := repoOnlineShop.UpdateProductpriceByName(ctx, data)
	if err != nil {
		log.Println("error function UpdateProductpriceByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", ProductID)

	body, err := json.Marshal(ProductID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameS := struct {
		Name string `db:"name" json:"name"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&nameS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	UserID, err := repoOnlineShop.DeleteUserByName(ctx, nameS.Name)
	if err != nil {
		log.Println("error function DeleteUserByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", UserID)

	body, err := json.Marshal(UserID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameS := struct {
		Name string `db:"name" json:"name"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&nameS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ProductID, err := repoOnlineShop.DeleteProductByName(ctx, nameS.Name)
	if err != nil {
		log.Println("error function DeleteProductByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", ProductID)

	body, err := json.Marshal(ProductID)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}

func HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		log.Println("incorrect method")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Method: %v\n", r.Method)

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	ctx := context.Background()
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Println("DSN: ", dbDSN)
	}

	db, err := pgdb.New(ctx, dbDSN, 1)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameS := struct {
		Name string `db:"name" json:"name"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&nameS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	name, err := repoOnlineShop.DeleteOrderByUsername(ctx, nameS.Name)
	if err != nil {
		log.Println("error function DeleteOrderByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", name)

	body, err := json.Marshal(name)
	if err != nil {
		log.Println("error marshalling: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		log.Println("error write to header: ", err)
	}
}
