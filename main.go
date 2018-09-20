package main

import (
	"./db"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if json.NewEncoder(w).Encode(res); err != nil {
		fmt.Fprint(w, "Error")
	}
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
	w.WriteHeader(http.StatusOK)
	req, err := db.NewRequest(time.Now(), r)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}
	bin.WriteLog(req)

	w.Write([]byte(r.RemoteAddr))
	println(vars["id"])
}

func main() {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/v1/{id}/{_:create/?}", CreateHandler).Methods("POST")
	router.HandleFunc("/v1/{id}/items", ItemsHandler).Methods("GET")
	router.HandleFunc("/v1/{id}/items/{num:[0-9]+(?:\\/)?}", ItemsHandler).Methods("GET")
	router.HandleFunc("/v1/{id}/{_:in(?:/.*|$)}", InHandler)
	srv := &http.Server{
		Addr:    "localhost:8081",
		Handler: router, // Pass our instance of gorilla/router in.
	}

	srv.ListenAndServe()
}
