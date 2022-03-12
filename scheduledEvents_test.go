package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

func TestListEventInvitees(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	invList, err := cw.ListEventInvitees(&calendly.ListEventInviteesInput{
		ID:        "AAAAAAAAAAAAAAAA",
		Count:     1,
		Email:     "test@example.com",
		PageToken: "test",
		Sort:      "created_at:asc",
		Status:    "active",
	})
	if err != nil {
		t.Fatal("Failed to get invitee list:", err)
	}

	if invList[0].Name != "John Doe" {
		t.Fatalf("Expected 'John Doe', received '%s'.", invList[0].Name)
	}
}

func TestListEvents(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	eventList, err := cw.ListEvents(&calendly.ListEventsInput{
		Count:        1,
		InviteeEmail: "user@example.com",
		MaxStartTime: "2019-08-24T14:15:22Z",
		MinStartTime: "2019-08-24T14:15:22Z",
		Organization: "https://api.calendly.com/organization/AAAAAAAAAAAAAAAA",
		PageToken:    "test",
		Sort:         "created_at:asc",
		Status:       "active",
		User:         "https://api.calendly.com/users/AAAAAAAAAAAAAAAA",
	})
	if err != nil {
		t.Fatal("Failed to get event list:", err)
	}

	if eventList[0].Name != "15 Minute Meeting" {
		t.Fatalf("Expected '15 Minute Meeting', received '%s'.", eventList[0].Name)
	}
}

func TestGetEventInvitee(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	inv, err := cw.GetEventInvitee("test-event", "test-invitee")
	if err != nil {
		t.Fatal("Failed to get event:", err)
	}

	if inv.Name != "John Doe" {
		t.Fatalf("Expected 'John Doe', received '%s'", inv.Name)
	}
}

func TestGetEvent(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	event, err := cw.GetEvent("test-event")
	if err != nil {
		t.Fatal("Failed to get event:", err)
	}

	if event.Name != "15 Minute Meeting" {
		t.Fatalf("Expected '15 Minute Meeting', received '%s'", event.Name)
	}
}
