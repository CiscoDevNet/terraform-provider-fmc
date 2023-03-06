package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessPolicySubConfig struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessPolicyDefaultAction struct {
	Intrusionpolicy *AccessPolicySubConfig `json:"intrusionPolicy,omitempty"`
	Syslogconfig    *AccessPolicySubConfig `json:"syslogConfig,omitempty"`
	Type            string                 `json:"type"`
	Logbegin        bool                   `json:"logBegin"`
	Logend          bool                   `json:"logEnd"`
	Sendeventstofmc bool                   `json:"sendEventsToFMC"`
	Action          string                 `json:"action"`
	ID              string                 `json:"id,omitempty"`
	// Variableset struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"variableSet"`
	// Snmpconfig struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"snmpConfig"`
}

type AccessPolicy struct {
	ID            string                    `json:"id,omitempty"`
	Type          string                    `json:"type"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	Defaultaction AccessPolicyDefaultAction `json:"defaultAction"`
}

type AccessPolicyResponse struct {
	Type  string `json:"type"`
	Rules struct {
		Reftype string `json:"refType"`
		Type    string `json:"type"`
		Links   struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"rules"`
	Name          string                    `json:"name"`
	Description   string                    `json:"description"`
	ID            string                    `json:"id"`
	Defaultaction AccessPolicyDefaultAction `json:"defaultAction"`
}

type AccessPoliciesResponse struct {
	Items []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"items"`
}

func (v *Client) GetFmcAccessPolicyByName(ctx context.Context, name string) (*AccessPolicyResponse, error) {
	var url string
	url := fmt.Sprintf("%s/policy/accesspolicies?name=%s&expanded=false", v.domainBaseURL, name)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access policy by name/value: %s - %s", url, err.Error())
	}
	resp := &AccessPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access policy by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcAccessPolicy(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == name {
				return v.GetFmcAccessPolicy(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no access policies found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("this should not be reachable, this is a bug")
}

// /fmc_config/v1/domain/DomainUUID/policy/accesspolicies?bulk=true ( Bulk POST operation on access policies. )

func (v *Client) CreateFmcAccessPolicy(ctx context.Context, accessPolicy *AccessPolicy) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies", v.domainBaseURL)
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	item := &AccessPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating access policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcAccessPolicy(ctx context.Context, id string) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access policies: %s - %s", url, err.Error())
	}
	item := &AccessPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access policies: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcAccessPolicy(ctx context.Context, acp_id string, accessPolicy *AccessPolicy) (*AccessPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s", v.domainBaseURL, acp_id)
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("updating access policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating access policies: %s - %s", url, err.Error())
	}
	item := &AccessPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating access policies: %s - %s,%+v", url, err.Error(), accessPolicy)
	}
	return item, nil
}

func (v *Client) DeleteFmcAccessPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting access policies: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
