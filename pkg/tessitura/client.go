package tessitura

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Folger-Shakespeare-Library/durb/pkg/config"
)

// Batch request/response types for the Tessitura batch endpoint.

type BatchRequest struct {
	Requests []BatchRequestItem `json:"Requests"`
}

type BatchRequestItem struct {
	HttpMethod string `json:"HttpMethod"`
	Id         int    `json:"Id"`
	Uri        string `json:"Uri"`
}

type BatchResponseItem struct {
	RequestId      int             `json:"RequestId"`
	ResponseObject json.RawMessage `json:"ResponseObject"`
	StatusCode     int             `json:"StatusCode"`
	ErrorMessages  []interface{}   `json:"ErrorMessages"`
}

type BatchResponseEnvelope struct {
	Responses   []BatchResponseItem `json:"Responses"`
	BatchFailed bool                `json:"BatchFailed"`
}

// APIError represents a non-2xx response from the Tessitura API.
type APIError struct {
	StatusCode int
	Status     string
	Body       string
}

func (e *APIError) Error() string {
	switch e.StatusCode {
	case 401, 403:
		return fmt.Sprintf("authentication failed (HTTP %d): check your credentials with 'tess configure'", e.StatusCode)
	case 404:
		return fmt.Sprintf("not found (HTTP %d)", e.StatusCode)
	case 409:
		return fmt.Sprintf("conflict (HTTP %d): the resource was modified by another request", e.StatusCode)
	default:
		msg := fmt.Sprintf("API error (HTTP %d)", e.StatusCode)
		if e.Body != "" {
			msg += ": " + e.Body
		}
		return msg
	}
}

// Client is a Tessitura REST API client.
type Client struct {
	BaseURL    string
	AuthHeader string
	HTTP       *http.Client
}

// NewClient creates a Client from the given config.
func NewClient(cfg config.Config) *Client {
	creds := strings.Join([]string{cfg.Username, cfg.UserGroup, cfg.Location, cfg.Password}, ":")
	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte(creds))

	return &Client{
		BaseURL:    cfg.Hostname + "/api",
		AuthHeader: auth,
		HTTP: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get performs an authenticated GET request. The path should start with "/".
func (c *Client) Get(ctx context.Context, path string) ([]byte, error) {
	url := c.BaseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Authorization", c.AuthHeader)
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	return body, nil
}

// Post performs an authenticated POST request with a JSON body.
func (c *Client) Post(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	jsonBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal request body: %w", err)
	}

	url := c.BaseURL + path

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("unable to create request: %w", err)
	}

	req.Header.Set("Authorization", c.AuthHeader)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			Body:       string(body),
		}
	}

	return body, nil
}

// Batch sends multiple API requests in a single HTTP call using the
// Tessitura batch endpoint. Returns the batch response with individual
// results keyed by request ID.
func (c *Client) Batch(ctx context.Context, items []BatchRequestItem) (*BatchResponseEnvelope, error) {
	// Batch URIs need the full URL prefix
	for i := range items {
		items[i].Uri = c.BaseURL + items[i].Uri
	}

	req := BatchRequest{Requests: items}
	data, err := c.Post(ctx, "/Batch", req)
	if err != nil {
		return nil, fmt.Errorf("batch request failed: %w", err)
	}

	var batchResp BatchResponseEnvelope
	if err := json.Unmarshal(data, &batchResp); err != nil {
		return nil, fmt.Errorf("unable to parse batch response: %w", err)
	}

	return &batchResp, nil
}
