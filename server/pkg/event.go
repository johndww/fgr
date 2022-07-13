package pkg

import (
	"errors"
	"github.com/google/uuid"
)

func CreateEvent(name string, ownerUserId string) string {
	event := Event{
		Id:          uuid.New().String(),
		Name:        name,
		OwnerUserId: ownerUserId,
	}

	allEvents = append(allEvents, event)

	membership := Membership{
		Id:      uuid.New().String(),
		EventId: event.Id,
		UserId:  ownerUserId,
	}

	memberships = append(memberships, &membership)
	return event.Id
}

func GetEventForUser(eventId string, userId string) (*Event, error) {
	isMember := false
	for _, membership := range memberships {
		if membership.EventId == eventId && membership.UserId == userId {
			isMember = true
		}
	}

	if !isMember {
		return nil, errors.New("not a member of the event")
	}

	for _, event := range allEvents {
		if event.Id == eventId {
			return &event, nil
		}
	}
	return nil, nil
}

func UpdateEvent(eventId string, userId string, name string, emails []string) error {
	//TODO so much work, just wait for postgres? not worth the logic..
	return errors.New("not implemented yet")
}

func GetEventsForUser(userId string) []Event {
	var eventIds []string
	for _, membership := range memberships {
		if membership.UserId == userId {
			eventIds = append(eventIds, membership.EventId)
		}
	}

	var events []Event
	for _, event := range allEvents {
		for _, lookForEventId := range eventIds {
			if lookForEventId == event.Id {
				events = append(events, event)
			}
		}
	}
	return events
}

type Event struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	OwnerUserId string `json:"ownerUserId"`
}

var allEvents = []Event{
	{
		Id:          "1",
		Name:        "2020 Wright's Christmas",
		OwnerUserId: "1",
	},
	{
		Id:          "2",
		Name:        "2021 Tapa's Thanksgiving",
		OwnerUserId: "2",
	},
}

type Membership struct {
	Id      string
	EventId string
	UserId  string
}

var memberships = []*Membership{
	{
		Id:      "1",
		EventId: "1",
		UserId:  "1",
	},
	{
		Id:      "2",
		EventId: "1",
		UserId:  "2",
	},
	{
		Id:      "3",
		EventId: "1",
		UserId:  "3",
	},
	{
		Id:      "4",
		EventId: "1",
		UserId:  "4",
	},
	{
		Id:      "5",
		EventId: "2",
		UserId:  "1",
	},
	{
		Id:      "6",
		EventId: "2",
		UserId:  "2",
	},
}
