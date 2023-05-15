package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NatPolicy struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id,omitempty"`
}

type NatPolicyResponse struct {
	Type  string `json:"type"`
	Rules struct {
		Reftype string `json:"refType"`
		Type    string `json:"type"`
	} `json:"rules"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
}

type NatPoliciesResponse struct {
	Items []NatPolicyResponse `json:"items"`
}

func (v *Client) GetFmcNatPolicyByName(ctx context.Context, name string) (*NatPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies?expanded=false&filter=name:%s", v.domainBaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting nat policy by name/value: %s - %s", url, err.Error())
	}
	resp := &NatPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting nat policy by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcNatPolicy(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == name {
				return v.GetFmcNatPolicy(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no nat policies found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

// /fmc_config/v1/domain/DomainUUID/policy/ftdnatpolicies?bulk=true ( Bulk POST operation on nat policies. )

func (v *Client) CreateFmcNatPolicy(ctx context.Context, accessPolicy *NatPolicy) (*NatPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies", v.domainBaseURL)
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating nat policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating nat policies: %s - %s", url, err.Error())
	}
	item := &NatPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating nat policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcNatPolicy(ctx context.Context, id string) (*NatPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting nat policies: %s - %s", url, err.Error())
	}
	item := &NatPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting nat policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcNatPolicy(ctx context.Context, natId string, natPolicy *NatPolicy) (*NatPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s", v.domainBaseURL, natId)
	body, err := json.Marshal(&natPolicy)
	if err != nil {
		return nil, fmt.Errorf("update nat policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("update nat policies: %s - %s", url, err.Error())
	}
	item := &NatPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("update nat policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcNatPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting nat policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
