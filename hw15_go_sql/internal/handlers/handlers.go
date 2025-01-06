package handlers

import (
	"encoding/json" //nolint
	"hw15_go_sql/internal/handlers/hfuncs"
	"hw15_go_sql/internal/repository/online_shop"
	"hw15_go_sql/internal/repository/transaction"
	"log" //nolint
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
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
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodGet)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
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

	data, err := repoOnlineShop.SelectTwoUsersByUsername(ctx, names)
	if err != nil {
		log.Println("error function SelectTwoUsersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleGetUsersByOrders(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodGet)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	data, err := repoOnlineShop.SelectUsersByOrders(ctx)
	if err != nil {
		log.Println("error function SelectUsersByOrders: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodGet)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
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

	data, err := repoOnlineShop.SelectProductsByPrices(ctx, prices)
	if err != nil {
		log.Println("error function SelectProductsByPrices: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodGet)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.SelectOrdersByUsername(ctx, nameD.Name)
	if err != nil {
		log.Println("error function SelectOrdersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandlePostUser(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodPost)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	dataS := online_shop.InsertUserParams{}
	err = json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.InsertUser(ctx, dataS)
	if err != nil {
		log.Println("error function InsertUser: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodPost)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	dataS := online_shop.InsertProductParams{}
	err = json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.InsertProduct(ctx, dataS)
	if err != nil {
		log.Println("error function InsertProduct: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodPost)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
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
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodPut)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	dataS := online_shop.UpdateUsernameByNameParams{}
	err = json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.UpdateUsernameByName(ctx, dataS)
	if err != nil {
		log.Println("error function UpdateUsernameByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandlePutProductprice(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodPut)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	dataS := online_shop.UpdateProductpriceByNameParams{}
	err = json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.UpdateProductpriceByName(ctx, dataS)
	if err != nil {
		log.Println("error function UpdateProductpriceByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodDelete)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteUserByName(ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteUserByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodDelete)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteProductByName(ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteProductByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}

func HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	ctx, db, err := hfuncs.HeadHandler(w, r, http.MethodDelete)
	if err != nil {
		log.Printf("error function HeadHandler: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()
	repoOnlineShop := online_shop.New(db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteOrderByUsername(ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteOrderByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.EndHandler(w, data)
}
