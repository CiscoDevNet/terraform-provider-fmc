package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecuritySecurityZonesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type SecurityZoneRequest struct {
	Type          string `json:"type"`
	Name          string `json:"name"`
	InterfaceMode string `json:"interfaceMode"`
}

type SecurityZone struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Name          string `json:"name"`
	InterfaceMode string `json:"interfaceMode"`
}

func (v *Client) GetFmcSecurityZoneByName(ctx context.Context, name string) (*SecurityZone, error) {
	url := fmt.Sprintf("%s/object/securityzones", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting security zone by name: %s - %s", url, err.Error())
	}
	securityZones := &SecuritySecurityZonesResponse{}
	err = v.DoRequest(req, securityZones, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting security zone by name: %s - %s", url, err.Error())
	}

	for _, securityZone := range securityZones.Items {
		if securityZone.Name == name {
			return &SecurityZone{
				ID:   securityZone.ID,
				Name: securityZone.Name,
				Type: securityZone.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no security zone found with name %s", name)
}

func (v *Client) GetFmcSecurityZone(ctx context.Context, id string) (*SecurityZone, error) {
	url := fmt.Sprintf("%s/object/securityzones/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting security zone by name: %s - %s", url, err.Error())
	}
	securityZone := &SecurityZone{}
	err = v.DoRequest(req, securityZone, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting security zone by name: %s - %s", url, err.Error())
	}

	return securityZone, nil
}

func (v *Client) CreateFmcSecurityZone(ctx context.Context, zone *SecurityZoneRequest) (*SecurityZone, error) {
	url := fmt.Sprintf("%s/object/securityzones", v.domainBaseURL)
	body, err := json.Marshal(&zone)
	if err != nil {
		return nil, fmt.Errorf("creating url objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating security zone: %s - %s", url, err.Error())
	}
	item := &SecurityZone{}
	err = v.DoRequest(req, item, http.StatusCreated)
	return item, err
	//return nil, nil
}

func (v *Client) DeleteFmcSecurityZone(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/securityzones/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting security zone: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}

func (v *Client) UpdateFmcSecurityZone(ctx context.Context, id string, object *SecurityZoneRequest) (*SecurityZone, error) {
	// security zone api is build in a way, where you are specifying interfaces
	// when creating or updating security zone, while from terraform it's easier when
	// you can create resource associations/attachments as separate resources.
	// In order to implement this for security zone, we need to get api object from device
	// first, update modified properties, and use this object in the PUT request

	// Get original object from device
	securityZoneOriginal, err := v.GetFmcSecurityZone(ctx, id)
	if err != nil {
		return nil, err
	}

	// update properties
	securityZoneOriginal.Name = object.Name

	// push changes to the device
	url := fmt.Sprintf("%s/object/securityzones/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&securityZoneOriginal)
	if err != nil {
		return nil, fmt.Errorf("updating security zone: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating security zone: %s - %s", url, err.Error())
	}
	item := &SecurityZone{}
	err = v.DoRequest(req, item, http.StatusOK)
	return item, err
}
