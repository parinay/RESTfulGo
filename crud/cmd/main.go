package main

import (
	"log"

	"github.com/parinay/RESTfulGo/crud/controller"
)

func main() {
	log.Println("Starting the server...")
	controller.HandleRequests()
}

func init() {

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	controller.Articles = []controller.Article{
		controller.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Hello world!"},
		controller.Article{Id: "2", Title: "Hello again", Desc: "Article Description", Content: "Hello world!!"},
	}
}
