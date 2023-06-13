package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type NetworkAnalysisPolicy struct {
	Type        string                               `json:"type"`
	Name        string                               `json:"name"`
	Description string                               `json:"description"`
	ID          string                               `json:"id,omitempty"`
	BasePolicy  NetworkAnalysisPolicyBasePolicyInput `json:"basePolicy"`
	SnortEngine string                               `json:"snortEngine"`
}

type NetworkAnalysisPolicyBasePolicyInput struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type NetworkAnalysisPolicyResponse struct {
	Type        string                               `json:"type"`
	Name        string                               `json:"name"`
	Description string                               `json:"description"`
	ID          string                               `json:"id"`
	BasePolicy  NetworkAnalysisPolicyBasePolicyInput `json:"basePolicy"`
	MetaData    struct {
		MappedPolicy struct {
			SnortEngine string `json:"snortEngine"`
		} `json:"mappedPolicy"`
	} `json:"metadata"`
}

type NetworkAnalysisPoliciesResponse struct {
	Items []struct {
		Name string `json:"name"`
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"items"`
}

func (v *Client) GetFmcNetworkAnalysisPolicyByName(ctx context.Context, name string) (*NetworkAnalysisPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/networkanalysispolicies?limit=1000", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting NWApolicy by name/value: %s - %s", url, err.Error())
	}
	resp := &NatPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting NWApolicy by name/value: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetFmcNetworkAnalysisPolicy(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == name {
				return v.GetFmcNetworkAnalysisPolicy(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no NWApolicies found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("no NWApolicies found, length of response is: %d, expected 1, please check your filter", len(resp.Items))
}

func (v *Client) GetFmcNetworkAnalysisPolicy(ctx context.Context, id string) (*NetworkAnalysisPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/networkanalysispolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting NWApolicy by id: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkAnalysisPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting NWApolicy by id: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) CreateFmcNetworkAnalysisPolicy(ctx context.Context, policy *NetworkAnalysisPolicy) (*NetworkAnalysisPolicyResponse, error) {
	url := fmt.Sprintf("%s/policy/networkanalysispolicies", v.domainBaseURL)
	body, err := json.Marshal(&policy)
	if err != nil {
		return nil, fmt.Errorf("creating nat policies: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating nat policies: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkAnalysisPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating NWApolicy: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcNetworkAnalysisPolicy(ctx context.Context, id string, policy *NetworkAnalysisPolicy) (*NetworkAnalysisPolicyResponse, error) {
	policy.Type = NWAtype

	url := fmt.Sprintf("%s/policy/networkanalysispolicies/%s", v.domainBaseURL, policy.ID)
	body, err := json.Marshal(&policy)
	if err != nil {
		return nil, fmt.Errorf("updating Network analysis policy: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating Network analysis policy: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &NetworkAnalysisPolicyResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating Network analysis policy: %s - %s,%+v", url, err.Error(), policy)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcNetworkAnalysisPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/networkanalysispolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting NWApolicies: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}
