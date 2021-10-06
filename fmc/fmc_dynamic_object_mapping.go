package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DynamicObjectMapping struct {
	Mappings []string `json:"mappings"`
	DynamicObject DynamicObjectUpdated `json:"dynamicObject"`
}

type DynamicObjectMappingRequest struct {
	Add []DynamicObjectMapping `json:"add"`
	Remove []DynamicObjectMapping `json:"remove"`
}

type DynamicObjectMappingsResponse struct {
	Items []struct{
		Mapping string `json:"mapping"`
	} `json:"items"`
}

func (v *Client) CreateFmcDynamicObjectMapping(ctx context.Context, dynamicObjectId string, mappings []string) (string, error) {
	url := fmt.Sprintf("%s/object/dynamicobjectmappings", v.domainBaseURL)
	body, err := json.Marshal(& DynamicObjectMappingRequest{
		Add: []DynamicObjectMapping{
			{
				DynamicObject: DynamicObjectUpdated{
					ID: dynamicObjectId,
				},
				Mappings: mappings,
			},
		},
	})
	if err != nil {
		return "", fmt.Errorf("creating dynamic object: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("creating dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMapping{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return "", fmt.Errorf("creating dynamic object mapping: %s - %s", url, err.Error())
	}

	return generateDynamicObjectMappingId(dynamicObjectId, mappings), nil
}

func (v *Client) DeleteFmcDynamicObjectMapping(ctx context.Context, dynamicObjectId string, mappings []string) error {
	url := fmt.Sprintf("%s/object/dynamicobjectmappings", v.domainBaseURL)
	body, err := json.Marshal(& DynamicObjectMappingRequest{
		Remove: []DynamicObjectMapping{
			{
				DynamicObject: DynamicObjectUpdated{
					ID: dynamicObjectId,
				},
				Mappings: mappings,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("deleting dynamic object: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("deleting dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMapping{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return fmt.Errorf("deleting dynamic object mapping: %s - %s", url, err.Error())
	}

	return nil
}

func (v *Client) GetFmcDynamicObjectMapping(ctx context.Context, id string) (*DynamicObjectMapping, error) {
	parsedDynamicObjectMapping, err := parseDynamicObjectMappingId(id)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/object/dynamicobjects/%s/mappings", v.domainBaseURL, parsedDynamicObjectMapping.DynamicObject.ID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object mapping: %s - %s", url, err.Error())
	}
	resp := &DynamicObjectMappingsResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting dynamic object mapping: %s - %s", url, err.Error())
	}

	for _, item := range resp.Items {
		if item.Mapping == parsedDynamicObjectMapping.Mappings {
			return &DynamicObjectMapping{
				DynamicObject: parsedDynamicObjectMapping.DynamicObject,
				Mappings: item.Mapping,
			}, nil
		}
	}

	return &DynamicObjectMapping{
		DynamicObject: parsedDynamicObjectMapping.DynamicObject,
		Mappings: "",
	}, nil
}

func generateDynamicObjectMappingId(dynamicObjectId string, mappings string) string {
	return fmt.Sprintf("%s-%s", dynamicObjectId, mappings)
}

func parseDynamicObjectMappingId(id string) (*DynamicObjectMapping, error) {
	substrings := strings.Split(id, "-")
	if len(substrings) != 2  {
		return nil, fmt.Errorf("unable to parse dynamic object mapping id")
	}

	return &DynamicObjectMapping{
		DynamicObject: DynamicObjectUpdated{
			ID: substrings[0],
		},
		Mappings: substrings[1],
	}, nil
}