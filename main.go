package main

import (
	"./db"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var man *db.BinManager = db.NewBinManager()

// CreateHandler is a handler of (POST) /v1/{id}/create
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok {
		log.Println("create error")
		w.WriteHeader(400)
		fmt.Fprint(w, "error")
	}

	if bin := man.Create(v); bin == nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "error")

	} else {
		w.WriteHeader(200)
		fmt.Fprint(w, "success create "+v)
	}
}

// ItemsHandler is a handler of (GET) /v1/{id}/items/{num}
func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, "bad id")
	}

	num, ok := vars["num"]
	if !ok {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad index")
	}

	index, err := strconv.Atoi(num)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, "Bad index")
		return
	}

	bin := man.Bin(v)
	if bin == nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "No bin")
		return
	}

	res, err := bin.ReadLog(index)
	if err != nil {
		fmt.Fprint(w, "No log")
		return
	}

	fmt.Fprint(w, res)
}

// InHandler is a handler of (ANY) /v1/id/in/*
func InHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bin := man.Bin(vars["id"])
	if bin == nil {
		w.WriteHeader(404)

		w.Write([]byte("BadRequest"))
		println(vars["id"])
		return
	}
	w.WriteHeader(200)
	bin.WriteLog(db.Request{Time: time.Now(), Request: r})

	w.Write([]byte(r.RemoteAddr))
	println(vars["id"])
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
