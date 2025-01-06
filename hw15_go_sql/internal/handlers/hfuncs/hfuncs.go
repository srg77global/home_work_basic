package hfuncs

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"hw15_go_sql/pkg/pgdb"
)

type NameStr struct {
	Name string `db:"name" json:"name"`
}

func HeadHandler(w http.ResponseWriter, r *http.Request, method string) (context.Context, *pgxpool.Pool, error) {
	if r.Method != method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return nil, nil, errors.New("incorrect method")
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
		w.WriteHeader(http.StatusInternalServerError)
		return nil, nil, err
	}
	return ctx, db, nil
}

func EndHandler(w http.ResponseWriter, data any) {
	body, err := json.Marshal(data)
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

func DecodeStr(w http.ResponseWriter, r *http.Request) (NameStr, error) {
	NameS := NameStr{}
	err := json.NewDecoder(r.Body).Decode(&NameS)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return NameS, err
	}
	return NameS, nil
}
