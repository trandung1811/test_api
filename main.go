package main

import (
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)
var rt = rate.Every(5*time.Second / 1)
var limiter = NewIPRateLimiter(rt, 1)
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/posts", middleware(http.HandlerFunc(getUploadFile))).Methods("POST")
	router.HandleFunc("/get", getRequest).Methods("GET")
	log.Println("server started and listening on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", router)
}


