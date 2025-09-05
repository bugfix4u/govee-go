/*
 The MIT License

 Permission is hereby granted, free of charge, to any person obtaining a copy
 of this software and associated documentation files (the "Software"), to deal
 in the Software without restriction, including without limitation the rights
 to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 copies of the Software, and to permit persons to whom the Software is
 furnished to do so, subject to the following conditions:

 The above copyright notice and this permission notice shall be included in
 all copies or substantial portions of the Software.

 THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 THE SOFTWARE.
*/

// Package govee provides a client for Govee APIs
package govee

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func New(apiKey string) (*Client, error) {

	// Check the API Key
	if apiKey == "" {
		return nil, fmt.Errorf("missing Govee API Key")
	}

	// Create client instance
	c := &Client{
		ApiKey:     apiKey,
		Ctx:        context.Background(),
		HTTPClient: http.DefaultClient,
	}

	return c, nil
}

// NewFromEnv creates new Client instance from environment variables.
// Supported variables:
//   - GOVEE_API_KEY - Govee API Key(required)
func NewFromEnv() (*Client, error) {
	apiKey := os.Getenv("GOVEE_API_KEY")
	if apiKey == "" {
		log.Fatal("Missing GOVEE_API_KEY environment variable")
	}
	return New(apiKey)
}

// GetDevices retrieves the list of devices associated with the API key.
func (c *Client) GetDevices() ([]DeviceInfo, error) {
	goveeURL, err := url.JoinPath(goveeApiUrl, goveeDevices)
	if err != nil {
		return nil, fmt.Errorf("error creating URL for devices: %v", err)
	}
	respBody, err := c.doRequest(http.MethodGet, goveeURL, nil)
	if err != nil {
		return nil, err
	}
	devicesResponse := DiscoveryApiResponse{}
	err = json.Unmarshal(respBody, &devicesResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling devices response: %v", err)
	}

	return devicesResponse.Data, nil
}

// GetDeviceState retrieves the current state of the specified device.
// It requires a Device.
func (c *Client) GetDeviceState(device DeviceInfo) (*DeviceInfo, error) {
	goveeURL, err := url.JoinPath(goveeApiUrl, goveeDeviceState)
	if err != nil {
		return nil, fmt.Errorf("error creating URL for device state: %v", err)
	}
	requestPayload := RequestPayload{
		RequestId: "uuid",
		Payload: DeviceStatePayload{
			SKU:    device.SKU,
			Device: device.Device,
		},
	}
	respBody, err := c.doRequest(http.MethodPost, goveeURL, requestPayload)
	if err != nil {
		return nil, err
	}

	deviceState := DeviceStateApiResponse{}
	err = json.Unmarshal(respBody, &deviceState)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling devices response: %v", err)
	}

	return deviceState.Payload, nil
}

// doRequest issues an HTTP request to Govee API url according to parameters.
// Additionally, sets headers and User-Agent.
// It returns http.Response or error. Error can be a *hostError if host responded with error.
func (c *Client) doRequest(method, url string, body interface{}) ([]byte, error) {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Govee-API-Key", c.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgent)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := resolveHTTPError(resp); err != nil {
		return nil, err
	}

	return io.ReadAll(resp.Body)
}

// resolveHTTPError parses host error response and returns error with human-readable message
func resolveHTTPError(r *http.Response) error {

	// successful status code range
	if r.StatusCode >= 200 && r.StatusCode < 300 {
		return nil
	}

	// Unauthorized request
	if r.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("%s: a valid Govee API key is required", r.Status)
	}

	// Too many request
	if r.StatusCode == http.StatusTooManyRequests {
		return fmt.Errorf("%s: accounts are limited to %d requests a day", r.Status, requestLimit)
	}

	return fmt.Errorf("%s: %s", r.Status, http.StatusText(r.StatusCode))
}

// Close closes all idle connections.
func (c *Client) Close() {
	c.HTTPClient.CloseIdleConnections()
}
