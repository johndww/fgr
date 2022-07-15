package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
	"time"
)

type UserGateway struct {
}

func (u UserGateway) LoginHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user := pkg.GetUser(userId)

	if user == nil {
		logrus.WithField("userId", userId).Warn("user not found for login")
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("User not found: " + userId))
		if err != nil {
			logrus.WithError(err).Error("unable to write user not found error")
			return
		}
	}

	expiration := time.Now().Add(365 * 24 * time.Hour)

	sessionIdCookie := http.Cookie{
		Name:    pkg.SessionCookieName,
		Value:   userId,
		Path:    "/",
		Expires: expiration,
		Secure:  false, //TODO should be true when we use https
	}
	http.SetCookie(w, &sessionIdCookie)

	logrus.WithField("userId", userId).Info("logged user in")
}

func (u UserGateway) CurrentUserHttp(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user := pkg.GetUser(userId)

	err = WriteResponse(w, UserOutput{*user})
	if err != nil {
		logrus.WithError(err).Error("unable to marshal user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type UserOutput struct {
	pkg.User // will probably need to diverge to SelfUserOutput and UserOutput to not give all details about other users
}

type UserInput struct {
	Name string `json:"name"`
}

func (u UserGateway) CreateUserHttp(w http.ResponseWriter, r *http.Request) {
	input := UserInput{}
	err := ReadBody(r.Body, &input)

	if err != nil {
		logrus.WithError(err).Error("unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId := pkg.CreateUser(input.Name)
	w.WriteHeader(http.StatusCreated)
	err = WriteResponse(w, struct {
		UserId string `json:"userId"`
	}{
		UserId: userId,
	})

	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (u UserGateway) AllUsersHttp(w http.ResponseWriter, r *http.Request) {
	users := pkg.AllUsers()

	output := ExternalUsersOutput{}
	for _, user := range users {
		output.Users = append(output.Users, ExternalUser{Id: user.Id, Name: user.Name})
	}

	err := WriteResponse(w, output)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type ExternalUsersOutput struct {
	Users []ExternalUser `json:"users"`
}

type ExternalUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type EventUsersOutput struct {
	Users []EventUser `json:"users"`
}

type EventUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (u UserGateway) EventUsersHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	eventId := vars["id"]

	users := pkg.Users(eventId)

	output := EventUsersOutput{}
	for _, user := range users {
		output.Users = append(output.Users, EventUser{Id: user.Id, Name: user.Name, Email: user.Email})
	}

	err := WriteResponse(w, output)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
