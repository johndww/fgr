package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/pkg/routes"
)

func main() {
	log.Println("FGR server startup")

	router := mux.NewRouter()
	routes.Define(router)

	log.Fatal(http.ListenAndServe(":80", router))
}
