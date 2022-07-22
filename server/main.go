package main

import (
	"log"
	"net/http"
	"server/pkg/routes"
)

func main() {
	log.Println("FGR server startup")

	router := routes.Define()

	//TODO change to localhost before prod
	log.Fatal(http.ListenAndServe("127.0.0.1:80", router))
}
