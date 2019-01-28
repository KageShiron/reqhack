package controller

import (
	"fmt"
	"github.com/KageShiron/reqhack/server/infrastracture"
	"github.com/KageShiron/reqhack/server/utils"
	"github.com/gorilla/mux"
	"crypto/rand"
	"math/big"
	"net/http"
)

// BinController represents Bin's controller
type BinController struct {
	Bin infrastracture.BinRepository
}

// NewBinController returns new bin controller
func NewBinController(bin infrastracture.BinRepository) *BinController {
	return &BinController{Bin: bin}
}

var rs1Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var rs1max = big.NewInt(int64(len(rs1Letters)))

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		num,_ := rand.Int(rand.Reader,rs1max)
		b[i] = rs1Letters[num.Int64()]
	}
	return string(b)
}

// Create is a handler of (POST) /v1/{name}/create
func (b *BinController) Create(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	v, ok := vars["name"]

	if !ok {
		utils.RestError(w, 400, "Invalid name")
		return
	}

	secret := ""
	r.ParseForm()
	if isPrivate := r.Form["isPrivate"]; isPrivate != nil && isPrivate[0] == "true" {
		secret = randString(64)
	}

	if _, err := b.Bin.Add(v, secret); err != nil {
		utils.RestError(w, 500, fmt.Sprintf("Bin %s already exists.", v))
		return
	}

	utils.RestSucceedWithObject(w, 200, "Created "+v+" bin", map[string]interface{}{"secret": secret})
}
