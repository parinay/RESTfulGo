package main

import (
	"fmt"
	"log"
	"net/http"
)

// homePage() to handle all request to root URL
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}

//handleRequests() matches URL paths to defined function
func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}
func main() {
	handleRequests()
}
