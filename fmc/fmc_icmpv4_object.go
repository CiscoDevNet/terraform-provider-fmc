package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ICMPV4ObjectUpdateInput struct {
	Name        string `json:"name"`
	Icmptype    string `json:"icmpType"`
	Overridable bool   `json:"overridable"`
	Code        *int   `json:"code,omitempty"`
	Type        string `json:"type"`
	ID          string `json:"id"`
}

type ICMPV4Object struct {
	Name        string `json:"name"`
	Icmptype    string `json:"icmpType"`
	Overridable bool   `json:"overridable"`
	Code        *int   `json:"code,omitempty"`
	Type        string `json:"type"`
}

type ICMPV4ObjectResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Code        int    `json:"code"`
	Icmptype    string `json:"icmpType"`
	Overridable bool   `json:"overridable"`
	Description string `json:"description"`
	Name        string `json:"name"`
	ID          string `json:"id"`
}

// /fmc_config/v1/domain/DomainUUID/object/icmpv4objects?bulk=true ( Bulk POST operation on icmv4 objects. )

func (v *Client) CreateFmcICMPV4Object(ctx context.Context, object *ICMPV4Object) (*ICMPV4ObjectResponse, error) {
	url := fmt.Sprintf("%s/object/icmpv4objects", v.domainBaseURL)
	body, err := json.Marshal(&object)
	//panic(fmt.Sprintf("Body of request: %s", body))
	if err != nil {
		return nil, fmt.Errorf("creating icmv4 objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating icmv4 objects: %s - %s", url, err.Error())
	}
	item := &ICMPV4ObjectResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("getting icmv4 objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcICMPV4Object(ctx context.Context, id string) (*ICMPV4ObjectResponse, error) {
	url := fmt.Sprintf("%s/object/icmpv4objects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting icmv4 objects: %s - %s", url, err.Error())
	}
	item := &ICMPV4ObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting icmv4 objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcICMPV4Object(ctx context.Context, id string, object *ICMPV4ObjectUpdateInput) (*ICMPV4ObjectResponse, error) {
	url := fmt.Sprintf("%s/object/icmpv4objects/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating icmv4 objects: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating icmv4 objects: %s - %s", url, err.Error())
	}
	item := &ICMPV4ObjectResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting icmv4 objects: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcICMPV4Object(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/object/icmpv4objects/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting icmv4 objects: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
