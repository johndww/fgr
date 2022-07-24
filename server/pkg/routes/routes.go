package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
)

func Define(config pkg.Config) *mux.Router {
	db, err := pkg.NewDatabase(config)
	if err != nil {
		logrus.WithError(err).Fatal("could not create database")
	}

	baseRouter := mux.NewRouter()

	baseRouter.Use(mux.CORSMethodMiddleware(baseRouter))
	baseRouter.Use(pkg.CORSOriginMiddleware{}.Middleware)
	baseRouter.Use(pkg.RequestMiddleware{}.Middleware)

	userGw := UserGateway{UserService: &pkg.UserService{Database: db, Config: config}}

	apiRouter := baseRouter.PathPrefix("/api").Subrouter()
	v1UnauthenticatedRouterRouter := apiRouter.PathPrefix("/v1").Subrouter()

	v1UnauthenticatedRouterRouter.HandleFunc("/login/google", userGw.LoginGoogleHttp).Methods(http.MethodPost, http.MethodOptions)

	v1AuthenticatedRouter := apiRouter.PathPrefix("/v1").Subrouter()
	v1AuthenticatedRouter.Use(pkg.SessionMiddleware{Database: db}.Middleware)

	v1AuthenticatedRouter.HandleFunc("/csrf", userGw.Csrf).Methods(http.MethodGet, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/logout", userGw.LogoutHttp).Methods(http.MethodPost, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/users/me", userGw.CurrentUserHttp).Methods(http.MethodGet, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/users/event/{id}", userGw.EventUsersHttp).Methods(http.MethodGet, http.MethodOptions)

	eventGw := EventGateway{EventService: &pkg.EventService{Database: db}}
	v1AuthenticatedRouter.HandleFunc("/events/create", eventGw.CreateEventHttp).Methods(http.MethodPost, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{id}", eventGw.GetEventHttp).Methods(http.MethodGet, http.MethodOptions, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events", eventGw.EventsHttp).Methods(http.MethodGet, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{id}/update", eventGw.UpdateEventHttp).Methods(http.MethodPost, http.MethodOptions)

	giftGw := GiftRequestGateway{EventService: &pkg.EventService{Database: db}}
	v1AuthenticatedRouter.HandleFunc("/events/{eventId}/gift-requests", giftGw.GetGiftRequestsHttp).Methods(http.MethodGet, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{eventId}/gift-requests/create", giftGw.CreateGiftRequestHttp).Methods(http.MethodPost, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{eventId}/gift-requests/{id}/delete", giftGw.DeleteGiftRequestsHttp).Methods(http.MethodDelete, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{eventId}/gift-requests/{id}/release", giftGw.ReleaseGiftRequestsHttp).Methods(http.MethodPost, http.MethodOptions)
	v1AuthenticatedRouter.HandleFunc("/events/{eventId}/gift-requests/{id}/claim", giftGw.ClaimGiftRequestsHttp).Methods(http.MethodPost, http.MethodOptions)

	v1AdminRouter := v1AuthenticatedRouter.PathPrefix("").Subrouter()
	v1AdminRouter.Use(pkg.AdminMiddleware{Database: db}.Middleware)
	v1AdminRouter.HandleFunc("/users/create", userGw.CreateUserHttp).Methods(http.MethodPost, http.MethodOptions)
	v1AdminRouter.HandleFunc("/users", userGw.AllUsersHttp).Methods(http.MethodGet, http.MethodOptions)
	v1AdminRouter.HandleFunc("/login/user/{id}", userGw.AdminLoginHttp).Methods(http.MethodPost, http.MethodOptions)

	return baseRouter
}
