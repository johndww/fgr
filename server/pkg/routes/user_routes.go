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
	Config      pkg.Config
}

type googleToken struct {
	Token string `json:"token"`
}

func (u UserGateway) DevAdminLogin(w http.ResponseWriter, r *http.Request) {
	u.AdminLoginHttp(w, r)
}

func (u UserGateway) LoginGoogleHttp(w http.ResponseWriter, r *http.Request) {
	input := googleToken{}
	err := ReadBody(r.Body, &input)

	if err != nil {
		logrus.WithError(err).Error("unable to read request body")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	session, err := u.UserService.GoogleLogin(input.Token)
	if err != nil {
		logrus.WithError(err).Error("unable to validate google id token")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	u.setCookie(w, session.Id)

	logrus.WithField("userId", session.UserId).Info("logged user in with google")
}

func (u UserGateway) AdminLoginHttp(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	session, err := u.UserService.AdminLogin(userId)
	if err != nil {
		logrus.WithError(err).Error("unable to admin login")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	u.setCookie(w, session.Id)

	logrus.WithField("userId", userId).Info("logged admin user in")
}

func (u UserGateway) setCookie(w http.ResponseWriter, userId string) {
	sessionIdCookie := http.Cookie{
		Name:   pkg.SessionCookieName,
		Value:  userId,
		Path:   "/",
		MaxAge: 30 * 24 * 60 * 60, // one month
		Secure: u.Config.Behavior == "prod",
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
		Secure:  u.Config.Behavior == "prod",
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

type CsrfResponse struct {
	Token string `json:"token"`
}

func (u UserGateway) Csrf(w http.ResponseWriter, r *http.Request) {
	userId, err := pkg.UserIdFromContext(r.Context())
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	csrf, err := u.UserService.GetCsrf(userId)
	if err != nil {
		logrus.WithError(err).Error("unable to pull userid from context")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = WriteResponse(w, CsrfResponse{Token: csrf})
	if err != nil {
		logrus.WithError(err).Error("unable to write response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
