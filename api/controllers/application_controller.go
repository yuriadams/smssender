package controllers

import (
	"fmt"
	"net/http"
)

type Headers map[string]string

func respondWith(w http.ResponseWriter, status int, headers Headers) {
	for k, v := range headers {
		w.Header().Set(k, v)
	}
	w.WriteHeader(status)
}

func RespondWithJSON(w http.ResponseWriter, response string) {
	respondWith(w, http.StatusOK, Headers{"Content-Type": "application/json"})
	fmt.Fprintf(w, response)
}
