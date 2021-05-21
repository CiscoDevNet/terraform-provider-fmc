package fmc

import (
	"context"
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

type SecurityZone struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetSecurityZoneByName(ctx context.Context, name string) (*SecurityZone, error) {
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
