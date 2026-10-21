package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	Timeout time.Duration
}

/**
 * constructor functions
 */

func NewHTTPClient(timeout *time.Duration) *HTTPClient {
	_timeout := time.Second * 10

	if timeout != nil {
		_timeout = *timeout
	}

	return &HTTPClient{
		Timeout: _timeout,
	}
}

/**
 * public functions
 */

func (c *HTTPClient) DoGetRequest(ctx context.Context, url string, body any, responseBody any) (*http.Response, error) {
	return c.DoRequest(ctx, http.MethodGet, url, body, responseBody)
}

func (c *HTTPClient) DoPostRequest(ctx context.Context, url string, body any, responseBody any) (*http.Response, error) {
	return c.DoRequest(ctx, http.MethodPost, url, body, responseBody)
}

func (c *HTTPClient) DoRequest(ctx context.Context, method string, url string, body any, responseBody any) (*http.Response, error) {
	var requestBody io.Reader

	_ctx, cancel := context.WithTimeout(ctx, c.Timeout)

	defer cancel()

	if body != nil {
		data, err := json.Marshal(body)

		if err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}

		requestBody = bytes.NewReader(data)
	}

	request, err := http.NewRequestWithContext(_ctx, method, url, requestBody)
	if err != nil {
		return nil, err
	}

	if requestBody != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return response, fmt.Errorf("unexpected status %d", response.StatusCode)
	}

	// attempt to get the raw response body
	if err = json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return response, fmt.Errorf("failed to decode response: %w", err)
	}

	return response, nil
}
