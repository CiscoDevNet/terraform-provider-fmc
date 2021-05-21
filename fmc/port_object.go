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
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
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

type PortObjectsResponse struct {
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

func (v *Client) GetPortObjectByNameOrPort(ctx context.Context, nameOrPort string) (*PortObjectResponse, error) {
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
	case l > 1:
		return nil, fmt.Errorf("duplicates found, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no port objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return v.GetPortObject(ctx, resp.Items[0].ID)
}

func (v *Client) CreatePortObject(ctx context.Context, object *PortObject) (*PortObjectResponse, error) {
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

func (v *Client) GetPortObject(ctx context.Context, id string) (*PortObjectResponse, error) {
	url := fmt.Sprintf("%s/object/protocolportobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	item := &PortObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting port objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdatePortObject(ctx context.Context, id string, object *PortObjectUpdateInput) (*PortObjectResponse, error) {
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

func (v *Client) DeletePortObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/protocolportobjects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting port objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
