//go:build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

// Runs go mod download and then installs the binary.
func Build() error {
	err := buildServer()
	if err != nil {
		logrus.WithError(err).Fatal("unable to build server")
	}

	err = buildUi()
	if err != nil {
		logrus.WithError(err).Fatal("unable to build ui")
	}

	return nil
}

func Deploy() error {
	mg.Deps(Build)

	err := sh.Run("sftp fgr@164.92.73.41:/var/www/html <<< $'put -r ../ui/dist'")
	if err != nil {
		logrus.WithError(err).Fatal("unable to deploy ui")
	}

	err = deployServer()
	if err != nil {
		logrus.WithError(err).Fatal("unable to deploy server")
	}
	return nil
}

func deployServer() error {
	err := sh.Run("ssh", "-A", "-t", "fgr@164.92.73.41", "\"sudo systemctl stop fgr\"")
	if err != nil {
		return errors.Wrap(err, "unable to stop backend service")
	}

	err = sh.Run("sftp", "fgr@164.92.73.41:/home/fgr", "<<<", "$'put -r build/'")
	if err != nil {
		return errors.Wrap(err, "unable to upload server artifacts")
	}

	err = sh.Run("ssh", "-A", "-t", "fgr@164.92.73.41", "\"sudo systemctl start fgr\"")
	if err != nil {
		return errors.Wrap(err, "unable to start backend service")
	}
	return nil
}

func buildUi() error {
	if err := sh.Run("npm", "run", "build", "--prefix", "../ui "); err != nil {
		return err
	}

	return nil
}

func buildServer() error {
	//GOOS=linux GOARCH=amd64 CGO_ENABLED=0
	if err := sh.Run("go", "build", "-o", "build/fgr"); err != nil {
		return err
	}

	err := sh.Run("cp", ".env", "build")
	if err != nil {
		return err
	}

	return sh.Run("cp", envFile(), "build")
}

func envFile() string {
	envFile := ".env.local"
	if os.Getenv("BEHAVIOR") == "prod" {
		envFile = ".env.prod"
	}
	return envFile
}
