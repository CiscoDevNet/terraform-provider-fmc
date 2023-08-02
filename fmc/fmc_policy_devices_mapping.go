package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// type PolicyDevicesAssignmentUpdateInput struct {
// 	Name          string `json:"name"`
// 	Value         string `json:"value"`
// 	DNSResolution string `json:"dnsResolution"`
// 	Description   string `json:"description"`
// 	Type          string `json:"type"`
// 	ID            string `json:"id"`
// }

type PolicyDevicesAssignmentSubConfig struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name,omitempty"`
}

type PolicyDevicesAssignment struct {
	Name        string                             `json:"name,omitempty"`
	ID          string                             `json:"id,omitempty"`
	Type        string                             `json:"type"`
	Targets     []PolicyDevicesAssignmentSubConfig `json:"targets"`
	Policy      PolicyDevicesAssignmentSubConfig   `json:"policy"`
}

// /fmc_config/v1/domain/DomainUUID/assignment/policyassignments?bulk=true ( Bulk POST operation on device policy assignments. )

func (v *Client) CreateFmcPolicyDevicesAssignment(ctx context.Context, object *PolicyDevicesAssignment) (*PolicyDevicesAssignment, error) {
	url := fmt.Sprintf("%s/assignment/policyassignments", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating device policy assignments: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PolicyDevicesAssignment{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcPolicyDevicesAssignment(ctx context.Context, id string) (*PolicyDevicesAssignment, error) {
	url := fmt.Sprintf("%s/assignment/policyassignments/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PolicyDevicesAssignment{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcPolicyDevicesAssignment(ctx context.Context, id string, object *PolicyDevicesAssignment) (*PolicyDevicesAssignment, error) {
	url := fmt.Sprintf("%s/assignment/policyassignments/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating device policy assignments: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &PolicyDevicesAssignment{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting device policy assignments: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}
