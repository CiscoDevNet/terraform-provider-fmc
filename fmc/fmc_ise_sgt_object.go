package fmc

import (
	"context"
	"fmt"
	"net/http"
)

type IseSGTObjectsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Type  string `json:"type"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type IseSGTObject struct {
	ID   string
	Type string
	Name string
}

func (v *Client) GetFmcIseSGTObjectsByName(ctx context.Context, name string) (*IseSGTObject, error) {
	url := fmt.Sprintf("%s/object/isesecuritygrouptags", v.domainBaseURL)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting ise security group tags object by name: %s - %s", url, err.Error())
	}
	iseSGTObjects := &IseSGTObjectsResponse{}
	err = v.DoRequest(req, iseSGTObjects, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting ise security group tags object by name: %s - %s", url, err.Error())
	}

	for _, iseSGTObject := range iseSGTObjects.Items {
		if iseSGTObject.Name == name {
			return &IseSGTObject{
				ID:   iseSGTObject.ID,
				Name: iseSGTObject.Name,
				Type: iseSGTObject.Type,
			}, nil
		}
	}
	return nil, fmt.Errorf("no ise security group tags object found with name %s", name)
}
