package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type FQDNObjectUpdateInput struct {
	Name          string `json:"name"`
	Value         string `json:"value"`
	DNSResolution string `json:"dnsResolution"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	ID            string `json:"id"`
}

type FQDNObject struct {
	Name          string `json:"name"`
	Value         string `json:"value"`
	DNSResolution string `json:"dnsResolution"`
	Description   string `json:"description"`
	Type          string `json:"type"`
}

type FQDNObjectResponse struct {
	Name          string `json:"name"`
	Value         string `json:"value"`
	DNSResolution string `json:"dnsResolution"`
	Description   string `json:"description"`
	Type          string `json:"type"`
	ID            string `json:"id"`
}

// /fmc_config/v1/domain/DomainUUID/object/fqdns?bulk=true ( Bulk POST operation on fqdn objects. )

func (v *Client) CreateFmcFQDNObject(ctx context.Context, object *FQDNObject) (*FQDNObjectResponse, error) {
	url := fmt.Sprintf("%s/object/fqdns", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating fqdn objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating fqdn objects: %s - %s", url, err.Error())
	}
	item := &FQDNObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting fqdn objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcFQDNObject(ctx context.Context, id string) (*FQDNObjectResponse, error) {
	url := fmt.Sprintf("%s/object/fqdns/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting fqdn objects: %s - %s", url, err.Error())
	}
	item := &FQDNObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting fqdn objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcFQDNObject(ctx context.Context, id string, object *FQDNObjectUpdateInput) (*FQDNObjectResponse, error) {
	url := fmt.Sprintf("%s/object/fqdns/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating fqdn objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating fqdn objects: %s - %s", url, err.Error())
	}
	item := &FQDNObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting fqdn objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcFQDNObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/fqdns/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting fqdn objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
