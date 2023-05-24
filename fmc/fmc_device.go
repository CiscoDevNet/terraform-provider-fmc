package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type AccessPolicyItem struct {
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
		Name         string            `json:"name"`
		HostName     string            `json:"hostName"`
		NatID        string            `json:"natID,omitempty"`
		RegKey       string            `json:"regKey,omitempty"`
		AccessPolicy *AccessPolicyItem `json:"accessPolicy"`
		PerformanceTier string         `json:"performanceTier,omitempty"`
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
	HostName    string `json:"hostName"`
	RegKey       string           `json:"regKey,omitempty"`
	NatID        string           `json:"natID,omitempty"`
	PerformanceTier string         `json:"performanceTier,omitempty"`
	AccessPolicy AccessPolicyItem `json:"accessPolicy"`
}

type Device struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Name         string            `json:"name"`
	HostName     string            `json:"hostName"`
	NatID        string            `json:"natID,omitempty"`
	RegKey       string            `json:"regKey,omitempty"`
	AccessPolicy *AccessPolicyItem `json:"accessPolicy"`
	PerformanceTier string         `json:"performanceTier,omitempty"`
	LicenseCaps  []string          `json:"license_caps,omitempty"`
}

func (v *Client) GetFmcDeviceByName(ctx context.Context, name string) (*Device, error) {
	url := fmt.Sprintf("%s/devices/devicerecords", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	devices := &DevicesResponse{}
	err = v.DoRequest(req, devices, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
	}

	for _, device := range devices.Items {
		if device.Name == name {
			Log.debug(device, "response")
			return &Device{
				ID:     device.ID,
				Name:   device.Name,
				Type:   device.Type,
				NatID:  device.NatID,
				RegKey: device.RegKey,
				// AccessPolicy: device.AccessPolicy,
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
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("adding device3: %s - %s", url, err.Error())
	}
	time.Sleep(300 * time.Second)
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcDevice(ctx context.Context, id string) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device1: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting device2: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcDevice(ctx context.Context, id string, device *Device) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&device)
	if err != nil {
		return nil, fmt.Errorf("updating device1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating device2: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating device3: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcDevice(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting Device: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}
