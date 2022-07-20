package pkg

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const SessionCookieName = "sessionId"
const csrfHeader = "X-CSRF-TOKEN"

type SessionMiddleware struct {
	Database *Database
}

func (s SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO users shouldn't be here, replace with real auth
		if strings.HasPrefix(r.URL.Path, "/login") || r.URL.Path == "/login/google" || r.URL.Path == "/users" || r.URL.Path == "/users/create" {
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

		session, err := s.Database.ReadSession(sessionId)
		if err != nil {
			logrus.WithError(err).Error("unable to read session")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if session == nil || !session.Active {
			logrus.WithField("sessionId", session.Id).Warn("no session found or session not active")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		if r.URL.Path != "/csrf" {
			csrfToken := r.Header.Get(csrfHeader)
			if csrfToken == "" {
				logrus.Warn("request with no CSRF token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			if session.CsrfToken != csrfToken {
				logrus.WithField("expectedCSRF", session.CsrfToken).WithField("actualCSRF", csrfToken).Warn("request with invalid csrf token")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		user, err := s.Database.ReadUser(session.UserId)
		if err != nil {
			logrus.WithError(err).Error("unable to read user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if user == nil {
			logrus.WithField("sessionid", sessionId).Warn("no user found for session")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		newReq := r.WithContext(NewUserIdContext(r.Context(), user.Id))
		next.ServeHTTP(w, newReq)
	})
}
