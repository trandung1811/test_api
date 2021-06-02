package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Post struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Token string `json:"token"`
}

var posts []Post


const token = "455"
const auHeader = "Buscu-bot"
func main() {
	//setupGoGuardian()
	router := mux.NewRouter()
	//posts = append(posts, Post{ID: "1", Title: "My first post", Body:"This is the content of my first post", Token: "1234"})
	//posts = append(posts, Post{ID: "2", Title: "My second post", Body:"This is the content of my second post", Token: "5678"})
	router.HandleFunc("/get",getPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", middleware(http.HandlerFunc(createPost))).Methods("POST")
	router.HandleFunc("/get/{id}", getPost).Methods("GET")
	router.HandleFunc("/put/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/delete/{id}", deletePost).Methods("DELETE")
	log.Println("server started and listening on http://127.0.0.1:8000")
	http.ListenAndServe(":8000", router)
}


