package pkg

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

//
//var allUsers = []User{
//	{
//		Id:    "1",
//		Name:  "John",
//		Email: "john.d.wright@gmail.com",
//	},
//	{
//		Id:    "2",
//		Name:  "Haritha",
//		Email: "haritha.tapa@gmail.com",
//	},
//	{
//		Id:    "3",
//		Name:  "Sue",
//		Email: "paubsue@gmail.com",
//	},
//	{
//		Id:    "4",
//		Name:  "Bruce",
//		Email: "bruce.d.wright@gmail.com",
//	},
//}
