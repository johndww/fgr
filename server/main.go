package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
	"server/pkg/conversion"
	"server/pkg/routes"
)

func main() {
	logrus.Info("FGR server startup")

	config := pkg.ReadConfig()

	db, err := pkg.NewDatabase(config)
	if err != nil {
		logrus.WithError(err).Fatal("could not create database")
	}

	convRunner := conversion.Runner{Database: db}
	err = convRunner.RunConversions()
	if err != nil {
		logrus.WithError(err).Fatal("unable to run conversions")
	}

	router := routes.Define(config, db)

	if config.Behavior == "dev" {
		logrus.Infof("Dev Startup complete listening on: " + "127.0.0.1:8085")
		logrus.Fatal(http.ListenAndServe("127.0.0.1:8085", router))
	} else {
		logrus.Infof("Prod Startup complete listening TLS on: " + "127.0.0.1:8085")
		logrus.Fatal(http.ListenAndServeTLS("127.0.0.1:8085", "../certs/simplegift_app.chained.crt", "../certs/simplegift_app.key", router))
	}
}
