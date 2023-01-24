package pkg

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RequestMiddleware struct {
}

func (s RequestMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.WithField("path", r.URL.Path).WithField("method", r.Method).Info("processing request")
		next.ServeHTTP(w, r)
	})
}

var demoUserApis = []string{
	"/api/v1/csrf",
	"/api/v1/users/logout",
	"/api/v1/users/me",
	"/api/v1/users/event/{id}",
	"/api/v1/events/{id}",
	"/api/v1/events",
	"/api/v1/events/{eventId}/gift-requests",
}

type DemoUserMiddleware struct {
}

func (s DemoUserMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path, err := mux.CurrentRoute(r).GetPathTemplate()
		if err != nil {
			logrus.WithError(err).Error("unable to get path template")
			return
		}

		if !isValidDemoPath(path) {
			logrus.WithField("api", path).Error("demo user attempted to access non-demo api")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func isValidDemoPath(path string) bool {
	for _, allowedPath := range demoUserApis {
		if allowedPath == path {
			return true
		}
	}

	return false
}
