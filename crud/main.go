package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
	Author  string `json:"Author"`
}

var Articles []Article

// homePage() to handle all request to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

func returnArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnArticles")

	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleArticle")

	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// createNewArticle() creates a new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Printf("Post Reuqst Body %v\n", string(reqBody))

	var article Article

	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Fatalf("Failed to Unmarshal request body %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Create new Article Failed due to some error"))
	}
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200  Create article successful"))
	json.NewEncoder(w).Encode(article)
}

// createNewArticleV2() creates a new article
func createNewArticleV2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	reqBody, _ := ioutil.ReadAll(r.Body)

	fmt.Fprintf(w, "%+v", string(reqBody))
	fmt.Printf("Post Reuqst Body %v\n", string(reqBody))

	var article Article

	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		log.Fatalf("Failed to Unmarshal request body %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Create new Article Failed due to some error"))
	}
	Articles = append(Articles, article)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Create new Article Successful"))
	json.NewEncoder(w).Encode(article)
}

// deleteArticle() deletes an article
func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")

	vars := mux.Vars(r)

	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200 - Delete article successful"))
		} else {
			log.Println("Record not found")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Delete article failed due to some error"))
		}
	}
}

// updateArticle() update existing article id
func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")

	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)

	id := vars["id"]
	for index, article := range Articles {
		if article.Id == id {
			var article Article
			json.Unmarshal(reqBody, &article)
			Articles[index].Id = article.Id
			Articles[index].Title = article.Title
			Articles[index].Desc = article.Desc
			Articles[index].Content = article.Content
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("200  Update article successful"))
			json.NewEncoder(w).Encode(article)

		} else {
			log.Println("Record not found")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Delete article failed due to some error"))
		}
	}

}

//handleRequests() matches URL paths to defined function
func handleRequests() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/v1/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/api/v2/article", createNewArticleV2).Methods("POST")
	myRouter.HandleFunc("/api/v1/article/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/api/v1", homePage)
	myRouter.HandleFunc("/api/v1/article", returnArticles)
	myRouter.HandleFunc("/api/v1/article/{id}", deleteArticle)
	myRouter.HandleFunc("/api/v1/article/{id}", returnSingleArticle)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Hello world!"},
		Article{Id: "2", Title: "Hello again", Desc: "Article Description", Content: "Hello world!!"},
	}
	handleRequests()
}
