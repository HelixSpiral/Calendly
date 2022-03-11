package calendly

import (
	"encoding/json"
	"fmt"
	"time"
)

// Event holds a Calendly Event object
type Event struct {
	URI              string             `json:"uri"`
	Name             string             `json:"name"`
	Status           string             `json:"status"`
	StartTime        time.Time          `json:"start_time"`
	EndTime          time.Time          `json:"end_time"`
	EventType        string             `json:"event_type"`
	Location         Location           `json:"location"`
	InviteesCounter  InviteesCounter    `json:"invitees_counter"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	EventMemberships []EventMemberships `json:"event_memberships"`
	EventGuests      []EventGuests      `json:"event_guests"`
}

type Location struct {
	Type     string `json:"type"`
	Location string `json:"location"`
}

type InviteesCounter struct {
	Total  int `json:"total"`
	Active int `json:"active"`
	Limit  int `json:"limit"`
}

type EventMemberships struct {
	User string `json:"user"`
}

type EventGuests struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// EventResponse is the Calendly Event response from their API
type EventResponse struct {
	Resource Event `json:"resource"`
}

// GetEvent gets the specified Calendly event
func (cw *CalendlyWrapper) GetEvent(id string) (Event, error) {
	var event Event
	var eventResponse EventResponse

	resp, err := cw.sendGetReq(fmt.Sprintf("%s%s%s", cw.baseApiUrl, "scheduled_events/", id))
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(resp, &eventResponse)
	if err != nil {
		return event, err
	}

	event = eventResponse.Resource

	return event, nil
}
