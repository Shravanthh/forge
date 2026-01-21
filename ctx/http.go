package ctx

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// HTTP client for API calls
var httpClient = &http.Client{Timeout: 30 * time.Second}

// Fetch makes an HTTP GET request and returns the response body.
func (c *Context) Fetch(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// FetchJSON makes a GET request and unmarshals JSON response.
func (c *Context) FetchJSON(url string, result any) error {
	data, err := c.Fetch(url)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Post makes an HTTP POST request with JSON body.
func (c *Context) Post(url string, body any) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	resp, err := httpClient.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// PostJSON makes a POST request and unmarshals JSON response.
func (c *Context) PostJSON(url string, body, result any) error {
	data, err := c.Post(url, body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Request makes a custom HTTP request.
func (c *Context) Request(method, url string, body any, headers map[string]string) ([]byte, error) {
	var bodyReader io.Reader
	if body != nil {
		data, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(data)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

// API is a helper for making requests to a base URL.
type API struct {
	BaseURL string
	Headers map[string]string
}

// NewAPI creates an API client with a base URL.
func NewAPI(baseURL string) *API {
	return &API{BaseURL: baseURL, Headers: make(map[string]string)}
}

// SetHeader sets a default header for all requests.
func (a *API) SetHeader(key, value string) *API {
	a.Headers[key] = value
	return a
}

// SetAuth sets the Authorization header.
func (a *API) SetAuth(token string) *API {
	return a.SetHeader("Authorization", "Bearer "+token)
}

// Get makes a GET request.
func (a *API) Get(c *Context, path string, result any) error {
	data, err := c.Request("GET", a.BaseURL+path, nil, a.Headers)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// Post makes a POST request.
func (a *API) Post(c *Context, path string, body, result any) error {
	data, err := c.Request("POST", a.BaseURL+path, body, a.Headers)
	if err != nil {
		return err
	}
	if result != nil {
		return json.Unmarshal(data, result)
	}
	return nil
}

// Put makes a PUT request.
func (a *API) Put(c *Context, path string, body, result any) error {
	data, err := c.Request("PUT", a.BaseURL+path, body, a.Headers)
	if err != nil {
		return err
	}
	if result != nil {
		return json.Unmarshal(data, result)
	}
	return nil
}

// Delete makes a DELETE request.
func (a *API) Delete(c *Context, path string, result any) error {
	data, err := c.Request("DELETE", a.BaseURL+path, nil, a.Headers)
	if err != nil {
		return err
	}
	if result != nil {
		return json.Unmarshal(data, result)
	}
	return nil
}
