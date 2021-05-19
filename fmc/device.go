package fmc

import (
	"context"
	"fmt"
	"net/http"
)

type DevicesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type DeviceResponse struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetDeviceByName(ctx context.Context, name string) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
	}
	devices := &DevicesResponse{}
	err = v.DoRequest(req, devices, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
	}

	for _, device := range devices.Items {
		if device.Name == name {
			return &DeviceResponse{
				ID:   device.ID,
				Name: device.Name,
				Type: device.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no device found with name %s", name)
}
