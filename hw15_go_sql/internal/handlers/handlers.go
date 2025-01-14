package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/srg77global/home_work_basic/hw15_go_sql/internal/handlers/hfuncs"
	"github.com/srg77global/home_work_basic/hw15_go_sql/internal/repository/online_shop"
	"github.com/srg77global/home_work_basic/hw15_go_sql/internal/repository/transaction"
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

type APIHandler struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func NewAPIHandler(ctx context.Context, db *pgxpool.Pool) (*APIHandler, error) {
	h := APIHandler{db: db}
	if db != h.db {
		return nil, errors.New("error creating APIHandler struct")
	}
	h.ctx = ctx
	return &h, nil
}

func (h *APIHandler) HandleGetUsersByUsername(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	names := online_shop.SelectTwoUsersByUsernameParams{}
	err := json.NewDecoder(r.Body).Decode(&names)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.SelectTwoUsersByUsername(h.ctx, names)
	if err != nil {
		log.Println("error function SelectTwoUsersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleGetUsersByOrders(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	data, err := repoOnlineShop.SelectUsersByOrders(h.ctx)
	if err != nil {
		log.Println("error function SelectUsersByOrders: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleGetProducts(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	prices := online_shop.SelectProductsByPricesParams{}
	err := json.NewDecoder(r.Body).Decode(&prices)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.SelectProductsByPrices(h.ctx, prices)
	if err != nil {
		log.Println("error function SelectProductsByPrices: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleGetOrder(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.SelectOrdersByUsername(h.ctx, nameD.Name)
	if err != nil {
		log.Println("error function SelectOrdersByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	dataS := online_shop.InsertUserParams{}
	err := json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.InsertUser(h.ctx, dataS)
	if err != nil {
		log.Println("error function InsertUser: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandlePostProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	dataS := online_shop.InsertProductParams{}
	err := json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.InsertProduct(h.ctx, dataS)
	if err != nil {
		log.Println("error function InsertProduct: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandlePostOrder(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

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
	err := json.NewDecoder(r.Body).Decode(&createOrderJSON)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = transaction.CreateOrderByUsernameProductnameAndQuantity(
		h.ctx, h.db, repoOnlineShop, *createOrder.Username, *createOrder.Productname, createOrderJSON.Quantity,
	)
	if err != nil {
		log.Println("error function CreateOrderByUsernameProductnameAndQuantity: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", txSucceed.String())
	hfuncs.Response(w, txSucceed.String())
}

func (h *APIHandler) HandlePutUsername(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	dataS := online_shop.UpdateUsernameByNameParams{}
	err := json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.UpdateUsernameByName(h.ctx, dataS)
	if err != nil {
		log.Println("error function UpdateUsernameByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandlePutProductprice(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	dataS := online_shop.UpdateProductpriceByNameParams{}
	err := json.NewDecoder(r.Body).Decode(&dataS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.UpdateProductpriceByName(h.ctx, dataS)
	if err != nil {
		log.Println("error function UpdateProductpriceByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteUserByName(h.ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteUserByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteProductByName(h.ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteProductByName: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}

func (h *APIHandler) HandleDeleteOrder(w http.ResponseWriter, r *http.Request) {
	log.Printf("new request: %+v\n", r)
	repoOnlineShop := online_shop.New(h.db)

	nameD, err := hfuncs.DecodeStr(w, r)
	if err != nil {
		log.Printf("error function DecodeStr: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := repoOnlineShop.DeleteOrderByUsername(h.ctx, nameD.Name)
	if err != nil {
		log.Println("error function DeleteOrderByUsername: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", data)
	hfuncs.Response(w, data)
}
