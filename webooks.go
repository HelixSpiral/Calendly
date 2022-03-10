package calendly

import "time"

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
	CancelURL           string        `json:"cancel_url"`
	CreatedAt           time.Time     `json:"created_at"`
	Email               string        `json:"email"`
	Event               string        `json:"event"`
	Name                string        `json:"name"`
	NewInvitee          interface{}   `json:"new_invitee"`
	OldInvitee          interface{}   `json:"old_invitee"`
	QuestionsAndAnswers []interface{} `json:"questions_and_answers"`
	RescheduleURL       string        `json:"reschedule_url"`
	Rescheduled         bool          `json:"rescheduled"`
	Status              string        `json:"status"`
	TextReminderNumber  interface{}   `json:"text_reminder_number"`
	Timezone            string        `json:"timezone"`
	Tracking            Tracking      `json:"tracking"`
	UpdatedAt           time.Time     `json:"updated_at"`
	URI                 string        `json:"uri"`
	Canceled            bool          `json:"canceled"`
}
