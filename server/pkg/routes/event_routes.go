package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
)

type EventGateway struct {
}

type CreateEventInput struct {
	Name string `json:"name"`
}

func (u EventGateway) CreateEventHttp(w http.ResponseWriter, r *http.Request) {
	input := CreateEventInput{}
	err := ReadBody(r.Body, &input)
	if err != nil {
		logrus.WithError(err).Error("")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	eventId := pkg.CreateEvent(input.Name, userId)
	w.WriteHeader(http.StatusCreated)
	err = WriteResponse(w, struct {
		EventId string `json:"eventId"`
	}{
		EventId: eventId,
	})

	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (u EventGateway) GetEventHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["id"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := pkg.GetEventForUser(eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to get event for user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if event == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = WriteResponse(w, event)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (u EventGateway) EventsHttp(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events := pkg.GetEventsForUser(userId)
	err = WriteResponse(w, events)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type UpdateEventInput struct {
	Name   string   `json:"name"`
	Emails []string `json:"emails"`
}

func (u EventGateway) UpdateEventHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["id"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	input := UpdateEventInput{}
	err = ReadBody(r.Body, &input)
	if err != nil {
		logrus.WithError(err).Error("")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	err = pkg.UpdateEvent(eventId, userId, input.Name, input.Emails)

	if err != nil {
		if err != nil {
			logrus.WithError(err).Error("unable to update event")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
