package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

func TestListUsersEventTypes(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	etl, err := cw.ListUsersEventTypes(&calendly.ListUsersEventTypesInput{
		Active: "true",
		Count:  10,
		User:   "https://api.calendly.com/users/AAAAAAAAAAAAAAAA", // Why we have to use a full url instead of just the uuid is beyond me :shrug:
	})
	if err != nil {
		t.Fatal("Failed to get event:", err)
	}

	if etl[0].Name != "15 Minute Meeting" {
		t.Fatalf("Expected '15 Minute Meeting', received '%s'", etl[0].Name)
	}

	etl, err = cw.ListUsersEventTypes(&calendly.ListUsersEventTypesInput{
		Active:       "false",
		Sort:         "name:asc",
		PageToken:    "sNjq4TvMDfUHEl7zHRR0k0E1PCEJWvdi",
		Organization: "https://api.calendly.com/organizations/AAAAAAAAAAAAAAAA",
	})
	if err != nil {
		t.Fatal("Failed to get event:", err)
	}

	if etl[0].Name != "15 Minute Meeting" {
		t.Fatalf("Expected '15 Minute Meeting', received '%s'", etl[0].Name)
	}
}

func TestGetEventType(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	event, err := cw.GetEventType("test-event")
	if err != nil {
		t.Fatal("Failed to get event:", err)
	}

	if event.Name != "15 Minute Meeting" {
		t.Fatalf("Expected '15 Minute Meeting', received '%s'", event.Name)
	}
}
