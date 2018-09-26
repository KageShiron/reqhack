package main

import (
	"encoding/json"
	"fmt"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"

	"github.com/KageShiron/reqhack/server/db"
	"github.com/gorilla/mux"
)

type simpleResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func restError(w http.ResponseWriter, code int, message string) {
	e := simpleResponse{Message: message, Code: code}
	ren.JSON(w, code, map[string]simpleResponse{"error": e})
}

func restSucceed(w http.ResponseWriter, code int, message string) {
	e := simpleResponse{Message: message, Code: code}
	ren.JSON(w, code, map[string]simpleResponse{"success": e})
}

var man = db.NewBinManager()
var ren *render.Render

// CreateHandler is a handler of (POST) /v1/{id}/create
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["id"]
	if !ok {
		restError(w, 400, "Invalid id")
		return
	}

	if bin := man.Create(v); bin == nil {
		restError(w, 500, fmt.Sprintf("Bin %s already exists.", v))
		return
	}

	restSucceed(w, 200, "Created "+v+" bin")
}

// ItemsHandler is a handler of (GET) /v1/{id}/items/{num}
func ItemsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		restError(w, 400, "Invalid id")
		return
	}

	bin := man.Bin(id)
	if bin == nil {
		restError(w, 404, fmt.Sprintf(`Bin "%s" not found`, id))
		return
	}

	num, ok := vars["num"]
	if !ok || num == "" {
		logs, err := bin.ReadLogs(0, 100)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if json.NewEncoder(w).Encode(logs); err != nil {
			restError(w, 500, fmt.Sprintf(`Failed to create JSON`))
		}
		return
	}

	index, err := strconv.Atoi(num)
	if err != nil {
		restError(w, 400, "Invalid index")
		return
	}

	res, err := bin.ReadLog(index)
	if err != nil {
		restError(w, 404, fmt.Sprintf(`Log "#%d" not found`, index))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if json.NewEncoder(w).Encode(res); err != nil {
		restError(w, 500, fmt.Sprintf(`Failed to create JSON`))
	}
}

// InHandler is a handler of (ANY) /v1/id/in/*
func InHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	bin := man.Bin(id)
	if bin == nil {
		restError(w, 404, fmt.Sprintf(`Bin "%s" not found`, id))
		return
	}
	req, err := db.NewRequest(time.Now(), r)
	if err != nil {
		restError(w, 500, "Failed to create JSON")
		return
	}
	bin.WriteLog(req)

	restSucceed(w, 200, r.RemoteAddr)
	println(vars["id"])
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	ren = render.New()

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
