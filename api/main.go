package main

import (
	"encoding/json"
	"jump/jump_interview/internal/db"
	"jump/jump_interview/internal/types"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	// Retrieve users from database
	users, err := db.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Into json and send
	jsonResp, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func createInvoice(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var invoice types.Invoice
	err := json.NewDecoder(r.Body).Decode(&invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert invoice into database
	err = db.CreateInvoice(invoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	// Decode request body
	var transaction types.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert transaction into database
	err = db.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	var err error
	// Connect to DB
	db.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	// HTTP methods
	router := mux.NewRouter()
	router.Methods("GET").Path("/users").HandlerFunc(getUsers)
	router.Methods("POST").Path("/invoice").HandlerFunc(createInvoice)
	router.Methods("POST").Path("/transaction").HandlerFunc(createTransaction)

	// CORS
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "X-Requested-With"})
	methods := handlers.AllowedMethods([]string{"GET", "POST"})
	origins := handlers.AllowedOrigins([]string{"*"})
	exposedHeaders := handlers.ExposedHeaders([]string{"Access-Control-Allow-Origin"})

	// Start server
	log.Println("Server started on port :8080")
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins, exposedHeaders)(router)))
}
