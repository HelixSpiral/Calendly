package calendly

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// WebhookPayload holds a Calendly Webhook Payload object
type WebhookPayload struct {
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	Event     string    `json:"event"`
	Payload   Payload   `json:"payload"`
}

type Tracking struct {
	UtmCampaign    interface{} `json:"utm_campaign"`
	UtmSource      interface{} `json:"utm_source"`
	UtmMedium      interface{} `json:"utm_medium"`
	UtmContent     interface{} `json:"utm_content"`
	UtmTerm        interface{} `json:"utm_term"`
	SalesforceUUID interface{} `json:"salesforce_uuid"`
}

type Payload struct {
	CancelURL           string                      `json:"cancel_url"`
	CreatedAt           time.Time                   `json:"created_at"`
	Email               string                      `json:"email"`
	Event               string                      `json:"event"`
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
}

type CustomQuestionsAndAnswers struct {
	Answer   string `json:"answer"`
	Position int    `json:"position"`
	Question string `json:"question"`
}

// WebhookSubscription holds a Calendly Webhook Subscription object
type WebhookSubscription struct {
	URI            string    `json:"uri"`
	CallbackUrl    string    `json:"callback_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	RetryStartedAt time.Time `json:"retry_started_at"`
	State          string    `json:"state"`
	Events         []string  `json:"events"`
	Scope          string    `json:"scope"`
	Organization   string    `json:"organization"`
	User           string    `json:"user"`
	Creator        string    `json:"creator"`
}

// ListWebhookSubscriptionsInput is used as input for the ListWebhookSubscriptions function
type ListWebhookSubscriptionsInput struct {
	Organization string
	Scope        string
	Count        int
	PageToken    string
	Sort         string
	User         string
}

// ListWebhookSubscriptions lists the webhook subscriptions for the provided user or organization
func (cw *CalendlyWrapper) ListWebhookSubscriptions(input *ListWebhookSubscriptionsInput) ([]WebhookSubscription, error) {
	var ws []WebhookSubscription
	var wsr map[string]json.RawMessage

	url := fmt.Sprintf("%s%s", cw.baseApiUrl, "webhook_subscriptions")

	if input.Organization == "" {
		return ws, fmt.Errorf("you must provide an organization")
	}

	if strings.ToLower(input.Scope) == "user" && input.User == "" {
		return ws, fmt.Errorf("with a user scope you must provide a user")
	}

	url += fmt.Sprintf("?organization=%s", input.Organization)

	if input.Count != 0 {
		url += fmt.Sprintf("&count=%d", input.Count)
	}

	if input.Scope != "" {
		url += fmt.Sprintf("&scope=%s", input.Scope)
	}

	if input.PageToken != "" {
		url += fmt.Sprintf("&page_token=%s", input.PageToken)
	}

	if input.User != "" {
		url += fmt.Sprintf("&user=%s", input.User)
	}

	resp, err := cw.sendGetReq(url)
	if err != nil {
		return ws, err
	}

	err = json.Unmarshal(resp, &wsr)
	if err != nil {
		return ws, err
	}

	err = json.Unmarshal(wsr["collection"], &ws)
	if err != nil {
		return ws, err
	}

	return ws, nil
}

// GetWebhookSubscription gets the webhook subscription by id
func (cw *CalendlyWrapper) GetWebhookSubscription(id string) (WebhookSubscription, error) {
	var ws WebhookSubscription
	var wsr map[string]WebhookSubscription

	url := fmt.Sprintf("%s%s%s", cw.baseApiUrl, "webhook_subscriptions/", id)

	resp, err := cw.sendGetReq(url)
	if err != nil {
		return ws, err
	}

	err = json.Unmarshal(resp, &wsr)
	if err != nil {
		return ws, err
	}

	ws = wsr["resource"]

	return ws, nil
}
