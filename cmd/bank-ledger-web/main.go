package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// api routes
	r.HandleFunc("/api/account", placeholder).Methods("PUT")             // login
	r.HandleFunc("/api/account", placeholder).Methods("POST")            // create account
	r.HandleFunc("/api/account/{accountID}", placeholder).Methods("GET") // get account details
	r.HandleFunc("/api/account/{accountID}", placeholder).Methods("PUT") // logout
	r.HandleFunc("/api/transactions", placeholder).Methods("POST")       // record new transaction
	r.HandleFunc("/api/transactions", placeholder).Methods("GET")        // get transaction history
	r.HandleFunc("/api/balance", placeholder).Methods("GET")             // get balance

	http.ListenAndServe(":8080", r)
}

func placeholder(w http.ResponseWriter, r *http.Request) {
	log.Println("placeholder route")
	_, err := w.Write([]byte("Hi, testing route"))
	log.Println(err.Error())
}
