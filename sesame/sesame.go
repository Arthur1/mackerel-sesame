package sesame

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{apiKey: apiKey}
}

type DeviceStatus struct {
	BatteryPercentage int64   `json:"batteryPercentage"`
	BatteryVoltage    float64 `json:"batteryVoltage"`
	Position          int64   `json:"position"`
	CHSesame2Status   string  `json:"CHSesame2Status"`
	Timestamp         int64   `json:"timestamp"`
	Wm2State          bool    `json:"wm2State"`
}

func (c *Client) GetDeviceStatus(deviceUUID string) (*DeviceStatus, error) {
	apiUrl, err := url.JoinPath("https://app.candyhouse.co/api/sesame2", deviceUUID)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-Api-Key", c.apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http error with status code %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	deviceStatus := &DeviceStatus{}
	if err := json.Unmarshal(resBody, deviceStatus); err != nil {
		return nil, err
	}

	return deviceStatus, nil
}
