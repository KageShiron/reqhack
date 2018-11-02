package controller

import (
	"fmt"
	"github.com/KageShiron/reqhack/server/usecase"
	"github.com/KageShiron/reqhack/server/utils"
	"github.com/gorilla/mux"
	"net/http"
)

// BinController represents Bin's controller
type BinController struct {
	Bin usecase.BinUsecase
}

// NewBinController returns new bin controller
func NewBinController(bin usecase.BinUsecase) *BinController {
	return &BinController{Bin: bin}
}

// Create is a handler of (POST) /v1/{name}/create
func (b *BinController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["name"]

	if !ok {
		utils.RestError(w, 400, "Invalid name")
		return
	}

	if _, err := b.Bin.Add(v); err != nil {
		utils.RestError(w, 500, fmt.Sprintf("Bin %s already exists.", v))
		return
	}

	utils.RestSucceed(w, 200, "Created "+v+" bin")
}
