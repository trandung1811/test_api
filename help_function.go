package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")

		//check header
		var header = r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)
		if header != auHeader {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing header")
			return
		}
		//check token
		params := mux.Vars(r)
		recievedToken:= params["id"]
		if recievedToken != token {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
