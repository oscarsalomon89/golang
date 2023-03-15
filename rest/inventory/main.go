package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8888"
	router := mux.NewRouter()

	router.HandleFunc("/ping", ping).Methods("GET")

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalln(err)
	}
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(ResponseInfo{
		Error: false,
		Data:  "pong",
	})
}
