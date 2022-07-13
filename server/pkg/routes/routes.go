package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"server/pkg"
)

func Define(router *mux.Router) {
	router.Use(pkg.SessionMiddleware{}.Middleware)
	router.Use(pkg.RequestMiddleware{}.Middleware)

	userGw := UserGateway{}
	router.HandleFunc("/login/{id}", userGw.LoginHttp).Methods(http.MethodPost)
	router.HandleFunc("/users/{id}", userGw.CurrentUserHttp).Methods(http.MethodGet)
	router.HandleFunc("/users/create", userGw.CreateUserHttp).Methods(http.MethodPost)
	router.HandleFunc("/users", userGw.AllUsersHttp).Methods(http.MethodGet)
	router.HandleFunc("/users/event/{id}", userGw.EventUsersHttp).Methods(http.MethodGet)

	eventGw := EventGateway{}
	router.HandleFunc("/events/create", eventGw.CreateEventHttp).Methods(http.MethodPost)
	router.HandleFunc("/events/{id}", eventGw.GetEventHttp).Methods(http.MethodGet)
	router.HandleFunc("/events", eventGw.EventsHttp).Methods(http.MethodGet)
	router.HandleFunc("/events/{id}/update", eventGw.UpdateEventHttp).Methods(http.MethodPost)

	giftGw := GiftRequestGateway{}
	router.HandleFunc("/events/{eventId}/gift-requests", giftGw.GetGiftRequestsHttp).Methods(http.MethodGet)
	router.HandleFunc("/events/{eventId}/gift-requests/create", giftGw.CreateGiftRequestHttp).Methods(http.MethodPost)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/delete", giftGw.DeleteGiftRequestsHttp).Methods(http.MethodDelete)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/release", giftGw.ReleaseGiftRequestsHttp).Methods(http.MethodPost)
	router.HandleFunc("/events/{eventId}/gift-requests/{id}/claim", giftGw.ClaimGiftRequestsHttp).Methods(http.MethodPost)
}
