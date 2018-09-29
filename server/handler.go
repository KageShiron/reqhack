package main

import (
	"encoding/json"
	"fmt"
	"github.com/KageShiron/reqhack/server/db"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"
)

var man = db.NewMysqlBinManager()
var ren = render.New()

type simpleResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func restError(w http.ResponseWriter, code int, message string) error {
	e := simpleResponse{Message: message, Code: code}
	return ren.JSON(w, code, map[string]simpleResponse{"error": e})
}

func restSucceed(w http.ResponseWriter, code int, message string) error {
	e := simpleResponse{Message: message, Code: code}
	return ren.JSON(w, code, map[string]simpleResponse{"success": e})
}

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
