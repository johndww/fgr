package pkg

import (
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
