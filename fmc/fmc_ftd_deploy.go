package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var deployment_type = "DeploymentRequest"

type FtdDeploy struct {
	Type          string   `json:"type"`
	Version       string   `json:"version"`
	Forcedeploy   bool     `json:"forceDeploy"`
	Ignorewarning bool     `json:"ignoreWarning"`
	Devicelist    []string `json:"deviceList"`
}

type DeployableDeviceResponse struct {
	Version string `json:"version"`
	Name    string `json:"name"`
	Device  struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"device"`
}

func (v *Client) GetFmcDeployableDevice(ctx context.Context, device_id string) (*DeployableDeviceResponse, error) {
	url := fmt.Sprintf("%s/deployment/deployabledevices?expanded=true", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting deployable devices: %s - %s", url, err.Error())
	}
	res := &struct {
		Items []DeployableDeviceResponse `json:"items"`
	}{}
	err = v.DoRequest(req, res, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting deployable devices: %s - %s", url, err.Error())
	}
	for _, item := range res.Items {
		if item.Device.ID == device_id {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("no devices found for deployment with ID %s", device_id)
}

func (v *Client) DeployToFTD(ctx context.Context, object FtdDeploy) error {
	url := fmt.Sprintf("%s/deployment/deploymentrequests", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return fmt.Errorf("deploying to device: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("deploying to device: %s - %s", url, err.Error())
	}
	item := &struct{}{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return fmt.Errorf("deploying to device: %s - %s", url, err.Error())
	}
	return nil
}
