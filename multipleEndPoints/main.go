package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

var Articles []Article

// homePage() to handle all request to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: homepage")
	json.NewEncoder(w).Encode(Articles)
}

//handleRequests() matches URL paths to defined function
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Hello world!"},
		Article{Title: "Hello again", Desc: "Article Description", Content: "Hello world!!"},
	}
	handleRequests()
}
