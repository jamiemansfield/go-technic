// Package solder provides a client for interacting with Technic
// Solder APIs.
//
// See https://docs.solder.io/
package solder

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

var (
	defaultBaseURL   = "https://solder.technicpack.net/api/"
	defaultUserAgent = "go-technic"
)

// A client manages communication with a Solder API, defaulting to
// Technic's instance.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// Base URL for API requests. Defaults to Technic's Solder
	// instance, but can be set to a domain endpoint to use with
	// other instances. BaseURL should always be set with a trailing
	// slash.
	BaseURL *url.URL

	// User Agent used when communicating with the Solder API.
	UserAgent string

	// Services used for accessing different parts of the Solder
	// API.
	Mod *ModService
	Modpack *ModpackService
}

type service struct {
	client *Client
}

// NewClient returns a new Solder API client. If a nil client is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{
		client: httpClient,
		BaseURL: baseURL,
		UserAgent: defaultUserAgent,
	}
	c.Mod = &ModService{client: c}
	c.Modpack = &ModpackService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided
// in urlStr, in which case it is resolved to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.
// If specified, the value pointed to by the body is JSON encoded and
// included as the request body.
func (c *Client) NewRequest(method string, urlStr string, body interface{}) (*http.Request, error) {
	// Resolve absolute URL
	u, err := c.BaseURL.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	// Encode body as JSON
	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	// Create the request
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	// Set request headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response
// is JSON decoded and stored in the value pointed to by v.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}

// GetApiInfo gets version information about the software running on the
// Solder instance.
func (c *Client) GetApiInfo() (*ApiInfo, error) {
	request, err := c.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		return nil, err
	}

	var response ApiInfo
	_, err = c.Do(request, &response)
	return &response, err
}
