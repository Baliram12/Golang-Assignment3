package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Article management system"},
	}
	fmt.Println("Endpoint Hit:AllArticles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Article management system")
}
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}
func main() {
	handleRequests()
}
