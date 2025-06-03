package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "admin"
	password = "secret"
	dbname   = "meli"
)

// Estrutura do usuário
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var db *sql.DB
var rdb *redis.Client

func main() {
	// Configurar conexión a PostgreSQL
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := mux.NewRouter()
	router.HandleFunc("/users/{id}", getUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}", updateUserHandler).Methods("PUT")

	fmt.Println("Servidor corriendo en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Handler para obter usuário
func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cacheKey := fmt.Sprintf("user_%d", id)

	// Tentar obter do Redis
	cachedUser, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		log.Println("🔵 Cache hit (Redis)!")
		w.Write([]byte(cachedUser))
		return
	}

	// Buscar no "banco de dados"
	var user User
	err = db.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusNotFound)
		return
	}

	// Salvar no Redis com expiração de 5 minutos
	userJSON, _ := json.Marshal(user)
	rdb.Set(ctx, cacheKey, userJSON, 5*time.Minute)
	log.Println("🟠 Cache miss! Salvando usuário no Redis.")

	w.Write(userJSON)
}

// Handler para atualizar usuário
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	_, err = db.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", user.Name, user.Email, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = id

	// Invalida cache no Redis
	cacheKey := fmt.Sprintf("user_%d", id)
	rdb.Del(ctx, cacheKey)
	log.Println("🔴 Cache do Redis invalidado!")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
