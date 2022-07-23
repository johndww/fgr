package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
	"server/pkg/routes"
)

const listenURL = "127.0.0.1:8085"

func main() {
	logrus.Info("FGR server startup")

	config := pkg.ReadConfig()

	router := routes.Define(config)

	//TODO change to localhost before prod
	logrus.Infof("Startup complete listening on: " + listenURL)
	logrus.Fatal(http.ListenAndServe(listenURL, router))
}
