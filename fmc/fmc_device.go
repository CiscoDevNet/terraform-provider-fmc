package fmc

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"encoding/json"
)

type AccessPolicyItem struct{
	ID   string `json:"id"`
	Type string `json:"type"`
}

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
		HostName        string `json:"hostName"`
		AccessPolicy  *AccessPolicyItem `json:"accessPolicy"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type DeviceResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	HostName        string `json:"hostName"`
	// Metadata    struct {
	// 	Lastuser struct {
	// 		Name string `json:"name"`
	// 	} `json:"lastUser"`
	// 	Domain struct {
	// 		Name string `json:"name"`
	// 		ID   string `json:"id"`
	// 	} `json:"domain"`
	// } `json:"metadata"`
	RegKey string `json:"regKey"`
	NatID string	`json:"natID,omitempty"`
	AccessPolicy  AccessPolicyItem `json:"accessPolicy"`
}

type Device struct {
	ID   string	`json:"id"`
	Type string	`json:"type"`
	Name string	`json:"name"`
	HostName string `json:"hostName"`
	NatID string	`json:"natID,omitempty"`
	RegKey string	`json:"regKey"`
	AccessPolicy *AccessPolicyItem `json:"accessPolicy"`
	LicenseCaps  []string   `json:"license_caps,omitempty"` 
}

func (v *Client) GetFmcDeviceByName(ctx context.Context, name string) (*Device, error) {
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
			return &Device{
				ID:   device.ID,
				Name: device.Name,
				Type: device.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no device found with name %s", name)
}

func (v *Client) CreateFmcDevice(ctx context.Context, device *Device) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords", v.domainBaseURL)
	body, err := json.Marshal(&device)
	if err != nil {
		return nil, fmt.Errorf("adding device1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("adding device2: %s - %s", url, err.Error())
	}
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("adding device3: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcDevice(ctx context.Context, id string) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device: %s - %s", url, err.Error())
	}
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	//ADD Timer //Accepted
	if err != nil {
		return nil, fmt.Errorf("getting device: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcDevice(ctx context.Context, id string, device *Device) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&device)
	if err != nil {
		return nil, fmt.Errorf("updating device: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating device: %s - %s", url, err.Error())
	}
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("Updating device: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcDevice(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting Device: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusAccepted)
	return err
}
