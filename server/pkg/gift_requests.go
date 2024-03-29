package pkg

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type GiftRequestOutput struct {
	Id             string  `json:"id"`
	UserId         string  `json:"userId"`
	EventId        string  `json:"eventId"`
	Name           string  `json:"name"`
	Description    *string `json:"description"`
	IsAssigned     bool    `json:"isAssigned"` // will always be false if gift request is for the logged in user
	IsAssignedToMe bool    `json:"isAssignedToMe"`
}

func (e EventService) GetGiftRequestsForEvent(eventId string, userId string) ([]GiftRequestOutput, error) {
	allGiftRequests, err := e.Database.ReadGiftRequestsForEvent(eventId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read gift requests")
	}

	giftRequests := []GiftRequestOutput{}
	for _, request := range allGiftRequests {
		if request.UserId == userId {
			// user's gift
			giftRequests = append(giftRequests, GiftRequestOutput{
				Id:          request.Id,
				UserId:      request.UserId,
				EventId:     request.EventId,
				Name:        request.Name,
				Description: request.Description,
				// don't output details on isAssigned for logged in user's gifts
			})
		} else {
			// someone elses gift

			if request.AssignedUserId != nil && *request.AssignedUserId != userId {
				// hide already claimed gifts by others to add more mystery
				// should probably push this down to the DB, but lazy right now
				continue
			}

			giftRequests = append(giftRequests, GiftRequestOutput{
				Id:             request.Id,
				UserId:         request.UserId,
				EventId:        request.EventId,
				Name:           request.Name,
				Description:    request.Description,
				IsAssigned:     request.AssignedUserId != nil,
				IsAssignedToMe: request.AssignedUserId != nil && *request.AssignedUserId == userId,
			})
		}
	}
	return giftRequests, nil
}

func (e EventService) CreateGiftRequest(eventId string, name string, description *string, userId string) (string, error) {
	request := GiftRequest{
		Id:             uuid.New().String(),
		UserId:         userId,
		EventId:        eventId,
		Name:           name,
		Description:    description,
		AssignedUserId: nil,
	}

	err := e.Database.WriteGiftRequestEnsureEventMembership(request)
	return request.Id, err
}

func (e EventService) DeleteGiftRequest(requestId string, eventId string, userId string) error {
	return e.Database.DeleteGiftRequest(requestId, eventId, userId)
}

func (e EventService) ReleaseGiftRequest(requestId string, eventId string, userId string) error {
	return e.Database.ReleaseGiftRequest(requestId, eventId, userId)
}

func (e EventService) ClaimGiftRequest(requestId string, eventId string, userId string) error {
	return e.Database.ClaimGiftRequest(requestId, eventId, userId)
}

type GiftRequest struct {
	Id             string
	UserId         string
	EventId        string
	Name           string
	Description    *string
	AssignedUserId *string
}
