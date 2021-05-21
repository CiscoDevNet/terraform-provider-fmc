package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type URLObjectUpdateInput struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

type URLObject struct {
	Type        string `json:"type"`
	Url         string `json:"url"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type URLObjectResponse struct {
	// Links struct {
	// 	Self string `json:"self"`
	// } `json:"links"`
	// Items []struct {
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
	Type        string `json:"type"`
	URL         string `json:"url"`
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
	} `json:"metadata"`
	// } `json:"items"`
}

func (v *Client) GetURLObjectByNameOrValue(ctx context.Context, nameOrValue string) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/object/urls?expanded=false&filter=nameOrValue:%s", v.domainBaseURL, nameOrValue)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting url object by name/value: %s - %s", url, err.Error())
	}
	resp := &URLObjectResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting url object by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l > 1:
		return nil, fmt.Errorf("duplicates found, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no url objects found, length of response is: %d, expected 1, please check your filter", l)
	}
	return v.GetURLObject(ctx, resp.Items[0].ID)
}



func (v *Client) CreateURLObject(ctx context.Context, object *URLObject) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/object/urls", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating url objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating url objects: %s - %s", url, err.Error())
	}
	item := &URLObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	return item, err
}

func (v *Client) GetURLObject(ctx context.Context, id string) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/object/urls/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting url objects: %s - %s", url, err.Error())
	}
	item := &URLObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	return item, err
}

func (v *Client) UpdateURLObject(ctx context.Context, id string, object *URLObjectUpdateInput) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/object/urls/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating url objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating url objects: %s - %s", url, err.Error())
	}
	item := &URLObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	return item, err
}

func (v *Client) DeleteURLObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/urls/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting url objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
