package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

func TestCreateSingleUseSchedulingLink(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=201",
		},
	})

	schedulingLink, err := cw.CreateSingleUseSchedulingLink(&calendly.SchedulingLinkInput{
		MaxEventCount: 1,
		Owner:         "https://api.calendly.com/event_types/012345678901234567890",
		OwnerType:     "EventType",
	})
	if err != nil {
		t.Fatal("Failed creating single use scheduling link:", err)
	}

	if schedulingLink.BookingURL != "https://calendly.com/d/abcd-brv8/15-minute-meeting" {
		t.Fatalf("Expected 'https://calendly.com/d/abcd-brv8/15-minute-meeting', received '%s'", schedulingLink.BookingURL)
	}
}
