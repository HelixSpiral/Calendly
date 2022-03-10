package calendly

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// User holds a Calendly User object
type User struct {
	URI                 string    `json:"uri"`
	Name                string    `json:"name"`
	Slug                string    `json:"slug"`
	Email               string    `json:"email"`
	SchedulingUrl       string    `json:"scheduling_url"`
	Timezone            string    `json:"timezone"`
	AvatarUrl           string    `json:"avatar_url"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	CurrentOrganization string    `json:"current_organization"`
}

type userApiResponse struct {
	Resource User `json:"resource"`
}

// GetCurrentUser returns the current user for the given API key
func (cw *CalendlyWrapper) GetCurrentUser() (User, error) {
	var user User
	var userResponse userApiResponse

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", cw.baseApiUrl, "users/me"), nil)
	if err != nil {
		return user, err
	}

	resp, err := cw.sendRawReq(req)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &userResponse)
	if err != nil {
		return user, err
	}

	user = userResponse.Resource

	return user, nil
}

// GetUser gets a user by ID
func (cw *CalendlyWrapper) GetUser(id string) (User, error) {
	var user User
	var userResponse userApiResponse

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s%s", cw.baseApiUrl, "users/", id), nil)
	if err != nil {
		return user, err
	}

	resp, err := cw.sendRawReq(req)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(resp, &userResponse)
	if err != nil {
		return user, err
	}

	user = userResponse.Resource

	return user, nil
}
