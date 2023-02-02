package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SubInterface struct {
	Ifname        string                        `json:"ifname"`
	ID            string                        `json:"id"`
	Type          string                        `json:"type"`
	Mode          string                        `json:"mode"`
	SubInterfaceID          int                        `json:"subIntfId"`
	VlanID          int                        `json:"vlanId"`
	MTU 			int                        `json:"MTU"`
	Enabled       bool                          `json:"enabled"`
	Ipv4          *SubInterfaceIpv4         `json:"ipv4"`
	Security_Zone *SubInterfaceSecurityZone `json:"securityZone"`
	Description   string                        `json:"description"`
	Name          string                        `json:"name"`
}

type SubInterfaceIpv4 struct {
	Static SubInterfaceIpv4Static `json:"static"`
	Dhcp   SubInterfaceIpv4Dhcp   `json:"dhcp"`
}
type SubInterfaceSecurityZone struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type SubInterfaceIpv4Static struct {
	Address string `json:"address"`
	Netmask string `json:"netmask"`
}

type SubInterfaceIpv4Dhcp struct {
	EnableDefaultRouteDHCP string `json:"enableDefaultRouteDHCP"`
	DhcpRouteMetric        int `json:"dhcpRouteMetric"`
}

type SubInterfaceResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Ifname      string `json:"ifname"`
	Enabled     bool   `json:"enabled"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type SubInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID     string `json:"id"`
		Ifname string `json:"ifname"`
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

func (v *Client) GetFmcSubInterface(ctx context.Context, deviceID string, name string) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces", v.domainBaseURL, deviceID)
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
		if SubInterface.Name == name {
			return &SubInterfaceResponse{
				ID:   SubInterface.ID,
				Name: SubInterface.Name,
				Type: SubInterface.Type,
			}, nil
		}
	}

	return nil, fmt.Errorf("no SubInterface found with name %s", name)
}
func (v *Client) CreateFmcSubInterface(ctx context.Context, deviceID string,  object *SubInterface) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces", v.domainBaseURL, deviceID)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	item := &SubInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	return item, nil
}
func (v *Client) UpdateFmcSubInterface(ctx context.Context, deviceID string, id string, object *SubInterface) (*SubInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/subinterfaces/%s", v.domainBaseURL, deviceID, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	item := &SubInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcSubInterface() {

}

