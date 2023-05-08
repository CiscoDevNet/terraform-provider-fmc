package fmc

import (
    "bytes"
    "context"
    "fmt"
    "net/http"
    "encoding/json"
)

type ObjectItem struct{
    ID   string `json:"id"`
    Type string `json:"type"`
    Name string `json:"name"`
}
type ObjectItems struct{
	Objects *ObjectItem `json:"object"`
}
type StaticIPv4sResponse struct {
    Links struct {
        Self string `json:"self"`
    } `json:"links"`
    Items []struct {
        ID    string `json:"id"`
        Type  string `json:"type"`
        Links struct {
            Self string `json:"self"`
        } `json:"links"`
        InterfaceName string `json:"interfaceName"`
        SelectedNetworks []ObjectItem `json:"selectedNetworks"`
		Gateway   ObjectItems `json:"gateway"`
		RouteTracking  *ObjectItem `json:"routeTracking"`
		MetricValue        int 			`json:"metricValue"`
    } `json:"items"`
    Paging struct {
        Offset int `json:"offset"`
        Limit  int `json:"limit"`
        Count  int `json:"count"`
        Pages  int `json:"pages"`
    } `json:"paging"`
}

type StaticIPv4Response struct {
    Links struct {
        Self   string `json:"self"`
        Parent string `json:"parent"`
    } `json:"links"`
    Type        string `json:"type"`
    ID          string `json:"id"`
    InterfaceName string `json:"interfaceName"`
    SelectedNetworks []ObjectItem `json:"selectedNetworks"`
	Gateway   ObjectItems `json:"gateway"`
	RouteTracking  ObjectItem `json:"routeTracking"`
	MetricValue        int 			`json:"metricValue"`
}

type StaticIpv4Route struct {
	ID                 string       `json:"id,omitempty"`
	InterfaceName 	   string       `json:"interfaceName"`
	Type               string       `json:"type"`
	Tunneled           bool         `json:"isTunneled"`
	SelectedNetworks   []ObjectItem `json:"selectedNetworks,omitempty"`
	Gateway            *ObjectItems   `json:"gateway,omitempty"`
	RouteTracking      *ObjectItem  `json:"routeTracking"`
	MetricValue        int 			`json:"metricValue"`
}

func (v *Client) GetStaticIPv4RouteByName(ctx context.Context, deviceID string, networkName string) (*StaticIPv4Response, error) {
    url := fmt.Sprintf("%s/devices/devicerecords/%s/routing/ipv4staticroutes?expanded=true", v.domainBaseURL, deviceID)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("getting static route: %s - %s", url, err.Error())
    }
    item := &StaticIPv4sResponse{}
    err = v.DoRequest(req, item, http.StatusOK)
    if err != nil {
        return nil, fmt.Errorf("getting static route: %s - %s", url, err.Error())
    }

    for _, j := range item.Items {
		for _, StaticIpv4Route := range j.SelectedNetworks {
			if StaticIpv4Route.Name == networkName {
				return &StaticIPv4Response{
					ID:   j.ID,
					InterfaceName: StaticIpv4Route.Name,
					Type: j.Type,
				}, nil
			}
		}
    }

    return nil, fmt.Errorf("no route found with network name %s", networkName)
}


func (v *Client) CreateFmcStaticIPv4Route(ctx context.Context, deviceID string, ipv4route *StaticIpv4Route) (*StaticIPv4Response, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/routing/ipv4staticroutes", v.domainBaseURL, deviceID)
	body, err := json.Marshal(&ipv4route)
	if err != nil {
		return nil, fmt.Errorf("Creating ipv4 route1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("Creating ipv4 route2: %s - %s", url, err.Error())
	}
	item := &StaticIPv4Response{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("Creating ipv4 route3: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcStaticIPv4Route(ctx context.Context, deviceID string, id string) (*StaticIPv4Response, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/routing/ipv4staticroutes/%s", v.domainBaseURL,deviceID, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting ipv4 route1: %s - %s - id: %s", url, err.Error(), id)
	}
	item := &StaticIPv4Response{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting ipv4 route2: %s - %s - id: %s", url, err.Error(),id)
	}
	return item, nil
}

func (v *Client) UpdateFmcStaticIPv4Response(ctx context.Context,deviceID string,  id string, ipv4route *StaticIpv4Route) (*StaticIPv4Response, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/routing/ipv4staticroutes/%s", v.domainBaseURL, deviceID, id)
	body, err := json.Marshal(&ipv4route)
	if err != nil {
		return nil, fmt.Errorf("updating ipv4 route1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating ipv4 route2: %s - %s", url, err.Error())
	}
	item := &StaticIPv4Response{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating ipv4 route3: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcStaticIPv4Response(ctx context.Context,deviceID string, id string) error {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/routing/ipv4staticroutes/%s", v.domainBaseURL, deviceID, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting routes: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
