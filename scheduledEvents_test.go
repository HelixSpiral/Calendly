package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

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
