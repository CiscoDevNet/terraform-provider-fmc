package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SourceInterface struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type NveNeighborAddress struct {
	Literal struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"literal"`
	Object struct {
		Overridable bool   `json:"overridable"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Type        string `json:"type"`
	} `json:"object"`
}
type VTEPEntry struct {
	SourceInterface          SourceInterface    `json:"sourceInterface"`
	NveNeighborAddress       NveNeighborAddress `json:"nveNeighborAddress"`
	NveVtepId                int                `json:"nveVtepId"`
	NveDestinationPort       int                `json:"nveDestinationPort"`
	NveEncapsulationType     string             `json:"nveEncapsulationType"`
	NveNeighborDiscoveryType string             `json:"nveNeighborDiscoveryType"`
}
type VTEPPolicyResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID          string      `json:"id"`
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

type VTEPPoliciesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items  []VTEPPolicyResponse `json:"items"`
	Paging struct {
		OffsetMTU int `json:"offset"`
		Limit     int `json:"limit"`
		Count     int `json:"count"`
		Pages     int `json:"pages"`
	} `json:"paging"`
}

type VTEPPolicy struct {
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

// /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vteppolicies/{objectId}
func (v *Client) GetVTEPPoliciesByName(ctx context.Context, acpId, name string) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies/%s", v.domainBaseURL, acpId, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
	}
	vtepPoliciesResponse := &VTEPPolicyResponse{}
	err = v.DoRequest(req, vtepPoliciesResponse, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
	}
	return vtepPoliciesResponse, fmt.Errorf("no vteppolicies found with name %s", name)
}

// /api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/vteppolicies
func (v *Client) GetVTEPPolicies(ctx context.Context, acpId, name string) (*VTEPPoliciesResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
	}
	vtepPoliciesResponse := &VTEPPoliciesResponse{}
	err = v.DoRequest(req, vtepPoliciesResponse, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies by name: %s - %s", url, err.Error())
	}
	return vtepPoliciesResponse, fmt.Errorf("no vteppolicies found with name %s", name)
}

func (v *Client) CreateVTEPPolicies(ctx context.Context, acpId string, vtepPolicy *VTEPPolicy) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	body, err := json.Marshal(&vtepPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	item := &VTEPPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating vtep policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateVTEPPolicies(ctx context.Context, acpId string, vtepPolicy *VTEPPolicy) (*VTEPPolicyResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	body, err := json.Marshal(&vtepPolicy)
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s", url, err.Error())
	}
	item := &VTEPPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating vtep policies: %s - %s,%+v", url, err.Error(), vtepPolicy)
	}
	return item, nil
}

func (v *Client) DeleteVTEPPolicies(ctx context.Context, acpId string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, acpId)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting vtep policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
