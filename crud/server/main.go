package main

import (
	"log"

	"github.com/parinay/RESTfulGo/crud/controller"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("Starting the server...")

	controller.InitialMigration()
	controller.HandleRequests()
}
