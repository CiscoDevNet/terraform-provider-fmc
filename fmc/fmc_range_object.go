package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RangeObjectUpdateInput struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
	ID          string `json:"id"`
}

type RangeObject struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

type RangeObjectResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Metadata    struct {
		Lastuser struct {
			Name string `json:"name"`
		} `json:"lastUser"`
		Domain struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"domain"`
		Iptype     string `json:"ipType"`
		Parenttype string `json:"parentType"`
	} `json:"metadata"`
}

// /fmc_config/v1/domain/DomainUUID/object/ranges?bulk=true ( Bulk POST operation on range objects. )

func (v *Client) CreateFmcRangeObject(ctx context.Context, object *RangeObject) (*RangeObjectResponse, error) {
	url := fmt.Sprintf("%s/object/ranges", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating range objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating range objects: %s - %s", url, err.Error())
	}
	item := &RangeObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting range objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcRangeObject(ctx context.Context, id string) (*RangeObjectResponse, error) {
	url := fmt.Sprintf("%s/object/ranges/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting range objects: %s - %s", url, err.Error())
	}
	item := &RangeObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting range objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcRangeObject(ctx context.Context, id string, object *RangeObjectUpdateInput) (*RangeObjectResponse, error) {
	url := fmt.Sprintf("%s/object/ranges/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating range objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating range objects: %s - %s", url, err.Error())
	}
	item := &RangeObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting range objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcRangeObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/ranges/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting range objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
