package pkg

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/idtoken"
)

const userContextKey = "userContextKey"

func NewUserIdContext(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, userContextKey, userId)
}

func UserIdFromContext(ctx context.Context) (string, error) {
	userId, ok := ctx.Value(userContextKey).(string)
	if !ok {
		return "", errors.New("no user Id in context")
	}
	return userId, nil
}

type UserService struct {
	Database *Database
	Config   Config
}

func (u UserService) CreateUser(name string) (string, error) {
	user := User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: name + "@Email.com",
		Admin: false,
		Demo:  false,
	}

	err := u.Database.WriteUser(user)
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (u UserService) GetUser(id string) (*User, error) {
	return u.Database.ReadUser(id)
}

func (u UserService) AllUsers() ([]User, error) {
	users, err := u.Database.ReadUsers()
	if err != nil {
		logrus.WithError(err).Error("unable to fetch all users")
		return []User{}, err
	}
	return users, nil
}

func (u UserService) UsersForEvent(eventId string) ([]User, error) {
	return u.Database.ReadUsersForEvent(eventId)
}

func (u UserService) AdminLogin(userId string) (*Session, error) {
	user, err := u.Database.ReadUser(userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("could not find user to login to")
	}

	return u.Database.CreateSession(user.Id)
}

func (u UserService) GoogleLogin(token string) (*Session, error) {
	payload, err := idtoken.Validate(context.Background(), token, u.Config.GoogleAuthClientId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to validate google jwt")
	}

	existingUser, err := u.Database.MapExternalIdToUser(payload.Subject, GoogleAuthSource)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return u.Database.CreateSession(existingUser.Id)
	}

	// never seen this user before. create user in context

	newUser := User{
		Id:    uuid.New().String(),
		Name:  payload.Claims["name"].(string),
		Email: payload.Claims["email"].(string),
	}

	logrus.WithField("externalId", payload.Subject).WithField("newUserId", newUser.Id).Info("new google login user. creating user..")

	mapping := UserIdMapping{
		UserId:     newUser.Id,
		ExternalId: payload.Subject,
		Source:     GoogleAuthSource,
	}

	updatedUser, err := u.Database.WriteExternalUser(newUser, mapping)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new user from google login")
	}
	return u.Database.CreateSession(updatedUser.Id)
}

func (u UserService) GetCsrf(sessionId string) (string, error) {
	session, err := u.Database.ReadSession(sessionId)
	if err != nil {
		return "", err
	}

	//if !session.Active {
	//	err := errors.New("session is not active")
	//	logrus.WithError(err).WithField("sessionId", session.Id).WithField("userId", userId).Error("unable to get csrf token")
	//	return "", err
	//}
	return session.CsrfToken, nil
}

func (u UserService) DemoLogin() (*Session, error) {
	return u.Database.CreateSession(DemoUserId1)
}

type UserIdMapping struct {
	UserId     string
	ExternalId string
	Source     AuthSource
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Admin bool   `json:"admin"`
	Demo  bool   `json:"demo"`
}
