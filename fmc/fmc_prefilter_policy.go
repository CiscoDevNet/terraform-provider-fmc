package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

var prefilterPolicyType string = "PrefilterPolicy"

type PrefilterPolicyDefaultActionInput struct {
	//LogEnd          bool `json:"logEnd"`
	LogBegin        bool   `json:"logBegin"`
	SendEventsToFMC bool   `json:"sendEventsToFMC"`
	Action          string `json:"action"`
}

type PrefilterPolicyDefaultAction struct {
	//LogEnd          bool `json:"logEnd"`
	LogBegin        bool   `json:"logBegin"`
	SendEventsToFMC bool   `json:"sendEventsToFMC"`
	Action          string `json:"action"`
	ID              string `json:"id"`
}

type PrefilterPolicyInput struct {
	DefaultAction PrefilterPolicyDefaultActionInput `json:"defaultAction"`
	Name          string                            `json:"name"`
	Description   string                            `json:"description"`
	Type          string                            `json:"type"`
}

type PrefilterPolicy struct {
	DefaultAction PrefilterPolicyDefaultAction `json:"defaultAction"`
	Name          string                       `json:"name"`
	Description   string                       `json:"description"`
	Type          string                       `json:"type"`
	ID            string                       `json:"id"`
}

func (v *Client) CreateFmcPrefilterPolicy(ctx context.Context, policy *PrefilterPolicyInput) (*PrefilterPolicy, error) {
	policy.Type = prefilterPolicyType

	url := fmt.Sprintf("%s/policy/prefilterpolicies", v.domainBaseURL)
	body, err := json.Marshal(&policy)
	if err != nil {
		return nil, fmt.Errorf("creating prefilter policy: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating prefilter policy: %s - %s", url, err.Error())
	}
	item := &PrefilterPolicy{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating prefilter policy: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcPrefilterPolicy(ctx context.Context, id string) (*PrefilterPolicy, error) {
	url := fmt.Sprintf("%s/policy/prefilterpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting prefilter policy: %s - %s", url, err.Error())
	}
	item := &PrefilterPolicy{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting prefilter policy: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcPrefilterPolicy(ctx context.Context, policy *PrefilterPolicy) (*PrefilterPolicy, error) {
	policy.Type = prefilterPolicyType

	url := fmt.Sprintf("%s/policy/prefilterpolicies/%s", v.domainBaseURL, policy.ID)
	body, err := json.Marshal(&policy)
	if err != nil {
		return nil, fmt.Errorf("updating prefilter policy: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating prefilter policy: %s - %s", url, err.Error())
	}
	item := &PrefilterPolicy{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating prefilter policy: %s - %s,%+v", url, err.Error(), policy)
	}
	return item, nil
}

func (v *Client) DeleteFmcPrefilterPolicy(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/policy/prefilterpolicies/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting prefilter policy: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
