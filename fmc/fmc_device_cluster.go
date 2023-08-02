package fmc

import (
	"context"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
	"time"
)
type ClustersResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []struct {
		ID    string `json:"id"`
		Type  string `json:"type"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
		Name         string            `json:"name"`
		ControlDevice  ControlDevice     `json:"controlDevice"`
		CommonBootstrap CommonBootstrap `json:"commonBootstrap"`
		DataDevices     []ControlDevice    `json:"dataDevices"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type ClusterResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	ControlDevice  ControlDevice     `json:"controlDevice"`
	CommonBootstrap CommonBootstrap `json:"commonBootstrap"`
	DataDevices     []ControlDevice     `json:"dataDevices"`
}
type Cluster struct {
	ID              string            	`json:"id,omitempty"`
	Type            string            	`json:"type"`
	Name            string            	`json:"name,omitempty"`
	ControlDevice   ControlDevice       `json:"controlDevice,omitempty"`
	CommonBootstrap CommonBootstrap     `json:"commonBootstrap,omitempty"`
	DataDevices     []ControlDevice     `json:"dataDevices,omitempty"`
	Action          string 				`json:"action,omitempty"`
}
type DataDevices struct {
	DataDevices     []ControlDevice     `json:"dataDevices"`
}

type ClusterNodeBootstrap struct {
	CclIP      string `json:"cclIp,omitempty"`
	Priority   int  `json:"priority,omitempty"`
}

type DeviceDetails struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Name         string            `json:"name"`
}

type ControlDevice struct {
	ClusterNode    *ClusterNodeBootstrap `json:"clusterNodeBootstrap,omitempty"`
	DeviceDetail   *DeviceDetails   `json:"deviceDetails,omitempty"`
}

type CommonBootstrap struct {
	CclInterface   *DeviceDetails   `json:"cclInterface,omitempty"`
	CclNetwork	   string 			`json:"cclNetwork,omitempty"`
	VniNetwork	   string 			`json:"vniNetwork,omitempty"`
}

func (v *Client) GetFmcDeviceClusterByName(ctx context.Context, name string) (*Cluster, error) {
    url := fmt.Sprintf("%s/deviceclusters/ftddevicecluster?offset=0&limit=25&expanded=true", v.domainBaseURL)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("getting cluster by name: %s - %s", url, err.Error())
    }
    Log.debug(req, "request")
    clusters := &ClustersResponse{}
    err = v.DoRequest(req, clusters, http.StatusOK)
    if err != nil {
        return nil, fmt.Errorf("getting cluster by name: %s - %s", url, err.Error())
    }

    for _, cluster := range clusters.Items {
        if cluster.Name == name {
            Log.debug(cluster, "response")
            return &Cluster{
                ID:     cluster.ID,
                Name:   cluster.Name,
                Type:   cluster.Type,
            }, nil
        }
    }
    return nil, fmt.Errorf("no device found with name %s", name)
}

func (v *Client) CreateFmcDeviceCluster(ctx context.Context, cluster *Cluster) (*ClusterResponse, error) {
	url := fmt.Sprintf("%s/deviceclusters/ftddevicecluster", v.domainBaseURL)
	body, err := json.Marshal(&cluster)
	if err != nil {
		return nil, fmt.Errorf("creating cluster: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating cluster: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &ClusterResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("creating cluster3: %s - %s", url, err.Error())
	}
	time.Sleep(180 * time.Second)
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcDeviceCluster(ctx context.Context, id string) (*ClusterResponse, error) {
	url := fmt.Sprintf("%s/deviceclusters/ftddevicecluster/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting cluster: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(req, "request")
	item := &ClusterResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting cluster: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcDeviceCluster(ctx context.Context, id string, cluster *Cluster) (*ClusterResponse, error) {
	url := fmt.Sprintf("%s/deviceclusters/ftddevicecluster/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&cluster)
	if err != nil {
		return nil, fmt.Errorf("updating cluster: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating cluster: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &ClusterResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating cluster: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcDeviceCluster(ctx context.Context, m interface{},id string) error {
    url := fmt.Sprintf("%s/deviceclusters/ftddevicecluster/%s", v.domainBaseURL, id)
    req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
    if err != nil {
        return fmt.Errorf("deleting cluster: %s - %s", url, err.Error())
    }
    Log.debug(req, "request")
    err = v.DoRequest(req, nil, http.StatusOK)
    Log.line()
    return err
}