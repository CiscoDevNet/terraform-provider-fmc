package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PortObjectUpdateInput struct {
	Name        string `json:"name"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ID          string `json:"id"`
}

type PortObject struct {
	Name        string `json:"name"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type PortObjectResponse struct {
	Type        string `json:"type"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type PortObjectsResponse struct {
	Items []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Port string `json:"port"`
		Name string `json:"name"`
	} `json:"items"`
}

func (v *Client) GetFmcPortObjectByNameOrPort(ctx context.Context, nameOrPort string) (*PortObjectResponse, error) {
	url := fmt.Sprintf("%s/object/protocolportobjects?expanded=false&filter=nameOrValue:%s", v.domainBaseURL, nameOrPort)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting port object by name/port: %s - %s", url, err.Error())
	}
	resp := &PortObjectsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting port object by name/port: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcPortObject(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == nameOrPort || item.Port == nameOrPort {
				return v.GetFmcPortObject(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no port objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

func (v *Client) CreateFmcPortObject(ctx context.Context, object *PortObject) (*PortObjectResponse, error) {
	url := fmt.Sprintf("%s/object/protocolportobjects", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating port objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating port objects: %s - %s", url, err.Error())
	}
	item := &PortObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcPortObject(ctx context.Context, id string) (*PortObjectResponse, error) {
	url := fmt.Sprintf("%s/object/protocolportobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	item := &PortObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcPortObject(ctx context.Context, id string, object *PortObjectUpdateInput) (*PortObjectResponse, error) {
	url := fmt.Sprintf("%s/object/protocolportobjects/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating port objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating port objects: %s - %s", url, err.Error())
	}
	item := &PortObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcPortObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/protocolportobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting port objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
