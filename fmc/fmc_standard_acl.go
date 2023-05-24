package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type StandardAclResponse struct {
	Type        string          `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	ID          string          `json:"id"`
	Entries     []ObjectEntries `json:"entries"`
}

type StandardAclResponses struct {
	Items []StandardAclResponse `json:"items"`
}

type StandardAcl struct {
	Name    string          `json:"name"`
	Type    string          `json:"type"`
	ID      string          `json:"id"`
	Entries []ObjectEntries `json:"entries"`
}

type ObjectEntries struct {
	Action   string  `json:"action,omitempty"`
	Networks Network `json:"networks,omitempty"`
}

type Network struct {
	Objects  []Object_nw  `json:"objects"`
	Literals []Literal_nw `json:"literals"`
}

type Literal_nw struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type Object_nw struct {
	ID string `json:"id,omitempty"`
}

func (v *Client) GetFmcStandardAclByName(ctx context.Context, name string) (*StandardAcl, error) {
	url := fmt.Sprintf("%s/object/standardaccesslists", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting standard ACL by name: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	acls := &StandardAclResponses{}
	err = v.DoRequest(req, acls, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting standard acl by name: %s - %s", url, err.Error())
	}

	for _, acl := range acls.Items {
		if acl.Name == name {
			Log.debug(acl, "response")
			Log.line()
			return &StandardAcl{
				ID:   acl.ID,
				Name: acl.Name,
				Type: acl.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no standard acl found with name %s", name)
}

func (v *Client) CreateFmcStandardAcl(ctx context.Context, object *StandardAcl) (*StandardAclResponse, error) {
	url := fmt.Sprintf("%s/object/standardaccesslists", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("Creating standard acl: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Creating standard acl: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &StandardAclResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("Creating standard acl: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}
func (v *Client) GetFmcStandardAcl(ctx context.Context, id string) (*StandardAclResponse, error) {
	url := fmt.Sprintf("%s/object/standardaccesslists/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting std acl1: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(req, "request")
	item := &StandardAclResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting std acl2: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcStandardAcl(ctx context.Context, id string, object *StandardAcl) (*StandardAclResponse, error) {
	url := fmt.Sprintf("%s/object/standardaccesslists/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating acl1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating acl2: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &StandardAclResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating acl 3: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}
func (v *Client) DeleteFmcStandardAcl(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/standardaccesslists/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting standard acl : %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}
