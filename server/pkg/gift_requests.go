package pkg

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GiftRequestOutput struct {
	Id         string `json:"id"`
	UserId     string `json:"userId"`
	EventId    string `json:"eventId"`
	Name       string `json:"name"`
	IsAssigned bool   `json:"isAssigned"` // will always be false if gift request is for the logged in user
}

func GetGiftRequestsForUser(eventId string, userId string) []GiftRequestOutput {
	var giftRequests []GiftRequestOutput
	for _, request := range allGiftRequests {
		if request.EventId == eventId {
			if request.UserId == userId {
				// user's gift
				giftRequests = append(giftRequests, GiftRequestOutput{
					Id:      request.Id,
					UserId:  request.UserId,
					EventId: request.EventId,
					Name:    request.Name,
					// don't output details on isAssigned for logged in user's gifts
				})
			} else {
				// someone elses gift
				giftRequests = append(giftRequests, GiftRequestOutput{
					Id:         request.Id,
					UserId:     request.UserId,
					EventId:    request.EventId,
					Name:       request.Name,
					IsAssigned: request.AssignedUserId != "",
				})
			}
		}
	}
	return giftRequests
}

func CreateGiftRequest(eventId string, name string, userId string) (string, error) {
	isMember := false
	for _, membership := range memberships {
		if membership.EventId == eventId && membership.UserId == userId {
			isMember = true
		}
	}

	if !isMember {
		return "", errors.New("not a member of the event specified (or that event doesn't exist")
	}

	request := GiftRequest{
		Id:             uuid.New().String(),
		UserId:         userId,
		EventId:        eventId,
		Name:           name,
		AssignedUserId: "",
	}

	allGiftRequests = append(allGiftRequests, &request)
	return request.Id, nil
}

func DeleteGiftRequest(requestId string, eventId string, userId string) error {
	indexOfGiftRequest := -1
	for i, request := range allGiftRequests {
		if request.Id == requestId {
			if request.EventId != eventId {
				return errors.New("trying to delete a request from a different event")
			}
			if request.UserId != userId {
				return errors.New("trying to delete a request from someone else")
			}

			indexOfGiftRequest = i
		}
	}

	if indexOfGiftRequest == -1 {
		return errors.New("invalid gift request id")
	}

	allGiftRequests = append(
		allGiftRequests[:indexOfGiftRequest],
		allGiftRequests[indexOfGiftRequest+1:]...,
	)
	return nil
}

func ReleaseGiftRequest(requestId string, eventId string, userId string) error {
	released := false
	for _, request := range allGiftRequests {
		if request.Id == requestId {
			if request.EventId != eventId {
				return errors.New("trying to release a request from a different event")
			}
			if request.UserId == userId {
				return errors.New("trying to release a request from your own gift request")
			}
			if request.AssignedUserId != userId {
				return errors.New("trying to release a request claimed by someone else")
			}

			request.AssignedUserId = ""
			released = true
		}
	}

	if !released {
		return errors.New("could not find requestId: " + requestId)
	}
	return nil
}

func ClaimGiftRequest(requestId string, eventId string, userId string) error {
	claimed := false
	for _, request := range allGiftRequests {
		if request.Id == requestId {
			if request.EventId != eventId {
				return errors.New("trying to claim a request from a different event")
			}
			if request.UserId == userId {
				return errors.New("trying to claim a request from your own gift request")
			}
			if request.AssignedUserId != "" {
				return errors.New("trying to claim a request that is already claimed")
			}

			request.AssignedUserId = userId
			claimed = true
		}
	}

	if !claimed {
		return errors.New("could not find requestId: " + requestId)
	}
	return nil
}

type GiftRequest struct {
	Id             string
	UserId         string
	EventId        string
	Name           string
	AssignedUserId string
}

var allGiftRequests = []*GiftRequest{
	{
		Id:             "111",
		UserId:         "1",
		EventId:        "1",
		Name:           "Xbox",
		AssignedUserId: "2",
	},
	{
		Id:             "112",
		UserId:         "1",
		EventId:        "1",
		Name:           "Ps6",
		AssignedUserId: "",
	},
	{
		Id:             "113",
		UserId:         "1",
		EventId:        "1",
		Name:           "Gamecube",
		AssignedUserId: "",
	},
	{
		Id:             "121",
		UserId:         "1",
		EventId:        "2",
		Name:           "Kite",
		AssignedUserId: "",
	},
	{
		Id:             "211",
		UserId:         "2",
		EventId:        "1",
		Name:           "Shoes",
		AssignedUserId: "",
	},
	{
		Id:             "212",
		UserId:         "2",
		EventId:        "1",
		Name:           "Ski goggles",
		AssignedUserId: "1",
	},
	{
		Id:             "213",
		UserId:         "2",
		EventId:        "1",
		Name:           "Hat",
		AssignedUserId: "",
	},
	{
		Id:             "221",
		UserId:         "2",
		EventId:        "2",
		Name:           "Purse",
		AssignedUserId: "",
	},
	{
		Id:             "222",
		UserId:         "2",
		EventId:        "2",
		Name:           "Skis",
		AssignedUserId: "",
	},
	{
		Id:             "7",
		UserId:         "3",
		EventId:        "1",
		Name:           "Skis",
		AssignedUserId: "",
	},
	{
		Id:             "8",
		UserId:         "4",
		EventId:        "1",
		Name:           "wine",
		AssignedUserId: "",
	},
	{
		Id:             "9",
		UserId:         "4",
		EventId:        "1",
		Name:           "beer",
		AssignedUserId: "3",
	},
}
