package pkg

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

const SessionCookieName = "sessionId"

type SessionMiddleware struct {
}

func (s SessionMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/login") {
			next.ServeHTTP(w, r)
			return
		}

		cookie, err := r.Cookie(SessionCookieName)
		if err != nil {
			logrus.WithError(err).Error("unable to read session cookie")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		sessionId := cookie.Value

		//TODO sessionId should map to userid. for now, its just the userid
		user := GetUser(sessionId)
		if user == nil {
			logrus.WithField("sessionid", sessionId).Warn("invalid user")
			w.WriteHeader(http.StatusForbidden)
			return
		}

		newReq := r.WithContext(NewUserIdContext(r.Context(), user.Id))
		next.ServeHTTP(w, newReq)
	})
}
