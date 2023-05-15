package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type IPSPoliciesResponse struct {
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

type IPSPolicyResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Base_policy struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type IPSPolicy struct {
	ID             string      `json:"id"`
	Type           string      `json:"type"`
	Name           string      `json:"name"`
	InspectionMode string      `json:"inspectionMode"`
	BasePolicy     Base_policy `json:"basePolicy"`
}

func (v *Client) GetFmcIPSPolicyByName(ctx context.Context, name string) (*IPSPolicy, error) {
	url := fmt.Sprintf("%s/policy/intrusionpolicies", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting IPS policy by name: %s - %s", url, err.Error())
	}
	ipsPolicies := &IPSPoliciesResponse{}
	err = v.DoRequest(req, ipsPolicies, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting IPS policy by name: %s - %s", url, err.Error())
	}

	for _, ipsPolicy := range ipsPolicies.Items {
		if ipsPolicy.Name == name {
			return &IPSPolicy{
				ID:   ipsPolicy.ID,
				Name: ipsPolicy.Name,
				Type: ipsPolicy.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no IPS policy found with name %s", name)
}

func (v *Client) CreateFmcIPS(ctx context.Context, ipsPolicy *IPSPolicy) (*IPSPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/intrusionpolicies", v.domainBaseURL)
	body, err := json.Marshal(&ipsPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating IPS policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating IPS policies: %s - %s", url, err.Error())
	}
	item := &IPSPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating IPS policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcIPSPolicy(ctx context.Context, id string) (*IPSPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/intrusionpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting IPS policies: %s - %s", url, err.Error())
	}
	item := &IPSPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting IPS policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcIPSPolicy(ctx context.Context, id string, ipsPolicy *IPSPolicy) (*IPSPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/intrusionpolicies/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&ipsPolicy)
	if err != nil {
		return nil, fmt.Errorf("updating IPS policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating IPS policies: %s - %s", url, err.Error())
	}
	item := &IPSPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating IPS policies: %s - %s,%+v", url, err.Error(), ipsPolicy)
	}
	return item, nil
}

func (v *Client) DeleteFmcIPSPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/intrusionpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting IPS policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
