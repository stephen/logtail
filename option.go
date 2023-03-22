package logtail

import (
	"fmt"
	"net/http"
)

type Option interface {
	apply(client *Client)
}

type serverOpt struct {
	server string
}

func (s serverOpt) apply(client *Client) {
	client.server = s.server
}

// WithCustomServer sets the custom server to connect .
func WithCustomServer(server string) Option {
	return serverOpt{server: server}
}

type authTokenOpt struct {
	authToken string
}

func (a authTokenOpt) apply(client *Client) {
	client.authToken = a.authToken
}

// WithAuthToken  sets the custom auth token.
func WithAuthToken(autoToken string) Option {
	return authTokenOpt{authToken: fmt.Sprintf("Bearer %s", autoToken)}
}

type httpTransportOpt struct {
	httpTransport *http.Transport
}

func (h httpTransportOpt) apply(client *Client) {
	if client.httpClient != nil {
		client.httpClient.Transport = h.httpTransport
	}
}

// WithHttpTransport sets the custom http transport.
func WithHttpTransport(transport *http.Transport) Option {
	return httpTransportOpt{httpTransport: transport}
}

type defaultContentTypeOpt struct {
	contentType ContentType
}

func (d defaultContentTypeOpt) apply(client *Client) {
	client.defaultContentType = d.contentType
}

// WithDefaultContentType sets the default content type to json.
func WithDefaultContentType(contentType ContentType) Option {
	return defaultContentTypeOpt{
		contentType: contentType,
	}
}

type marshalerOpt struct {
	marshaler Marshaler
}

func (m marshalerOpt) apply(client *Client) {
	client.defaultMarshaler = m.marshaler
}

// WithDefaultMarshaler sets the default marshaler.
func WithDefaultMarshaler(marshaler Marshaler) Option {
	return marshalerOpt{marshaler: marshaler}
}
