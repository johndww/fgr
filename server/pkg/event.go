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

type Membership struct {
	Id      string
	EventId string
	UserId  string
}
