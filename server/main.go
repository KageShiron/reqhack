package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/v1/{id}/{_:create/?}", CreateHandler).Methods("POST")
	router.HandleFunc("/v1/{id}/items", ItemsHandler).Methods("GET")
	router.HandleFunc("/v1/{id}/items/{num:[0-9]*(?:\\/)?}", ItemsHandler).Methods("GET")
	router.HandleFunc("/v1/{id}/{_:in(?:/.*|$)}", InHandler)
	srv := &http.Server{
		Addr:    "localhost:8081",
		Handler: router, // Pass our instance of gorilla/router in.
	}

	srv.ListenAndServe()
}
