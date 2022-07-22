package pkg

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type AdminMiddleware struct {
	Database *Database
}

func (a AdminMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId, err := UserIdFromContext(r.Context())
		if err != nil {
			logrus.WithError(err).Error("unable to pull userid from context")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		user, err := a.Database.ReadUser(userId)
		if err != nil {
			logrus.WithError(err).Error("unable to read user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !user.Admin {
			logrus.WithField("userid", userId).Warn("user is not an admin and is attempting to view an admin page")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		logrus.WithField("userid", user.Id).Info("passed admin check")
		next.ServeHTTP(w, r)
	})
}
