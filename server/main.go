package main

import (
	"github.com/KageShiron/reqhack/server/infrastracture"
	"github.com/KageShiron/reqhack/server/infrastracture/controller"
	"github.com/KageShiron/reqhack/server/infrastracture/mysql"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(false)

	sql := infrastracture.NewSQLHandler()
	br := mysql.NewMysqlBinRepository(*sql)
	rr := mysql.NewMysqlRequestRepository(*sql)
	bc := controller.NewBinController(br)
	rc := controller.NewRequestController(rr, br)
	router.Methods("POST").Path("/v1/{name}/{_:create/?}").HandlerFunc(bc.Create)
	router.Methods("GET").Path("/v1/{name}/items").HandlerFunc(rc.Items)
	router.Methods("GET").Path("/v1/{name}/items/{num:[0-9]*(?:\\/)?}").HandlerFunc(rc.Items)
	router.Methods("GET").Path("/v1/{name}/items/{num:[0-9]+}/body").HandlerFunc(rc.Body)
	router.Path("/v1/{name}/{_:in(?:/.*|$)}").HandlerFunc(rc.In)
	srv := &http.Server{
		Addr:    ":8081",
		Handler: router, // Pass our instance of gorilla/router in.
	}

	srv.ListenAndServe()
}
