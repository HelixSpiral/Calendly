package calendly

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// CalendlyWrapper holds the main Calendly client
type CalendlyWrapper struct {
	apiKey        string
	baseApiUrl    string
	customHeaders map[string]string
}

// CalendlyWrapperInput is used as input for the New function
type CalendlyWrapperInput struct {
	ApiKey        string
	BaseApiUrl    string
	CustomHeaders map[string]string
}

// New returns a CalendlyWrapper to be used
func New(input *CalendlyWrapperInput) *CalendlyWrapper {
	var baseApiUrl string

	if input.BaseApiUrl != "" {
		baseApiUrl = input.BaseApiUrl
	} else {
		baseApiUrl = "https://api.calendly.com/"
	}

	cw := &CalendlyWrapper{
		apiKey:        input.ApiKey,
		baseApiUrl:    baseApiUrl,
		customHeaders: input.CustomHeaders,
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

func (cw *CalendlyWrapper) sendPostReq(url string, payload []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
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
	for key, value := range cw.customHeaders {
		req.Header.Add(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}
