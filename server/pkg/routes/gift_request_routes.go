package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
)

type GiftRequestGateway struct {
	EventService *pkg.EventService
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

	giftRequests, err := u.EventService.GetGiftRequestsForEvent(eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to get gift requests")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = WriteResponse(w, struct {
		Gifts []pkg.GiftRequestOutput `json:"gifts"`
	}{
		Gifts: giftRequests,
	})

	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type GiftRequestInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
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

	requestId, err := u.EventService.CreateGiftRequest(eventId, input.Name, input.Description, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to create gift request")
		w.WriteHeader(http.StatusBadRequest)
		return
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

	err = u.EventService.DeleteGiftRequest(requestId, eventId, userId)
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

	err = u.EventService.ReleaseGiftRequest(requestId, eventId, userId)
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

	err = u.EventService.ClaimGiftRequest(requestId, eventId, userId)
	if err != nil {
		logrus.WithError(err).Error("unable to claim gift request")
		w.WriteHeader(http.StatusBadRequest)
	}
}
