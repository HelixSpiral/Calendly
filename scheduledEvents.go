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
	EventGuests      []Guest            `json:"event_guests"`
	Cancellation     Cancellation       `json:"cancellation"`
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

// Guest holds a Calendly Guest object
type Guest struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Invitee holds a Calendly Invitee object
type Invitee struct {
	CancelURL           string                      `json:"cancel_url"`
	CreatedAt           time.Time                   `json:"created_at"`
	Email               string                      `json:"email"`
	Event               string                      `json:"event"`
	FirstName           string                      `json:"first_name"`
	LastName            string                      `json:"last_name"`
	Name                string                      `json:"name"`
	NewInvitee          interface{}                 `json:"new_invitee"`
	OldInvitee          interface{}                 `json:"old_invitee"`
	QuestionsAndAnswers []CustomQuestionsAndAnswers `json:"questions_and_answers"`
	RescheduleURL       string                      `json:"reschedule_url"`
	Rescheduled         bool                        `json:"rescheduled"`
	Status              string                      `json:"status"`
	TextReminderNumber  interface{}                 `json:"text_reminder_number"`
	Timezone            string                      `json:"timezone"`
	Tracking            Tracking                    `json:"tracking"`
	UpdatedAt           time.Time                   `json:"updated_at"`
	URI                 string                      `json:"uri"`
	Canceled            bool                        `json:"canceled"`
	Cancellation        Cancellation                `json:"cancellation"`
	Payment             Payment                     `json:"payment"`
	NoShow              string                      `json:"no_show"`
}

type Cancellation struct {
	CancelerType string `json:"canceler_type"`
	CanceledBy   string `json:"canceled_by"`
	Reason       string `json:"reason"`
}

type Payment struct {
	ExternalID string  `json:"external_id"`
	Provider   string  `json:"provider"`
	Amount     float64 `json:"amount"`
	Currency   string  `json:"currency"`
	Terms      string  `json:"terms"`
	Successful bool    `json:"successful"`
}

type ListEventInviteesInput struct {
	ID        string
	Count     int
	Email     string
	PageToken string
	Sort      string
	Status    string
}

type ListEventsInput struct {
	Count        int
	InviteeEmail string
	MaxStartTime string
	MinStartTime string
	Organization string
	PageToken    string
	Sort         string
	Status       string
	User         string
}

// ListEventInvitees lists all the invitees for a specific event
func (cw *CalendlyWrapper) ListEventInvitees(input *ListEventInviteesInput) ([]Invitee, error) {
	var invList []Invitee
	var invListResp map[string]json.RawMessage

	if input.ID == "" {
		return invList, fmt.Errorf("you must provide an ID")
	}

	url := fmt.Sprintf("%s%s%s%s", cw.baseApiUrl, "scheduled_events/", input.ID, "/invitees")

	if input.Count >= 1 && input.Count <= 100 {
		url += fmt.Sprintf("?count=%d", input.Count)
	} else {
		// Because the first param uses ? instead of & we always use count
		// Easier than trying to dynamically add ? to the first param
		url += "?count=20"
	}

	if input.Email != "" {
		url += fmt.Sprintf("&email=%s", input.Email)
	}

	if input.PageToken != "" {
		url += fmt.Sprintf("&page_token=%s", input.PageToken)
	}

	if input.Sort != "" {
		url += fmt.Sprintf("&sort=%s", input.Sort)
	}

	if input.Status != "" {
		url += fmt.Sprintf("&status=%s", input.Status)
	}

	resp, err := cw.sendGetReq(url)
	if err != nil {
		return invList, err
	}

	err = json.Unmarshal(resp, &invListResp)
	if err != nil {
		return invList, err
	}

	err = json.Unmarshal(invListResp["collection"], &invList)
	if err != nil {
		return invList, err
	}

	return invList, nil
}

// ListEvents returns an event list from the Calendly API
func (cw *CalendlyWrapper) ListEvents(input *ListEventsInput) ([]Event, error) {
	var events []Event
	var eventsResponse map[string]json.RawMessage

	url := fmt.Sprintf("%s/scheduled_events", cw.baseApiUrl)

	if input.Count >= 1 && input.Count <= 100 {
		url += fmt.Sprintf("?count=%d", input.Count)
	} else {
		// Because the first param uses ? instead of & we always use count
		// Easier than trying to dynamically add ? to the first param
		url += "?count=20"
	}

	if input.MaxStartTime != "" {
		url += fmt.Sprintf("&max_start_time=%s", input.MaxStartTime)
	}

	if input.MinStartTime != "" {
		url += fmt.Sprintf("&min_start_time=%s", input.MinStartTime)
	}

	if input.Organization != "" {
		url += fmt.Sprintf("&organization=%s", input.Organization)
	}

	if input.InviteeEmail != "" {
		url += fmt.Sprintf("&email=%s", input.InviteeEmail)
	}

	if input.PageToken != "" {
		url += fmt.Sprintf("&page_token=%s", input.PageToken)
	}

	if input.Sort != "" {
		url += fmt.Sprintf("&sort=%s", input.Sort)
	}

	if input.Status != "" {
		url += fmt.Sprintf("&status=%s", input.Status)
	}

	if input.User != "" {
		url += fmt.Sprintf("&user=%s", input.User)
	}

	resp, err := cw.sendGetReq(url)
	if err != nil {
		return events, err
	}

	err = json.Unmarshal(resp, &eventsResponse)
	if err != nil {
		return events, err
	}

	err = json.Unmarshal(eventsResponse["collection"], &events)
	if err != nil {
		return events, err
	}

	return events, nil
}

// GetEventInvitee returns an invitee for a specific event
func (cw *CalendlyWrapper) GetEventInvitee(eventid, inviteeid string) (Invitee, error) {
	var inv Invitee
	var invResp map[string]Invitee

	resp, err := cw.sendGetReq(fmt.Sprintf("%sscheduled_events/%s/invitees/%s", cw.baseApiUrl, eventid, inviteeid))
	if err != nil {
		return inv, err
	}

	err = json.Unmarshal(resp, &invResp)
	if err != nil {
		return inv, err
	}

	inv = invResp["resource"]

	return inv, nil
}

// GetEvent gets the specified Calendly event
func (cw *CalendlyWrapper) GetEvent(id string) (Event, error) {
	var event Event
	var eventResponse map[string]Event

	resp, err := cw.sendGetReq(fmt.Sprintf("%sscheduled_events/%s", cw.baseApiUrl, id))
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(resp, &eventResponse)
	if err != nil {
		return event, err
	}

	event = eventResponse["resource"]

	return event, nil
}
