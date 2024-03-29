package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
)

type EventGateway struct {
	EventService *pkg.EventService
}

type CreateEventInput struct {
	Name string `json:"name"`
}

func (e EventGateway) CreateEventHttp(w http.ResponseWriter, r *http.Request) {
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

	eventId, err := e.EventService.CreateEvent(input.Name, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to write event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

func (e EventGateway) DeleteEventHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["id"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = e.EventService.DeleteEvent(eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to delete event")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (e EventGateway) GetEventHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["id"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	event, err := e.EventService.GetEventForUser(eventId, userId)
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

type EventsOutput struct {
	Events []pkg.EventWithDetails `json:"events"`
}

func (e EventGateway) EventsHttp(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	events, err := e.EventService.GetEventsForUser(userId)
	if err != nil {
		logrus.WithError(err).Error("unable to get events for user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := EventsOutput{Events: events}
	err = WriteResponse(w, output)
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

func (e EventGateway) UpdateEventHttp(w http.ResponseWriter, r *http.Request) {
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

	err = e.EventService.UpdateEvent(eventId, userId, input.Name, input.Emails)

	if err != nil {
		if err != nil {
			logrus.WithError(err).Error("unable to update event")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
