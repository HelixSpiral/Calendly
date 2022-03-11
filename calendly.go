package calendly

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// CalendlyWrapper holds the main Calendly client
type CalendlyWrapper struct {
	apiKey     string
	baseApiUrl string
}

// New returns a CalendlyWrapper to be used
func New(apiKey string) *CalendlyWrapper {
	cw := &CalendlyWrapper{
		apiKey:     apiKey,
		baseApiUrl: "https://api.calendly.com/",
	}

	return cw
}

func (cw *CalendlyWrapper) sendGetReq(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := cw.sendRawReq(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cw *CalendlyWrapper) sendRawReq(req *http.Request) ([]byte, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cw.apiKey))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}
