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
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	// Items []struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
	Metadata    struct {
		Timestamp int `json:"timestamp"`
		LastUser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
		IPType     string `json:"ipType"`
		ParentType string `json:"parentType"`
	} `json:"metadata"`
	// } `json:"items"`
}

type NetworkObjectsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		Links struct {
			Self   string `json:"self"`
			Parent string `json:"parent"`
		} `json:"links"`
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
	Paging struct {
		Offset int      `json:"offset"`
		Limit  int      `json:"limit"`
		Count  int      `json:"count"`
		Next   []string `json:"next"`
		Pages  int      `json:"pages"`
	} `json:"paging"`
}

func (v *Client) GetNetworkObjectByNameOrValue(ctx context.Context, nameOrValue string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks?expanded=false&filter=nameOrValue:%s", v.domainBaseURL, nameOrValue)
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
	case l > 1:
		return nil, fmt.Errorf("duplicates found, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no network objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return v.GetNetworkObject(ctx, resp.Items[0].ID)
}

// /fmc_config/v1/domain/DomainUUID/object/networks?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) CreateNetworkObject(ctx context.Context, object *NetworkObject) (*NetworkObjectResponse, error) {
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

func (v *Client) GetNetworkObject(ctx context.Context, id string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/object/networks/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	item := &NetworkObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting network objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateNetworkObject(ctx context.Context, id string, object *NetworkObjectUpdateInput) (*NetworkObjectResponse, error) {
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

func (v *Client) DeleteNetworkObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/networks/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting network objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
