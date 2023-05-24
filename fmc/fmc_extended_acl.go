package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ExtendedAclResponse struct {
	Type    string         `json:"type"`
	Name    string         `json:"name"`
	ID      string         `json:"id"`
	Entries []Entries_data `json:"entries"`
}

type ExtendedAclResponses struct {
	Items []ExtendedAclResponse `json:"items"`
}

type ExtendedAcl struct {
	Name    string         `json:"name"`
	Type    string         `json:"type"`
	ID      string         `json:"id,omitempty"`
	Entries []Entries_data `json:"entries"`
}

type Entries_data struct {
	LogLevel            string     `json:"logLevel"`
	Action              string     `json:"action"`
	Logging             string     `json:"logging"`
	LogInterval         int        `json:"logInterval"`
	SourcePorts         Data_Ports `json:"sourcePorts"`
	SourceNetworks      Data_Nw    `json:"sourceNetworks"`
	DestinationNetworks Data_Nw    `json:"destinationNetworks"`
}

type Data_Ports struct {
	Objects  []Object_data        `json:"objects,omitempty"`
	Literals []Literals_Port_data `json:"literals,omitempty"`
}
type Data_Nw struct {
	Objects  []Object_data      `json:"objects,omitempty"`
	Literals []Literals_Nw_data `json:"literals,omitempty"`
}
type Object_data struct {
	ID string `json:"id"`
}

type Literals_Port_data struct {
	Type     string `json:"type"`
	Port     string `json:"port,omitempty"`
	Protocol string `json:"protocol,omitempty"`
}
type Literals_Nw_data struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func (v *Client) GetFmcExtendedAclByName(ctx context.Context, name string) (*ExtendedAcl, error) {
	url := fmt.Sprintf("%s/object/extendedaccesslists", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting extended ACL by name: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	acls := &ExtendedAclResponses{}
	err = v.DoRequest(req, acls, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting extended acl by name: %s - %s", url, err.Error())
	}

	for _, acl := range acls.Items {
		if acl.Name == name {
			Log.debug(acl, "response")
			Log.line()
			return &ExtendedAcl{
				ID:   acl.ID,
				Name: acl.Name,
				Type: acl.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no extended acl found with name %s", name)
}

func (v *Client) CreateFmcExtendedAcl(ctx context.Context, object *ExtendedAcl) (*ExtendedAclResponse, error) {
	url := fmt.Sprintf("%s/object/extendedaccesslists", v.domainBaseURL)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("Creating extended acl: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Creating extended acl: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &ExtendedAclResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("Creating extended acl: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcExtendedAcl(ctx context.Context, id string) (*ExtendedAclResponse, error) {
	url := fmt.Sprintf("%s/object/extendedaccesslists/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting extended acl: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(req, "request")
	item := &ExtendedAclResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting extended acl: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcExtendedAcl(ctx context.Context, id string, object *ExtendedAcl) (*ExtendedAclResponse, error) {
	url := fmt.Sprintf("%s/object/extendedaccesslists/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating extended accesslist: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating  extended accesslist : %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &ExtendedAclResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating  extended accesslist : %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcExtendedAcl(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/extendedaccesslists/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting extended acl : %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	err = v.DoRequest(req, nil, http.StatusOK)
	Log.line()
	return err
}
