package pkg

import (
	"github.com/google/uuid"
)

type EventService struct {
	Database *Database
}

func (e EventService) CreateEvent(name string, ownerUserId string) (string, error) {
	event := Event{
		Id:          uuid.New().String(),
		Name:        name,
		OwnerUserId: ownerUserId,
	}

	membership := Membership{
		Id:      uuid.New().String(),
		EventId: event.Id,
		UserId:  ownerUserId,
	}

	err := e.Database.WriteEventAndMembership(event, membership)

	return event.Id, err
}

// GetEventForUser gets an event, just ensure the user is a member of that event
func (e EventService) GetEventForUser(eventId string, userId string) (*Event, error) {
	return e.Database.ReadEventForUser(eventId, userId)
}

func (e EventService) UpdateEvent(eventId string, userId string, name string, emails []string) error {
	return e.Database.UpdateEvent(eventId, userId, name, emails)
}

func (e EventService) GetEventsForUser(userId string) ([]Event, error) {
	return e.Database.ReadEventsForUser(userId)
}

type Event struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	OwnerUserId string `json:"ownerUserId"`
}

//var allEvents = []Event{
//	{
//		Id:          "1",
//		Name:        "2020 Wright's Christmas",
//		OwnerUserId: "1",
//	},
//	{
//		Id:          "2",
//		Name:        "2021 Tapa's Thanksgiving",
//		OwnerUserId: "2",
//	},
//}

type Membership struct {
	Id      string
	EventId string
	UserId  string
}

//var memberships = []*Membership{
//	{
//		Id:      "3692dfe3-f17a-47e0-bf9f-0c85f55f9860",
//		EventId: "f677cbb0-5caf-4a8d-a74a-dd0b1660faca",
//		UserId: "e36afe1d-ef84-42a8-af38-0e6d60745e9e",
//	},
//	{
//		Id:      "41eb853b-26ab-418a-9641-689cbb17e2dc",
//		EventId: "f677cbb0-5caf-4a8d-a74a-dd0b1660faca",
//		UserId: "78f4b4d6-a5d2-403b-a389-b49e59e78385",
//	},
//	{
//		Id:      "b81677a2-ae47-402c-b551-a88e057f0be5",
//		EventId: "f677cbb0-5caf-4a8d-a74a-dd0b1660faca",
//		UserId: "2a0e83b0-48bb-4fa8-b93c-54c894b8b142",
//	},
//	{
//		Id:      "c4160f0b-71b6-43a1-9b0c-1468bee2f87a",
//		EventId: "f677cbb0-5caf-4a8d-a74a-dd0b1660faca",
//		UserId: "6ae71694-36a1-4fbd-bd73-f495738980cb",
//	},
//	{
//		Id:      "76aba7e7-fe25-4c3b-bc21-3bdb774979c9",
//		EventId: "863ed653-8e1f-48d3-bfbb-8b3942af7990",
//		UserId: "e36afe1d-ef84-42a8-af38-0e6d60745e9e",
//	},
//	{
//		Id:      "6ee109c5-b671-4588-a0b3-7a4dac77a1cf",
//		EventId: "863ed653-8e1f-48d3-bfbb-8b3942af7990",
//		UserId: "78f4b4d6-a5d2-403b-a389-b49e59e78385",
//	},
//}
