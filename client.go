package judge0

import (
	"bytes"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	authProvider AuthProvider
	httpClient   *http.Client
}

func (c *Client) Request(
	method, path string,
	headers map[string]string,
	body []byte,
) (*http.Response, error) {
	fullURL, err := url.JoinPath(c.authProvider.GetBaseURL(), path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fullURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	c.authProvider.SetAuthHeaders(req)

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return c.httpClient.Do(req)
}

func NewClient(authProvider AuthProvider, timeout time.Duration) *Client {
	return &Client{
		authProvider: authProvider,
		httpClient:   &http.Client{Timeout: timeout},
	}
}
