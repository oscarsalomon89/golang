package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const port = ":9000"

type Book struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Price  int    `json:"price"`
}

var db []Book

func main() {
	b1 := Book{
		ID:     1,
		Title:  "Dune",
		Price:  1965,
		Author: "Frank Herbert",
	}

	b2 := Book{
		ID:     2,
		Title:  "Cita con Rama",
		Price:  1974,
		Author: "Arthur C. Clarke",
	}

	b3 := Book{
		ID:     3,
		Title:  "Un guijarro en el cielo",
		Price:  500,
		Author: "Isaac Asimov",
	}

	db = append(db, b1, b2, b3)

	r := gin.Default()

	r.GET("/", index)
	r.GET("/books", getBooks)
	r.GET("/books/{id}", getBook)
	r.POST("/books", addBook)
	r.PUT("/books/{id}", updateBook)
	r.DELETE("/books/{id}", deleteBook)

	log.Println("Server listening on port", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalln(err)
	}
}

/*
w response: respuesta del servidor al cliente
r request: peticion del cliente al servidor
*/
func index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "Bienvenido a mi increible API!")
}

func getBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"error": false,
		"data":  db,
	})
}

type ResponseInfo struct {
	Error bool   `json:"error"`
	Data  string `json:"data"`
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	idRequest, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	for _, v := range db {
		if v.ID == idRequest {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(&v)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ResponseInfo{
		Error: true,
		Data:  "book not found",
	})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	db = append(db, book)

	json.NewEncoder(w).Encode(db)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseInfo{
			Error: true,
			Data:  err.Error(),
		})
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, v := range db {
		if v.ID == id {
			db[i] = book
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func updateBook2(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		encodeJSON(w, ResponseInfo{
			Error: true,
			Data:  err.Error(),
		}, http.StatusBadRequest)
		return
	}

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, v := range db {
		if v.ID == id {
			db[i] = book
		}
	}

	encodeJSON(w, db, http.StatusOK)
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ResponseInfo{
			Error: true,
			Data:  err.Error(),
		})
	}

	for i, v := range db {
		if v.ID == id {
			db = append(db[:i], db[i+1:]...)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(db)
}

func encodeJSON(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}
