package routes

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"server/pkg"
	"time"
)

type UserGateway struct {
	UserService *pkg.UserService
}

type googleToken struct {
	Token string `json:"token"`
}

func (u UserGateway) LoginGoogleHttp(w http.ResponseWriter, r *http.Request) {
	input := googleToken{}
	err := ReadBody(r.Body, &input)

	if err != nil {
		logrus.WithError(err).Error("unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := u.UserService.GoogleLogin(input.Token)
	if err != nil {
		logrus.WithError(err).Error("unable to validate google id token")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.setCookie(w, userId)

	logrus.WithField("userId", userId).Info("logged user in with google")
}

func (u UserGateway) LoginHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	user, err := u.UserService.GetUser(userId)
	if err != nil {
		logrus.WithError(err).Error("unable to read user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		logrus.WithField("userId", userId).Warn("user not found for login")
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("User not found: " + userId))
		if err != nil {
			logrus.WithError(err).Error("unable to write user not found error")
			return
		}
	}

	u.setCookie(w, userId)

	logrus.WithField("userId", userId).Info("logged user in")
}

func (u UserGateway) setCookie(w http.ResponseWriter, userId string) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	sessionIdCookie := http.Cookie{
		Name:    pkg.SessionCookieName,
		Value:   userId,
		Path:    "/",
		Expires: expiration,
		Secure:  false, //TODO should be true when we use https
	}
	http.SetCookie(w, &sessionIdCookie)
}

func (u UserGateway) LogoutHttp(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionIdCookie := http.Cookie{
		Name:    pkg.SessionCookieName,
		Value:   "",
		Path:    "/",
		Expires: time.Now(),
		MaxAge:  0,
		Secure:  false, //TODO should be true when we use https
	}
	http.SetCookie(w, &sessionIdCookie)

	logrus.WithField("userId", userId).Info("logged user out")
}

func (u UserGateway) CurrentUserHttp(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := u.UserService.GetUser(userId)
	if err != nil {
		logrus.WithError(err).Error("unable to read user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil {
		logrus.WithError(err).Error("unable to find user for current user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	userId, err := u.UserService.CreateUser(input.Name)
	if err != nil {
		logrus.WithError(err).Error("unable to write user")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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
	users, err := u.UserService.AllUsers()
	if err != nil {
		logrus.WithError(err).Error("unable to read users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := ExternalUsersOutput{}
	for _, user := range users {
		output.Users = append(output.Users, ExternalUser{Id: user.Id, Name: user.Name})
	}

	err = WriteResponse(w, output)
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

	users, err := u.UserService.UsersForEvent(eventId)
	if err != nil {
		logrus.WithError(err).Error("unable to read users")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	output := EventUsersOutput{}
	for _, user := range users {
		output.Users = append(output.Users, EventUser{Id: user.Id, Name: user.Name, Email: user.Email})
	}

	err = WriteResponse(w, output)
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
