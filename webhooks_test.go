package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

func TestListWebhookSubscriptions(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	ws, err := cw.ListWebhookSubscriptions(&calendly.ListWebhookSubscriptionsInput{
		Organization: "https://api.calendly.com/organizations/AAAAAAAAAAAAAAAA",
		User:         "https://api.calendly.com/users/AAAAAAAAAAAAAAAA",
		Scope:        "user",
		Count:        1,
		PageToken:    "test",
	})
	if err != nil {
		t.Fatal("Failed listing webhook subscriptions:", err)
	}

	if ws[0].URI != "https://api.calendly.com/webhook_subscriptions/AAAAAAAAAAAAAAAA" {
		t.Fatalf("Expected 'https://api.calendly.com/webhook_subscriptions/AAAAAAAAAAAAAAAA', received: '%s'", ws[0].URI)
	}
}

func TestGetWebhookSubscriptions(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	ws, err := cw.GetWebhookSubscription("AAAAAAAAAAAAAAAA")
	if err != nil {
		t.Fatal("Failed getting webhook subscription:", err)
	}

	if ws.URI != "https://api.calendly.com/webhook_subscriptions/AAAAAAAAAAAAAAAA" {
		t.Fatalf("Expected 'https://api.calendly.com/webhook_subscriptions/AAAAAAAAAAAAAAAA', received: '%s'", ws.URI)
	}
}
