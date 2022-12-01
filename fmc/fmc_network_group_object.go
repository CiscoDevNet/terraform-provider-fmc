package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkGroupObjectUpdateInput struct {
	Name        string                       `json:"name"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Description string                       `json:"description"`
	Type        string                       `json:"type"`
	ID          string                       `json:"id"`
}

type NetworkGroupObjectObjects struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

type NetworkGroupObjectLiterals struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type NetworkGroupObject struct {
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Type        string                       `json:"type"`
}

type NetworkGroupObjectResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type        string                       `json:"type"`
	Literals    []NetworkGroupObjectLiterals `json:"literals"`
	Objects     []NetworkGroupObjectObjects  `json:"objects"`
	Overridable bool                         `json:"overridable"`
	Name        string                       `json:"name"`
	ID          string                       `json:"id"`
	Description string                       `json:"description"`
	Metadata    struct {
		Timestamp int `json:"timestamp"`
		Lastuser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
	} `json:"metadata"`
}

// /fmc_config/v1/domain/DomainUUID/object/networkgroups?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) CreateFmcNetworkGroupObject(ctx context.Context, object *NetworkGroupObject) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating network objects: %s - %s", url, err.Error())
	}
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcNetworkGroupObject(ctx context.Context, id string) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcNetworkGroupObject(ctx context.Context, id string, object *NetworkGroupObjectUpdateInput) (*NetworkGroupObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating network objects: %s - %s", url, err.Error())
	}
	item := &NetworkGroupObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcNetworkGroupObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/networkgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting network objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
