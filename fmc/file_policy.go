package fmc

import (
	"context"
	"fmt"
	"net/http"
)

type FilePoliciesResponse struct {
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

type FilePolicy struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetFilePolicyByName(ctx context.Context, name string) (*FilePolicy, error) {
	url := fmt.Sprintf("%s/policy/filepolicies", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting File policy by name: %s - %s", url, err.Error())
	}
	filePolicies := &FilePoliciesResponse{}
	err = v.DoRequest(req, filePolicies, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting File policy by name: %s - %s", url, err.Error())
	}

	for _, filePolicy := range filePolicies.Items {
		if filePolicy.Name == name {
			return &FilePolicy{
				ID:   filePolicy.ID,
				Name: filePolicy.Name,
				Type: filePolicy.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no File policy found with name %s", name)
}
