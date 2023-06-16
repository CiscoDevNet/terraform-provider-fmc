package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"strings"
)

type AccessPolicyItem struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type getDeviceID struct {
	Data struct {
		Devices struct {
			Items []struct {
				Name string `json:"name"`
				UID  string `json:"uid"`
			} `json:"items"`
		} `json:"devices"`
	} `json:"data"`
}
type FmcSpecific struct {
	Tags struct {
	} `json:"tags"`
	TagKeys             []any  `json:"tagKeys"`
	TagValues           []any  `json:"tagValues"`
	UID                 string `json:"uid"`
	Name                string `json:"name"`
	Namespace           string `json:"namespace"`
	Type                string `json:"type"`
	Version             int    `json:"version"`
}
type DevicesResponse struct {
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
		HostName     string            `json:"hostName"`
		NatID        string            `json:"natID,omitempty"`
		RegKey       string            `json:"regKey,omitempty"`
		AccessPolicy *AccessPolicyItem `json:"accessPolicy"`
		PerformanceTier string         `json:"performanceTier,omitempty"`
	} `json:"items"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
}

type DeviceResponse struct {
	Links struct {
		Self   string `json:"self"`
		Parent string `json:"parent"`
	} `json:"links"`
	Type        string `json:"type"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	HostName    string `json:"hostName"`
	RegKey       string           `json:"regKey,omitempty"`
	NatID        string           `json:"natID,omitempty"`
	PerformanceTier string         `json:"performanceTier,omitempty"`
	AccessPolicy AccessPolicyItem `json:"accessPolicy"`
}

type Device struct {
	ID           string            `json:"id"`
	Type         string            `json:"type"`
	Name         string            `json:"name"`
	HostName     string            `json:"hostName"`
	NatID        string            `json:"natID,omitempty"`
	RegKey       string            `json:"regKey,omitempty"`
	AccessPolicy *AccessPolicyItem `json:"accessPolicy"`
	PerformanceTier string         `json:"performanceTier,omitempty"`
	LicenseCaps  []string          `json:"license_caps,omitempty"`
}

func (v *Client) GetFmcDeviceByName(ctx context.Context, name string) (*Device, error) {
    url := fmt.Sprintf("%s/devices/devicerecords?filter=name:%s", v.domainBaseURL, name)
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
    }
    Log.debug(req, "request")
    devices := &DevicesResponse{}
    err = v.DoRequest(req, devices, http.StatusOK)
    if err != nil {
        return nil, fmt.Errorf("getting device by name: %s - %s", url, err.Error())
    }

    for _, device := range devices.Items {
        if device.Name == name {
            Log.debug(device, "response")
            return &Device{
                ID:     device.ID,
                Name:   device.Name,
                Type:   device.Type,
                NatID:  device.NatID,
                RegKey: device.RegKey,
            }, nil
        }
    }
    return nil, fmt.Errorf("no device found with name %s", name)
}

func (v *Client) CreateFmcDevice(ctx context.Context, device *Device) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords", v.domainBaseURL)
	body, err := json.Marshal(&device)
	if err != nil {
		return nil, fmt.Errorf("adding device1: %s - %s", url, err.Error())
	}
	var reqType string
	if v.is_cdfmc {
		reqType = "PUT"
	} else {
		reqType = "POST"
	}
	req, err := http.NewRequestWithContext(ctx, reqType, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("adding device2: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusAccepted)
	if err != nil {
		return nil, fmt.Errorf("adding device3: %s - %s", url, err.Error())
	}
	time.Sleep(300 * time.Second)
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) GetFmcDevice(ctx context.Context, id string) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting device1: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting device2: %s - %s - id: %s", url, err.Error(), id)
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) UpdateFmcDevice(ctx context.Context, id string, device *Device) (*DeviceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
	body, err := json.Marshal(&device)
	if err != nil {
		return nil, fmt.Errorf("updating device1: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating device2: %s - %s", url, err.Error())
	}
	Log.debug(req, "request")
	item := &DeviceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("Updating device3: %s - %s", url, err.Error())
	}
	Log.debug(item, "response")
	Log.line()
	return item, nil
}

func (v *Client) DeleteFmcCDODevice(ctx context.Context, name string,cdoHost string, cdoRegion string) error {
    // fetch ftd id
    url := fmt.Sprintf("https://edge.%s.cdo.cisco.com/api/public",cdoRegion)
    
    method := "POST"
    payload_temp := fmt.Sprintf("{\"query\":\"query {\\n        devices(name: \\\"%s\\\") {\\n            items {\\n                name\\n                 uid\\n            }\\n        }\\n}\",\"variables\":{}}", name)
    payload2 := strings.NewReader("{\"query\":\"query {\\n        devices(name: \\\"FMC\\\") {\\n            items {\\n                name\\n                 uid\\n            }\\n        }\\n}\",\"variables\":{}}")
    
    payload := strings.NewReader(payload_temp)

    req,err := http.NewRequestWithContext( ctx,method, url, payload)
    if err != nil {
        fmt.Println(err)
        return err
    }

    res := &getDeviceID{}
    err = v.DoRequest(req, res, http.StatusOK)

    if err != nil {
        fmt.Println(err)
        return err
    }

    device_uid := res.Data.Devices.Items[0].UID
    fmt.Println(device_uid)

    // fetch fmc id

    req2,err := http.NewRequestWithContext( ctx,method, url, payload2)
    if err != nil {
        fmt.Println(err)
        return err
    }

    res2 := &getDeviceID{}
    err = v.DoRequest(req2, res2, http.StatusOK)
    if err != nil {
        fmt.Println(err)
        return err
    }

    fmc_uid := res2.Data.Devices.Items[0].UID
    fmt.Println(fmc_uid)
    // fetch specific fmc id

    url2 := fmt.Sprintf("https://%s/aegis/rest/v1/device/%s/specific-device", cdoHost,fmc_uid)
    fmcReq, err := http.NewRequestWithContext(ctx,"GET", url2, nil)
    if err != nil {
        fmt.Println(err)
        return err
    }
    fmc_res := &FmcSpecific{}
    err = v.DoRequest(fmcReq, fmc_res, http.StatusOK)

    if err != nil {
        fmt.Println(err)
        return err
    }

    fmc_specific_id := fmc_res.UID
    fmt.Println(fmc_specific_id)
    
    // delete device

    url3 := fmt.Sprintf("https://%s/aegis/rest/v1/services/fmc/appliance/%s",  cdoHost,fmc_specific_id)
    payload_temp2 := fmt.Sprintf("{\"queueTriggerState\": \"PENDING_DELETE_FTDC\",\"stateMachineContext\": {\"ftdCDeviceIDs\": \"%s\"}}", device_uid)
    payload3 := strings.NewReader(payload_temp2)
    
    del, err := http.NewRequestWithContext(ctx, "PUT", url3, payload3)

    if err != nil {
        fmt.Println(err)
        return err
    }
    err = v.DoRequest(del, nil, http.StatusOK)
    if err != nil {
        fmt.Println(err)
        return err
    }
    return err
}

func (v *Client) DeleteFmcDevice(ctx context.Context, m interface{},id string, name string,cdoHost string, cdoRegion string) error {
    var err error
    if v.is_cdfmc {
        c := m.(*Client)
        err := c.DeleteFmcCDODevice(ctx, name, cdoHost, cdoRegion)
        if err != nil {
            return err
        }
    } else {
        url := fmt.Sprintf("%s/devices/devicerecords/%s", v.domainBaseURL, id)
        req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
        if err != nil {
            return fmt.Errorf("deleting Device: %s - %s", url, err.Error())
        }
        Log.debug(req, "request")
        err = v.DoRequest(req, nil, http.StatusOK)
        Log.line()
        return err
    }
    return err  
}