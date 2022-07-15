package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/pkg"
)

func Define(router *mux.Router) {
	router.Use(mux.CORSMethodMiddleware(router))
	router.Use(pkg.CORSOriginMiddleware{}.Middleware)
	router.Use(pkg.SessionMiddleware{}.Middleware)
	router.Use(pkg.RequestMiddleware{}.Middleware)

	userGw := UserGateway{}
	router.HandleFunc("/login/{id}", userGw.LoginHttp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users/me", userGw.CurrentUserHttp).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/create", userGw.CreateUserHttp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/users", userGw.AllUsersHttp).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/users/event/{id}", userGw.EventUsersHttp).Methods(http.MethodGet, http.MethodOptions)

	eventGw := EventGateway{}
	router.HandleFunc("/events/create", eventGw.CreateEventHttp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/events/{id}", eventGw.GetEventHttp).Methods(http.MethodGet, http.MethodOptions, http.MethodOptions)
	router.HandleFunc("/events", eventGw.EventsHttp).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/events/{id}/update", eventGw.UpdateEventHttp).Methods(http.MethodPost, http.MethodOptions)

	giftGw := GiftRequestGateway{}
	router.HandleFunc("/events/{eventId}/gift-requests", giftGw.GetGiftRequestsHttp).Methods(http.MethodGet, http.MethodOptions)
	router.HandleFunc("/events/{eventId}/gift-requests/create", giftGw.CreateGiftRequestHttp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/delete", giftGw.DeleteGiftRequestsHttp).Methods(http.MethodDelete, http.MethodOptions)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/release", giftGw.ReleaseGiftRequestsHttp).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/claim", giftGw.ClaimGiftRequestsHttp).Methods(http.MethodPost, http.MethodOptions)
}
