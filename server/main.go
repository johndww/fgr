package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
	"server/pkg/routes"
)

func main() {
	logrus.Info("FGR server startup")

	config := pkg.ReadConfig()

	router := routes.Define(config)

	logrus.Infof("Startup complete listening on: " + config.ListenURL)
	logrus.Fatal(http.ListenAndServe(config.ListenURL, router))
}
