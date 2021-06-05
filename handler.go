package main

import (
	"encoding/json"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"net/http"
	"strings"
)

func middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		nLimiter := limiter.GetLimiter(r.RemoteAddr)
		if !nLimiter.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		//check header
		headerName:= r.Header.Get("name")
		json.NewEncoder(w).Encode(r)
		headerName = strings.TrimSpace(headerName)

		if headerName != auHeader {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing header")
			return
		}
		token:= r.Header.Get("x-access-token")
		json.NewEncoder(w).Encode(r)
		token = strings.TrimSpace(token)
		if !checkToken(token, fileName) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing header")
			return
		}
		next.ServeHTTP(w, r)
	})
}
func checkToken(token string, fileName string) bool {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		println(err.Error())
	}
	//cell, _ := f.GetCellValue("Sheet1", "A1")
	rows, err := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			if colCell == token {
				return true
			}
		}
	}
	return false
}
