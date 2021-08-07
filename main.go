package main

import (
	"log"
	"net/http"
	"post/routes"
	"post/services"
	"post/utility"
)

func main() {
	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))
	//_ = utility.GetConnection()
}
