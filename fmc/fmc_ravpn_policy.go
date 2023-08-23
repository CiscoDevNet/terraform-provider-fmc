package fmc

import (
	// "bytes"
	"context"
	// "encoding/json"
	"fmt"
	"net/http"
)

type RavpnsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Name            string            `json:"name"`
		HostName        string            `json:"hostName"`
		NatID           string            `json:"natID,omitempty"`
		RegKey          string            `json:"regKey,omitempty"`
		AccessPolicy    *AccessPolicyItem `json:"accessPolicy"`
		PerformanceTier string            `json:"performanceTier,omitempty"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type RavpnResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type            string           `json:"type"`
	Description     string           `json:"description"`
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	HostName        string           `json:"hostName"`
	RegKey          string           `json:"regKey,omitempty"`
	NatID           string           `json:"natID,omitempty"`
	PerformanceTier string           `json:"performanceTier,omitempty"`
	AccessPolicy    AccessPolicyItem `json:"accessPolicy"`
}

type Ravpn struct {
	ID              string            `json:"id"`
	Type            string            `json:"type"`
	Name            string            `json:"name"`
	HostName        string            `json:"hostName"`
	NatID           string            `json:"natID,omitempty"`
	RegKey          string            `json:"regKey,omitempty"`
	AccessPolicy    *AccessPolicyItem `json:"accessPolicy"`
	PerformanceTier string            `json:"performanceTier,omitempty"`
	LicenseCaps     []string          `json:"license_caps,omitempty"`
}

func (v *Client) GetFmcRavpnByName(ctx context.Context, name string) (*Ravpn, error) {
	url := fmt.Sprintf("%s/policy/ravpns?expanded=true", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting ravpn policy by name: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	ravpns := &RavpnsResponse{}
	err = v.DoRequest(req, ravpns, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting ravpn policy by name: %s - %s", url, err.Error())
	}

	for _, ravpn := range ravpns.Items {
		if ravpn.Name == name {
			Log.debug(ravpn, "response")
			return &Ravpn{
				ID:     ravpn.ID,
				Name:   ravpn.Name,
				Type:   ravpn.Type,
				NatID:  ravpn.NatID,
				RegKey: ravpn.RegKey,
			}, nil
		}
	}
	return nil, fmt.Errorf("no ravpn policy found with name %s", name)
}

// func (v *Client) GetFmcDevice(ctx context.Context, id string) (*DeviceResponse, error) {
// 	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting device1: %s - %s - id: %s", url, err.Error(), id)
// 	}
// 	Log.debug(req, "request")
// 	item := &DeviceResponse{}
// 	err = v.DoRequest(req, item, http.StatusOK)
// 	if err != nil {
// 		return nil, fmt.Errorf("getting device2: %s - %s - id: %s", url, err.Error(), id)
// 	}
// 	Log.debug(item, "response")
// 	Log.line()
// 	return item, nil
// }
