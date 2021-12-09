package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type DynamicObjectMappingObject struct {
	ID string `json:"id"`
}

type DynamicObjectMapping struct {
	Mappings      []string                   `json:"mappings"`
	DynamicObject DynamicObjectMappingObject `json:"dynamicObject"`
}

type DynamicObjectMappingRequest struct {
	Add    []DynamicObjectMapping `json:"add"`
	Remove []DynamicObjectMapping `json:"remove"`
}

type DynamicObjectMappingsResponse struct {
	Items []struct {
		Mapping string `json:"mapping"`
	} `json:"items"`
}

func (v *Client) CreateFmcDynamicObjectMapping(ctx context.Context, dynamicObjectMapping *DynamicObjectMapping) error {
	url := fmt.Sprintf("%s/object/dynamicobjectmappings", v.domainBaseURL)
	body, err := json.Marshal(&DynamicObjectMappingRequest{
		Add: []DynamicObjectMapping{
			{
				DynamicObject: DynamicObjectMappingObject{
					ID: dynamicObjectMapping.DynamicObject.ID,
				},
				Mappings: dynamicObjectMapping.Mappings,
			},
		},
		Remove: []DynamicObjectMapping{},
	})
	if err != nil {
		return fmt.Errorf("creating dynamic object mapping: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("creating dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMapping{}
	err = v.DoRequest(req, resp, http.StatusCreated)
	if err != nil {
		return fmt.Errorf("creating dynamic object mapping: %s - %s", url, err.Error())
	}

	return nil
}

func (v *Client) DeleteFmcDynamicObjectMapping(ctx context.Context, dynamicObjectMapping *DynamicObjectMapping) error {
	url := fmt.Sprintf("%s/object/dynamicobjectmappings", v.domainBaseURL)
	body, err := json.Marshal(&DynamicObjectMappingRequest{
		Remove: []DynamicObjectMapping{
			{
				DynamicObject: DynamicObjectMappingObject{
					ID: dynamicObjectMapping.DynamicObject.ID,
				},
				Mappings: dynamicObjectMapping.Mappings,
			},
		},
		Add: []DynamicObjectMapping{},
	})
	if err != nil {
		return fmt.Errorf("deleting dynamic object mapping: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("deleting dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMapping{}
	err = v.DoRequest(req, resp, http.StatusCreated)
	if err != nil {
		return fmt.Errorf("deleting dynamic object mapping: %s - %s", url, err.Error())
	}

	return nil
}

func (v *Client) GetFmcDynamicObjectMapping(ctx context.Context, dynamicObjectMapping *DynamicObjectMapping) (*DynamicObjectMapping, error) {

	url := fmt.Sprintf("%s/object/dynamicobjects/%s/mappings", v.domainBaseURL, dynamicObjectMapping.DynamicObject.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMappingsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object mapping: %s - %s", url, err.Error())
	}

	if resp.Items == nil || len(resp.Items) < 1 {
		return &DynamicObjectMapping{
			Mappings: []string{},
			DynamicObject: DynamicObjectMappingObject{
				ID: dynamicObjectMapping.DynamicObject.ID,
			},
		}, nil
	}

	// convert array to a map in order to simplify search
	mappingsTable := make(map[string]string)
	for _, item := range resp.Items {
		mappingsTable[item.Mapping] = item.Mapping
	}

	existingMappings := []string{}
	for _, item := range dynamicObjectMapping.Mappings {
		if _, exists := mappingsTable[item]; exists {
			existingMappings = append(existingMappings, item)
		}
	}

	return &DynamicObjectMapping{
		Mappings: existingMappings,
		DynamicObject: DynamicObjectMappingObject{
			ID: dynamicObjectMapping.DynamicObject.ID,
		},
	}, nil
}
