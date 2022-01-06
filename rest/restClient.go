package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// Post sends a post request to the URL with the body
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	if err != nil {
		return nil, err
	}
	request.Header = headers
	return Client.Do(request)
}

func Get(url string, params map[string]string) (*http.Response, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)
	q := request.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}

	// assign encoded query string to http request
	request.URL.RawQuery = q.Encode()
	if err != nil {
		return nil, err
	}
	return Client.Do(request)
}
