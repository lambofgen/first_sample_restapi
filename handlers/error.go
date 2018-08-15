package handlers

import (
	"geon/api3/response"
	"net/http"
)

func errHandler(err *error, w http.ResponseWriter) bool {
	if *err != nil {
		response.InternalServerError(w)
		return true
	}
	return false
}
