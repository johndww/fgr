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

	logrus.Infof("Startup complete listening TLS on: " + "127.0.0.1:8085")
	logrus.Fatal(http.ListenAndServeTLS("127.0.0.1:8085", "../certs/simplegift_app.chained.crt", "../certs/simplegift.app.key", router))
}
