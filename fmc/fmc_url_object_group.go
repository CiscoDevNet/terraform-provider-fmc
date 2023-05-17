package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type URLObjectGroupUpdateInput struct {
	Type        string                   `json:"type"`
	Objects     []URLObjectGroupObjects  `json:"objects"`
	Literals    []URLObjectGroupLiterals `json:"literals"`
	Description string                   `json:"description"`
	Name        string                   `json:"name"`
	ID          string                   `json:"id"`
}

type URLObjectGroupObjects struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

type URLObjectGroupLiterals struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type URLObjectGroup struct {
	Type        string      `json:"type"`
	Objects     interface{} `json:"objects"`
	Literals    interface{} `json:"literals"`
	Description string      `json:"description"`
	Name        string      `json:"name"`
}

type URLObjectGroupResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Type        string                   `json:"type"`
	Literals    []URLObjectGroupLiterals `json:"literals"`
	Objects     []URLObjectGroupObjects  `json:"objects"`
	Overridable bool                     `json:"overridable"`
	Description string                   `json:"description"`
	Name        string                   `json:"name"`
	ID          string                   `json:"id"`
	Metadata    struct {
		Timestamp int `json:"timestamp"`
		LastUser  struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
	} `json:"metadata"`
	// } `json:"items"`
}

// {
//     "name": "net1",
//     "value": "1.0.0.0/24",
//     "overridable": false,
//     "description": "Network obj 1",
//     "type": "Network"
// }
// POST /fmc_config/v1/domain/DomainUUID/object/networks

func (v *Client) CreateFmcURLObjectGroup(ctx context.Context, object *URLObjectGroup) (*URLObjectGroupResponse, error) {
	url := fmt.Sprintf("%s/object/urlgroups", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating url group objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating url group objects: %s - %s", url, err.Error())
	}
	item := &URLObjectGroupResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting url group objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcURLObjectGroup(ctx context.Context, id string) (*URLObjectGroupResponse, error) {
	url := fmt.Sprintf("%s/object/urlgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting url group objects: %s - %s", url, err.Error())
	}
	item := &URLObjectGroupResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting url group objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcURLObjectGroup(ctx context.Context, id string, object *URLObjectGroupUpdateInput) (*URLObjectGroupResponse, error) {
	url := fmt.Sprintf("%s/object/urlgroups/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating url group objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating url group objects: %s - %s", url, err.Error())
	}
	item := &URLObjectGroupResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting url group objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcURLObjectGroup(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/urlgroups/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting url group objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
