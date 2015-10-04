package population

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Client for the World Population API
type Client struct {
	// BaseURL is the base url for the World Population API.
	BaseURL string

	// User agent used for HTTP requests to the World Population API.
	UserAgent string

	// HTTP client used to communicate with the World Population API.
	httpClient *http.Client
}

// NewClient returns a new World Population API client.
// If no *http.Client were provided then http.DefaultClient is used.
func NewClient(httpClients ...*http.Client) *Client {
	var httpClient *http.Client

	if len(httpClients) > 0 && httpClients[0] != nil {
		httpClient = httpClients[0]
	} else {
		cloned := *http.DefaultClient
		httpClient = &cloned
	}

	c := &Client{
		BaseURL:    Env("POPULATION_BASE_URL", "http://api.population.io/1.0"),
		UserAgent:  Env("POPULATION_USER_AGENT", "population.go"),
		httpClient: httpClient,
	}

	return c
}

// NewRequest creates a new API request.
func (c *Client) NewRequest(resourcePath string) (*http.Request, error) {
	u, err := url.Parse(c.BaseURL + resourcePath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", c.UserAgent)

	return req, nil
}

// Do sends an API request and returns the API response.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Decode API response into the value pointed to by v.
func (c *Client) Decode(req *http.Request, v interface{}) error {
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if ct := resp.Header.Get("Content-Type"); !strings.Contains(ct, "application/json") {
		return fmt.Errorf("Content-Type %q", ct)
	}

	if resp.StatusCode >= 400 {
		doc := map[string]string{}

		if err := json.NewDecoder(resp.Body).Decode(&doc); err != nil {
			return err
		}

		return fmt.Errorf(doc["detail"])
	}

	return json.NewDecoder(resp.Body).Decode(v)
}

// Get creates a new request with the provided resource path
// then decodes the response into the value pointed to by v.
func (c *Client) Get(resourcePath string, v interface{}) error {
	req, err := c.NewRequest(resourcePath)
	if err != nil {
		return err
	}

	return c.Decode(req, v)
}

// Env returns a string from the ENV, or fallback variable
func Env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}

	return fallback
}
