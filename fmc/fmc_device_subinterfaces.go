package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SubInterface struct {
	Ifname        string                        `json:"ifname,omitempty"`
	ID            string                        `json:"id,omitempty"`
	Type          string                        `json:"type"`
	Mode          string                        `json:"mode,omitempty"`
	SubInterfaceID          int                 `json:"subIntfId"`
	VlanID          int                         `json:"vlanId,omitempty"`
	MTU 			int                         `json:"MTU,omitempty"`
	Enabled       bool                          `json:"enabled,omitempty"`
	MgmntOnly     bool 						    `json:"managementOnly,omitempty"`
	Priority 	   int 							`json:"priority,omitempty"`
	IPv4          IPv4                          `json:"ipv4,omitempty"`
	IPv6         IPv6                          `json:"ipv6,omitempty"`
	SecurityZone PhysicalInterfaceSecurityZone  `json:"securityZone,omitempty"`
	Description   string                        `json:"description,omitempty"`
	Name          string                        `json:"name"`
}

type SubInterfaceResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Ifname      string `json:"ifname"`
	SubInterfaceID    int       `json:"subIntfId"`
	Enabled     bool   `json:"enabled"`
	MgmntOnly     bool 						    `json:"managementOnly,omitempty"`
	Priority 	   int 							`json:"priority,omitempty"`
	VlanID          int                         `json:"vlanId,omitempty"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	MTU 			int                        `json:"MTU,omitempty"`
	SecurityZone PhysicalInterfaceSecurityZone `json:"securityZone"`
	Mode          string                        `json:"mode"`
	IPv4          IPv4                          `json:"ipv4,omitempty"`
}

type SubInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID     string `json:"id"`
		Ifname string `json:"ifname"`
		SubInterfaceID    int       `json:"subIntfId"`
		MTU 			int                        `json:"MTU,omitempty"`
		SecurityZone PhysicalInterfaceSecurityZone `json:"securityZone"`
		IPv4          IPv4                          `json:"ipv4,omitempty"`
		VlanID          int                         `json:"vlanId,omitempty"`
		Type   string `json:"type"`
		Links  struct {
			Self string `json:"self"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

func (v *Client) GetFmcSubInterfaceByName(ctx context.Context, deviceID string, subinterface_id int) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces?expanded=true", v.domainBaseURL, deviceID)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting sub interfaces: %s - %s", url, err.Error())
	}
	subInterfaces := &SubInterfacesResponse{}
	err = v.DoRequest(req, subInterfaces, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting sub interfaces: %s - %s", url, err.Error())
	}

	for _, SubInterface := range subInterfaces.Items {
		if SubInterface.SubInterfaceID == subinterface_id {
			return &SubInterfaceResponse{
				ID:   SubInterface.ID,
				Name: SubInterface.Name,
				SubInterfaceID: SubInterface.SubInterfaceID,
				Type: SubInterface.Type,
				Ifname: SubInterface.Ifname,
				MTU: SubInterface.MTU,
				SecurityZone: SubInterface.SecurityZone,
				IPv4: SubInterface.IPv4,
				VlanID: SubInterface.VlanID,
			}, nil
		}
	}

	return nil, fmt.Errorf("no SubInterface found with id %d", subinterface_id)
}
func (v *Client) GetFmcSubInterface(ctx context.Context, deviceID string, id string) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting sub interfaces: %s - %s", url, err.Error())
	}
	subInterfaces := &SubInterfaceResponse{}
	err = v.DoRequest(req, subInterfaces, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting sub interfaces: %s - %s", url, err.Error())
	}

	if subInterfaces != nil {
		return &SubInterfaceResponse{
			ID:           subInterfaces.ID,
			Name:         subInterfaces.Name,
			Enabled:      subInterfaces.Enabled,
			Ifname:       subInterfaces.Ifname,
			Type:         subInterfaces.Type,
			MTU:          subInterfaces.MTU,
			Mode:         subInterfaces.Mode,
			SecurityZone: subInterfaces.SecurityZone,
		}, nil
	}

	return nil, fmt.Errorf("no Interface found with physicalInterfaceId %s", id)
}
func (v *Client) CreateFmcSubInterface(ctx context.Context, deviceID string,  object *SubInterface) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces", v.domainBaseURL, deviceID)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("creating sub interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating sub interfaces: %s - %s", url, err.Error())
	}
	item := &SubInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating sub interfaces: %s - %s", url, err.Error())
	}
	return item, nil
}
func (v *Client) UpdateFmcSubInterface(ctx context.Context, deviceID string, id string, object *SubInterface) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating subinterfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating subinterfaces: %s - %s", url, err.Error())
	}
	item := &SubInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting subinterfaces: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcSubInterface(ctx context.Context, deviceID string, id string) error{
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting subinterface: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)

	return err
}

