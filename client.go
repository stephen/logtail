package logtail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	server             string
	authToken          string
	defaultContentType ContentType
	defaultMarshaler   Marshaler
	httpClient         *http.Client
}

var defaultServerUrl = "https://in.logtail.com"

func NewClient(authToken string, options ...Option) *Client {
	client := &Client{
		authToken:          authToken,
		server:             defaultServerUrl,
		httpClient:         http.DefaultClient,
		defaultContentType: Json,
		defaultMarshaler:   json.Marshal,
	}
	for _, option := range options {
		option.apply(client)
	}

	return client
}

func (c Client) Send(payload any) (n int, err error) {
	body, err := c.defaultMarshaler(payload)
	if err != nil {
		return 0, err
	}
	return c.Write(body)
}
func (c Client) Write(body []byte) (n int, err error) {
	request, _ := http.NewRequest(http.MethodPost, c.server, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", string(c.defaultContentType))
	request.Header.Set("Authorization", c.authToken)

	response, err := c.httpClient.Do(request)
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return 0, fmt.Errorf("log send: %w", err)
	}
	defer response.Body.Close()

	// Check response status, statusCode 202 means success
	if response.StatusCode != 202 {
		switch response.StatusCode {
		case 504:
			return 0, InvalidSourceToken
		case 406:
			return 0, InvalidBodyFormat
		default:
			return 0, fmt.Errorf("log send failed: %s", response.Status)
		}
	}

	return len(body), nil
}
