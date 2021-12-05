package gohttp

import (
	"errors"
	"net/http"
)

func (c *httpClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
	//httpMethod := "GET"
	//url := "http://localhost:8080"
	client := http.Client{}

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, errors.New("unable to create new request")
	}

	fullHeaders := c.getRequestHeaders(headers)
	request.Header = fullHeaders

	//request.Header.Set("Accept", "application/json")
	return client.Do(request)

}

func (c *httpClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)
	// Add custom headers to the request
	for header, value := range c.Headers {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}
	return result
}