package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkObjectUpdateInput struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ID          string `json:"id"`
}

type NetworkObject struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type NetworkObjectResponse struct {
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type NetworkObjectsResponse struct {
	Items []struct {
		Type  string `json:"type"`
		ID    string `json:"id"`
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"items"`
}

func (v *Client) GetFmcNetworkObjectByNameOrValue(ctx context.Context, nameOrValue string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks?expanded=true&limit=1000&filter=nameOrValue:%s", v.domainBaseURL, nameOrValue)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network object by name/value: %s - %s", url, err.Error())
	}
	resp := &NetworkObjectsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting network object by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcNetworkObject(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == nameOrValue || item.Value == nameOrValue {
				return v.GetFmcNetworkObject(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no network objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

// /fmc_config/v1/domain/DomainUUID/object/networks?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) CreateFmcNetworkObject(ctx context.Context, object *NetworkObject) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	item := &NetworkObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcNetworkObject(ctx context.Context, id string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	item := &NetworkObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcNetworkObject(ctx context.Context, id string, object *NetworkObjectUpdateInput) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	item := &NetworkObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcNetworkObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/networks/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting network objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
