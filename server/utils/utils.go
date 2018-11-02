package utils

import (
	"github.com/unrolled/render"
	"net/http"
)

type simpleResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

var ren = render.New()

// RestError responses error json
func RestError(w http.ResponseWriter, code int, message string) error {
	e := simpleResponse{Message: message, Code: code}
	return ren.JSON(w, code, map[string]simpleResponse{"error": e})
}

// RestSucceed returns success json
func RestSucceed(w http.ResponseWriter, code int, message string) error {
	e := simpleResponse{Message: message, Code: code}
	return ren.JSON(w, code, map[string]simpleResponse{"success": e})
}
