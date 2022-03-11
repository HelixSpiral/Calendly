package calendly_test

import (
	"testing"

	"github.com/helixspiral/calendly"
)

func TestGetCurrentUser(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})

	currentUser, err := cw.GetCurrentUser()
	if err != nil {
		t.Fatal("Failed to get current user:", err)
	}

	if currentUser.Name != "John Doe" {
		t.Fatalf("Expected 'John Doe', received '%s'", currentUser.Name)
	}

	t.Log(currentUser)
}

func TestGettUser(t *testing.T) {
	cw := calendly.New(&calendly.CalendlyWrapperInput{
		BaseApiUrl: "https://stoplight.io/mocks/calendly/api-docs/395/",
		CustomHeaders: map[string]string{
			"Prefer": "code=200",
		},
	})
	user, err := cw.GetUser("AAAAAAAAAAAAAAAA")
	if err != nil {
		t.Fatal("Failed to get user:", err)
	}

	if user.Name != "John Doe" {
		t.Fatalf("Expected 'John Doe', received '%s'", user.Name)
	}
}
