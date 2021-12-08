package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessPolicyCategory struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type AccessPolicyCategoryResponse struct {
	Name string `json:"name"`
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (v *Client) GetFmcAccessPoliciesCategory(ctx context.Context, id, accessPolicyId string) (*AccessPolicyCategoryResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/categories/%s", v.domainBaseURL, accessPolicyId, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access policy category: %s - %s", url, err.Error())
	}
	resp := &AccessPolicyCategoryResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access policy category: %s - %s", url, err.Error())
	}

	return resp, nil
}

// /fmc_config/v1/domain/DomainUUID/object/networks?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) CreateFmcAccessPoliciesCategory(ctx context.Context, accessPolicyId string, object *AccessPolicyCategory) (*AccessPolicyCategoryResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/categories", v.domainBaseURL, accessPolicyId)

	object.Type = "Category"

	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating access policy category: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating access policy category: %s - %s", url, err.Error())
	}
	item := &AccessPolicyCategoryResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating access policy category: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcAccessPoliciesCategory(ctx context.Context, id, accessPolicyId string) error {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/categories/%s", v.domainBaseURL, accessPolicyId, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting access policy category: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
