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
