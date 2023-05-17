package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SecurityGroupTag struct {
	Type        string `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

type SecurityGroupTagUpdateInput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag,omitempty"`
}

type SecurityGroupTagResponse struct {
	Name        string `json:"name"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Description string `json:"description"`
	// Tag         string `json:"tag,omitempty"`
}

type SecurityGroupTagsResponse struct {
	Links struct {
		Self string `json:"self"`
	}
	Items []SecurityGroupTagResponse `json:"items"`
}

func (v *Client) GetFmcSGTObjectsByName(ctx context.Context, name string) (*SecurityGroupTagResponse, error) {
	url := fmt.Sprintf("%s/object/securitygrouptags?limit=1000", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting security group tags by name: %s - %s", url, err.Error())
	}
	securityGroupTags := &SecurityGroupTagsResponse{}
	err = v.DoRequest(req, securityGroupTags, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting security group tags by name: %s - %s", url, err.Error())
	}

	for _, securityGroupTag := range securityGroupTags.Items {
		if securityGroupTag.Name == name {
			return &SecurityGroupTagResponse{
				ID:   securityGroupTag.ID,
				Name: securityGroupTag.Name,
				Type: securityGroupTag.Type,
				// Tag:  securityGroupTag.Tag,
			}, nil
		}
	}
	return nil, fmt.Errorf("no security group tags found with name %s", name)
}

func (v *Client) GetFmcSGTObjects(ctx context.Context, id string) (*SecurityGroupTagResponse, error) {
	url := fmt.Sprintf("%s/object/securitygrouptags/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting security group tags by id: %s - %s", url, err.Error())
	}
	securityGroupTag := &SecurityGroupTagResponse{}
	err = v.DoRequest(req, securityGroupTag, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting security group tags by id: %s - %s", url, err.Error())
	}

	return securityGroupTag, nil
}

func (v *Client) CreateFmcSGTObjects(ctx context.Context, object *SecurityGroupTag) (*SecurityGroupTagResponse, error) {
	url := fmt.Sprintf("%s/object/securitygrouptags", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating security group tags: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating security group tags: %s - %s", url, err.Error())
	}
	item := &SecurityGroupTagResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting security group tags: %s - %s", url, err.Error())
	}

	return item, nil
}

func (v *Client) UpdateFmcSGTObjects(ctx context.Context, id string, object *SecurityGroupTagUpdateInput) (*SecurityGroupTagResponse, error) {
	url := fmt.Sprintf("%s/object/securitygrouptags/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating SGT objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating SGT objects: %s - %s", url, err.Error())
	}
	item := &SecurityGroupTagResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting SGT objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcSGTObjects(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/securitygrouptags/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting sgt objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
