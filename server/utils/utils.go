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

// RestSucceedWithObject returns success json and object
func RestSucceedWithObject(w http.ResponseWriter, code int, message string, object map[string]interface{}) error {
	e := simpleResponse{Message: message, Code: code}
	if object == nil {
		object = map[string]interface{}{}
	}
	object["success"] = e
	return ren.JSON(w, code, object)
}

// RestSucceed returns success json
func RestSucceed(w http.ResponseWriter, code int, message string) error {
	return RestSucceedWithObject(w, code, message, nil)
}

// RestSucceedObject returns success json with a result object
func RestSucceedObject(w http.ResponseWriter, code int, obj interface{}) error {
	e := map[string]interface{}{"result": obj, "code": code}
	return ren.JSON(w, code, map[string]interface{}{"success": e})
}
