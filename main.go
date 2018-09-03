package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

var man *BinManager= NewBinManager()

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(200)
	v, ok := vars["id"]
	if !ok {
		log.Println("create error")
	}

	if bin := man.Create(v); bin == nil {
		fmt.Fprint(w, "error")
	} else {
		fmt.Fprint(w, "success "+v)
	}
}
func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Items!"))
	log.Println(r.Method)

}
func InHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!"))
	log.Println(r.Method)
}

func main() {
	mux := mux.NewRouter().StrictSlash(false)

	mux.HandleFunc("/v1/{id}/{_:create/?}", CreateHandler).Methods("POST")
	mux.HandleFunc("/v1/{id}/items", ItemsHandler).Methods("GET")
	mux.HandleFunc("/v1/{id}/items/{num:[0-9]+(?:\\/)?}", ItemsHandler).Methods("GET")
	mux.HandleFunc("/v1/{id}/{_:in(?:/.*|$)}", InHandler)
	srv := &http.Server{
		Addr:    "localhost:8081",
		Handler: mux, // Pass our instance of gorilla/mux in.
	}

	srv.ListenAndServe()
}

func trimSlashMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
