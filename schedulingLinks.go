package calendly

import (
	"encoding/json"
	"fmt"
)

// Note this part of the wrapper may change in the future, this endpoint doesn't appear fully finished from Calendly
// https://developer.calendly.com/api-docs/b3A6MzQyNTM0OQ-create-single-use-scheduling-link

// SchedulingLink holds a Calendly Scheduled Link
type SchedulingLink struct {
	BookingURL string `json:"booking_url"`
	Owner      string `json:"owner"`
	OwnerType  string `json:"owner_type"`
}

// SchedulingLinkInput is used as input to the CreateSingleUseSchedulingLink function
type SchedulingLinkInput struct {
	MaxEventCount int    `json:"max_event_count"`
	Owner         string `json:"owner"`
	OwnerType     string `json:"owner_type"`
}

// CreateSingleUseSchedulingLink creates a single use link for Calendly
func (cw *CalendlyWrapper) CreateSingleUseSchedulingLink(input *SchedulingLinkInput) (SchedulingLink, error) {
	var scheduledLink SchedulingLink
	var slr map[string]SchedulingLink

	url := fmt.Sprintf("%sscheduling_links", cw.baseApiUrl)

	payload, err := json.Marshal(input)
	if err != nil {
		return scheduledLink, err
	}

	resp, err := cw.sendPostReq(url, payload)
	if err != nil {
		return scheduledLink, err
	}

	err = json.Unmarshal(resp, &slr)
	if err != nil {
		return scheduledLink, err
	}

	scheduledLink = slr["resource"]

	return scheduledLink, nil
}
