package logtail

import "net/http"

type Client struct {
	authToken  string
	httpClient *http.Client
}
