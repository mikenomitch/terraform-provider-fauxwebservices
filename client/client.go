package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HostURL - Default Hashicups URL
const HostURL string = "https://tfe-zone-cc09c2e7.ngrok.io"

// DefaultToken - Default Token to be removed
const DefaultToken string = "TODOFILLIN"

// Client -
type Client struct {
	BaseURL    *url.URL
	HostURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(host, token *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
		Token:   DefaultToken,
	}

	if host != nil {
		c.HostURL = *host
	}

	if token != nil {
		c.Token = *token
	}

	baseURL, err := url.Parse(c.HostURL)
	if err != nil {
		return nil, fmt.Errorf("bad url: %v", err)
	}

	c.BaseURL = baseURL

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

// NewRequest creates an API request. A relative URL path can be provided in
// path, in which case it is resolved relative to the apiVersionPath of the
// Client. Relative URL paths should always be specified without a preceding
// slash.
// If v is supplied, the value will be JSONAPI encoded and included as the
// request body. If the method is GET, the value will be parsed and added as
// query parameters.
func (c *Client) NewRequest(method, path string, body *strings.Reader) (*http.Request, error) {
	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	// Create a request specific headers map.
	reqHeaders := make(http.Header)
	reqHeaders.Set("Authorization", "Bearer "+c.Token)

	switch method {
	case "GET":
		reqHeaders.Set("Accept", "application/vnd.api+json")
	case "DELETE", "PATCH", "POST":
		reqHeaders.Set("Accept", "application/vnd.api+json")
		reqHeaders.Set("Content-Type", "application/vnd.api+json")
	case "PUT":
		reqHeaders.Set("Accept", "application/json")
		reqHeaders.Set("Content-Type", "application/octet-stream")
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	// Set the request specific headers.
	for k, v := range reqHeaders {
		req.Header[k] = v
	}

	return req, nil
}
