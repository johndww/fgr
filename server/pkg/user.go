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
}

func (u UserService) CreateUser(name string) (string, error) {
	user := User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: name + "@Email.com",
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

//TODO pull out to config. not really a secret, but still shouldn't be committed
const aud = "186100627326-iqnh1vj4bbbse1i1qh24p1br61c9hgjh.apps.googleusercontent.com"

func (u UserService) GoogleLogin(token string) (string, error) {
	payload, err := idtoken.Validate(context.Background(), token, aud)
	if err != nil {
		return "", errors.Wrap(err, "unable to validate google jwt")
	}

	user, err := u.Database.MapExternalIdToUser(payload.Subject, GoogleAuthSource)
	if err != nil {
		return "", err
	}

	if user != nil {
		return user.Id, nil
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

	err = u.Database.WriteExternalUser(newUser, mapping)
	if err != nil {
		return "", errors.Wrap(err, "unable to create new user from google login")
	}
	return newUser.Id, nil
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
}
