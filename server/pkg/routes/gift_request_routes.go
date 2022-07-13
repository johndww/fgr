package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
)

type GiftRequestGateway struct {
}

func (u GiftRequestGateway) GetGiftRequestsHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	giftRequests := pkg.GetGiftRequestsForUser(eventId, userId)
	err = WriteResponse(w, giftRequests)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type GiftRequestInput struct {
	Name string `json:"name"`
}

func (u GiftRequestGateway) CreateGiftRequestHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["eventId"]

	input := GiftRequestInput{}
	err := ReadBody(r.Body, &input)

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	requestId, err := pkg.CreateGiftRequest(eventId, input.Name, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to create gift request")
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	err = WriteResponse(w, struct {
		RequestId string `json:"requestId"`
	}{
		RequestId: requestId,
	})

	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (u GiftRequestGateway) DeleteGiftRequestsHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["id"]
	eventId := vars["eventId"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pkg.DeleteGiftRequest(requestId, eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to delete gift request")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (u GiftRequestGateway) ReleaseGiftRequestsHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["id"]
	eventId := vars["eventId"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pkg.ReleaseGiftRequest(requestId, eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to release gift request")
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (u GiftRequestGateway) ClaimGiftRequestsHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	requestId := vars["id"]
	eventId := vars["eventId"]

	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = pkg.ClaimGiftRequest(requestId, eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to claim gift request")
		w.WriteHeader(http.StatusBadRequest)
	}
}
