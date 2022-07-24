package pkg

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

const behaviorKey = "BEHAVIOR"

func ReadConfig() Config {
	behavior := os.Getenv(behaviorKey)
	if behavior == "" {
		behavior = "dev"
		err := godotenv.Load(".env.local")
		if err != nil {
			logrus.WithError(err).Fatal("cannot load behavior dev config")
		}
		logrus.Info("loaded dev config")
	} else {
		err := godotenv.Load(".env." + behavior)
		if err != nil {
			logrus.WithError(err).Fatal("cannot load env props for behavior: " + behavior)
		}
		logrus.Infof("loaded %s config", behavior)
	}

	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Fatal()
	}

	return Config{
		DbConnectString:    mustEnv("DB_CONNECT_STRING"),
		ListenURL:          mustEnv("LISTEN_URL"),
		GoogleAuthClientId: mustEnv("GOOGLE_AUTH_CLIENT_ID"),
	}
}

func mustEnv(key string) string {
	result, ok := os.LookupEnv(key)
	if !ok {
		panic("no env var set for: " + key)
	}
	return result
}

type Config struct {
	DbConnectString    string
	ListenURL          string
	GoogleAuthClientId string
}
