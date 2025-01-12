package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/jackc/pgx/stdlib" // driver
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
)

func NewConnDB(ctx context.Context) (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("no .env file found")
	}
	dbDSN, exists := os.LookupEnv("DB_DSN")
	if exists {
		log.Printf("db DSN: %v", dbDSN)
	}

	db, err := sql.Open("pgx", dbDSN)
	if err != nil {
		return nil, fmt.Errorf("error loading driver: %w", err)
	}

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(0)

	if err = db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	return db, nil
}

func HandleGetUsersByUsernameV1(w http.ResponseWriter, r *http.Request) {
	log.Printf("Method: %+v\n", r)

	ctx := context.Background()

	db, err := NewConnDB(ctx)
	if err != nil {
		log.Printf("error creating new connection DB: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer db.Close()

	names := struct {
		Name1 string `db:"name1" json:"name1"`
		Name2 string `db:"name2" json:"name2"`
	}{}

	err = json.NewDecoder(r.Body).Decode(&names)
	if err != nil {
		log.Printf("error decoding: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	query := `
	SELECT * FROM users
	WHERE name=$1 OR name=$2;
	`

	rows, err := db.QueryContext(ctx, query, names.Name1, names.Name2)
	if err != nil {
		log.Printf("error method QueryContext: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type data = struct {
		ID       pgtype.UUID `db:"id" json:"id"`
		Name     string      `db:"name" json:"name"`
		Email    string      `db:"email" json:"email"`
		Password string      `db:"password" json:"password"`
	}
	dataSlice := []data{}
	dataRows := data{}

	for rows.Next() {
		if err = rows.Scan(&dataRows.ID, &dataRows.Name, &dataRows.Email, &dataRows.Password); err != nil {
			log.Printf("error rows scanning: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		dataSlice = append(dataSlice, dataRows)
	}

	if err = rows.Err(); err != nil {
		log.Printf("error getting rows: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("%+v\n", dataSlice)

	body, err := json.Marshal(dataSlice)
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
