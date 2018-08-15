package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type response struct {
}

func ResponseStatus(code int, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
}

func ResponseStatusWithJSON(code int, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Fatal("[API RESPONSE]: The website encountered an unexpected error.")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func InternalServerError(w http.ResponseWriter) {
	resError := struct {
		Code         int
		TextResponse string
	}{Code: http.StatusInternalServerError, TextResponse: "internal server error"}
	ResponseStatusWithJSON(http.StatusInternalServerError, resError, w)
}
