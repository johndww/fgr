package pkg

import (
	"context"
	"errors"
	"github.com/google/uuid"
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

func CreateUser(name string) string {
	user := User{
		Id:    uuid.New().String(),
		Name:  name,
		Email: name + "@Email.com",
	}

	allUsers = append(allUsers, user)
	return user.Id
}

func GetUser(id string) *User {
	for _, user := range allUsers {
		if user.Id == id {
			return &user
		}
	}
	return nil
}

func AllUsers() []User {
	return allUsers
}

func Users(eventId string) []User {
	var users []User
	for _, membership := range memberships {
		if membership.EventId == eventId {
			for _, user := range allUsers {
				if user.Id == membership.UserId {
					users = append(users, user)
				}
			}
		}
	}
	return users
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var allUsers = []User{
	{
		Id:    "1",
		Name:  "John",
		Email: "john.d.wright@gmail.com",
	},
	{
		Id:    "2",
		Name:  "Haritha",
		Email: "haritha.tapa@gmail.com",
	},
	{
		Id:    "3",
		Name:  "Sue",
		Email: "paubsue@gmail.com",
	},
	{
		Id:    "4",
		Name:  "Bruce",
		Email: "bruce.d.wright@gmail.com",
	},
}
