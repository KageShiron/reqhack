package controller

import (
	"encoding/json"
	"fmt"
	"github.com/KageShiron/reqhack/server/domain"
	"github.com/KageShiron/reqhack/server/infrastracture"
	"github.com/KageShiron/reqhack/server/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"time"
)

// NowFunc is time.Now. (when testing, NowFunc will be overwrited)
var NowFunc = time.Now

// RequestController represents Requests' controller
type RequestController struct {
	Request infrastracture.RequestRepository
	Bin     infrastracture.BinRepository
}

// NewRequestController returns new RequestController
func NewRequestController(req infrastracture.RequestRepository, bin infrastracture.BinRepository) *RequestController {
	return &RequestController{
		Request: req,
		Bin:     bin,
	}
}

// Items is a handler of (GET) /v1/{name}/items/{num}
func (rc *RequestController) Items(w http.ResponseWriter, r *http.Request) {
	println(r.URL.String())
	vars := mux.Vars(r)
	name, ok := vars["name"]
	r.ParseForm()
	secret := ""
	if sf := r.Form["secret"]; sf != nil {
		secret = sf[0]
	}

	if !ok {
		utils.RestError(w, 400, "Invalid id")
		return
	}

	println(secret)
	bin, err := rc.Bin.Get(name, secret)
	if err != nil {
		utils.RestError(w, 404, fmt.Sprintf(`Bin "%s" not found`, name))
		return
	}

	num, ok := vars["num"]
	if !ok || num == "" {
		//todo://
		logs, err := rc.Request.GetRange(bin.ID, 0, 100)
		if err != nil {
			utils.RestError(w, 500, fmt.Sprintf("Internal Error %s", err))
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if json.NewEncoder(w).Encode(logs); err != nil {
			utils.RestError(w, 500, fmt.Sprintf(`Failed to create JSON`))
		}
		return
	}

	index, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		utils.RestError(w, 400, "Invalid index")
		return
	}

	res, err := rc.Request.Get(bin.ID, index)
	if err != nil {
		utils.RestError(w, 404, fmt.Sprintf(`Log "#%d" not found`, index))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if json.NewEncoder(w).Encode(res); err != nil {
		utils.RestError(w, 500, fmt.Sprintf(`Failed to create JSON`))
	}
}

// In is a handler of (ANY) /v1/{name}/in/*
func (rc *RequestController) In(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	bin, err := rc.Bin.GetWithoutSecret(name)
	if err != nil {
		utils.RestError(w, 404, fmt.Sprintf(`Bin "%s" not found`, name))
		return
	}
	req, err := domain.NewRequest(NowFunc(), r)
	req.Bin = bin
	if err != nil {
		utils.RestError(w, 500, "Failed to create JSON")
		return
	}
	rc.Request.Add(req)
	utils.RestSucceedObject(w, 200, req)
}

// Body is a handler of (GET) /v1/{name}/items/{num}/body
func (rc *RequestController) Body(w http.ResponseWriter, r *http.Request) {
	println(r.URL.String())
	vars := mux.Vars(r)
	name, ok := vars["name"]
	if !ok {
		utils.RestError(w, 400, "Invalid id")
		return
	}

	r.ParseForm()
	secret := ""
	if sf := r.Form["secret"]; sf != nil {
		secret = sf[0]
	}

	bin, err := rc.Bin.Get(name,secret)
	if err != nil {
		utils.RestError(w, 404, fmt.Sprintf(`Bin "%s" not found`, name))
		return
	}

	num, ok := vars["num"]
	index, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		utils.RestError(w, 400, "Invalid index")
		return
	}

	res, err := rc.Request.Get(bin.ID, index)
	if err != nil {
		utils.RestError(w, 404, fmt.Sprintf(`Log "#%d" not found`, index))
		return
	}

	utils.RestBinary(w, 200, res.Body)
}
