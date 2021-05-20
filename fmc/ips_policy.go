package fmc

import (
	"context"
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

type IPSPolicy struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetIPSPolicyByName(ctx context.Context, name string) (*IPSPolicy, error) {
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
