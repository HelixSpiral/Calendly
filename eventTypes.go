package calendly

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// EventType holds a Calendly EventType object
type EventType struct {
	URI              string           `json:"uri"`
	Name             string           `json:"name"`
	Active           bool             `json:"active"`
	Slug             string           `json:"slug"`
	SchedulingUrl    string           `json:"scheduling_url"`
	Duration         int              `json:"duration"`
	Kind             string           `json:"kind"`
	PoolingType      string           `json:"pooling_type"`
	Type             string           `json:"type"`
	Color            string           `json:"color"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
	InternalNote     string           `json:"internal_note"`
	DescriptionPlain string           `json:"description_plain"`
	DescriptionHtml  string           `json:"description_html"`
	Profile          Profile          `json:"profile"`
	Secret           bool             `json:"secret"`
	CustomQuestions  []CustomQuestion `json:"custom_questions"`
}

// Profile holds a Calendly Profile object
type Profile struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

// CustomQuestion holds a Calendly CustomQuestion object
type CustomQuestion struct {
	Name          string   `json:"name"`
	Type          string   `json:"type"`
	Position      int      `json:"position"`
	Enabled       bool     `json:"enabled"`
	Required      bool     `json:"required"`
	AnswerChoices []string `json:"answer_choices"`
	IncludeOther  bool     `json:"include_other"`
}

// ListUsersEventTypesInput is used as input for the ListUsersEventTypes function
type ListUsersEventTypesInput struct {
	Active       string // This should be a bool, but Go doesn't have a null value for bools so we use string
	Count        int
	Organization string
	PageToken    string
	Sort         string
	User         string
}

// ListUsersEventTypesResponse is the response from the Calendly API
type listUsersEventTypesResponse struct {
	Collection []EventType `json:"collection"`
	Pagination struct {
		Count    int    `json:"count"`
		NextPage string `json:"next_page"`
	} `json:"pagination"`
}

// ListUsersEventTypes list a users event types
// Either User or Organization are required
func (cw *CalendlyWrapper) ListUsersEventTypes(input *ListUsersEventTypesInput) ([]EventType, error) {
	var etl []EventType
	var etlResponse listUsersEventTypesResponse

	url := fmt.Sprintf("%s%s", cw.baseApiUrl, "event_types")

	// Only user or organization can be supplied, not both
	if input.User != "" {
		url += fmt.Sprintf("?user=%s", input.User)
	} else if input.Organization != "" {
		url += fmt.Sprintf("?organization=%s", input.Organization)
	}

	switch strings.ToLower(input.Active) {
	case "true":
		url += "&active=true"
	case "false":
		url += "&active=false"
	}

	if input.Count != 0 && (input.Count >= 1 && input.Count <= 100) {
		url += fmt.Sprintf("&count=%d", input.Count)
	}

	if input.PageToken != "" {
		url += fmt.Sprintf("&page_token=%s", input.PageToken)
	}

	if input.Sort != "" {
		url += fmt.Sprintf("&sort=%s", input.Sort)
	}

	resp, err := cw.sendGetReq(url)
	if err != nil {
		return etl, err
	}

	err = json.Unmarshal(resp, &etlResponse)
	if err != nil {
		return etl, err
	}

	etl = etlResponse.Collection

	return etl, nil
}

// GetEventType returns the event type for the given id
func (cw *CalendlyWrapper) GetEventType(id string) (EventType, error) {
	var et EventType
	var eventTypeResponse map[string]EventType

	resp, err := cw.sendGetReq(fmt.Sprintf("%s%s%s", cw.baseApiUrl, "event_types/", id))
	if err != nil {
		return et, err
	}

	err = json.Unmarshal(resp, &eventTypeResponse)
	if err != nil {
		return et, err
	}

	et = eventTypeResponse["resource"]

	return et, nil
}
