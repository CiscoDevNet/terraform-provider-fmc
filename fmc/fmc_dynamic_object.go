package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DynamicObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ObjectType  string `json:"objectType"`
}

type DynamicObjectUpdated struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ObjectType  string `json:"objectType"`
	ID          string `json:"id"`
}

type DynamicObjectResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ObjectType  string `json:"objectType"`
	ID          string `json:"id"`
}

type DynamicObjectsResponse struct {
	Items []DynamicObjectResponse `json:"items"`
}

func (v *Client) GetFmcDynamicObjectByName(ctx context.Context, name string) (*DynamicObjectResponse, error) {
	url := fmt.Sprintf("%s/object/dynamicobjects?expanded=true&name=%s", v.domainBaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object by name: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object by name: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return &resp.Items[0], nil
	case l > 1:
		// seems that FMC does not allow that, but api returns list - so it's need to be handled somehow
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1", l)
	case l == 0:
		return nil, fmt.Errorf("no network objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

// /fmc_config/v1/domain/DomainUUID/object/networks?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) CreateFmcDynamicObject(ctx context.Context, object *DynamicObject) (*DynamicObjectResponse, error) {
	url := fmt.Sprintf("%s/object/dynamicobjects", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating dynamic object: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating dynamic objects: %s - %s", url, err.Error())
	}
	item := &DynamicObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating dynamic objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcDynamicObject(ctx context.Context, id string) (*DynamicObjectResponse, error) {
	url := fmt.Sprintf("%s/object/dynamicobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object: %s - %s", url, err.Error())
	}
	item := &DynamicObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcDynamicObject(ctx context.Context, id string, object *DynamicObjectUpdated) (*DynamicObjectResponse, error) {
	url := fmt.Sprintf("%s/object/dynamicobjects/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating dynamic objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating dynamic objects: %s - %s", url, err.Error())
	}
	item := &DynamicObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating dynamic objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcDynamicObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/dynamicobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting dynamic object: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
