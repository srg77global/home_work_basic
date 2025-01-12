package hfuncs

import (
	"encoding/json"
	"log"
	"net/http"
)

type NameStr struct {
	Name string `db:"name" json:"name"`
}

func Response(w http.ResponseWriter, data any) {
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
