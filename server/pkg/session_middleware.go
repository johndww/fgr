package pkg

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const SessionCookieName = "sessionId"

type SessionMiddleware struct {
	Database *Database
}

func (s SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO users shouldn't be here, replace with real auth
		if strings.HasPrefix(r.URL.Path, "/login") || r.URL.Path == "/users" || r.URL.Path == "/users/create" {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie(SessionCookieName)
		if err != nil {
			if r.URL.Path == "/users/me" {
				// dont log, using this for session timeouts
			} else {
				logrus.WithError(err).WithField("route", r.URL.Path).Warn("unable to read session cookie")
			}
			w.WriteHeader(http.StatusForbidden)
			return
		}

		sessionId := cookie.Value

		//TODO sessionId should map to userid. for now, its just the userid
		user, err := s.Database.ReadUser(sessionId)
		if err != nil {
			logrus.WithError(err).Error("unable to read user / session id")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if user == nil {
			logrus.WithField("sessionid", sessionId).Warn("invalid user")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		newReq := r.WithContext(NewUserIdContext(r.Context(), user.Id))
		next.ServeHTTP(w, newReq)
	})
}
