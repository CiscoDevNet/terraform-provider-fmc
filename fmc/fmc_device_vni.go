package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type VNISecurityZone struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type VNIsResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []VNIResponse `json:"items"`
}

type VNIRequest struct {
	ID                    string          `json:"id,omitempty"`
	Name                  string          `json:"name,omitempty"`
	Enabled               bool 			  `json:"enabled"`
	Ifname                string          `json:"ifname"`
	Description           string          `json:"description"`
	Priority              int             `json:"priority"`
	VniId                 int             `json:"vniId"`
	VtepID				  int             `json:"vtepID"`
	MulticastGroupAddress string          `json:"multicastGroupAddress,omitempty"`
	SegmentId             int             `json:"segmentId,omitempty"`
	EnableProxy           bool            `json:"enableProxy,omitempty"`
	SecurityZone          VNISecurityZone `json:"securityZone"`
	IPv4                  IPv4            `json:"ipv4,omitempty"`
	Type				  string 			`json:"type"`
}

type VNIResponse struct {
	Type                  string          `json:"type"`
	ID                    string          `json:"id"`
	Name                  string          `json:"name"`
	Enabled               bool 			  `json:"enabled"`
	Description           string          `json:"description"`
	SegmentId             int             `json:"segmentId"`
	VNIID                 int             `json:"vniId"`
	VtepID				  int             `json:"vtepID"`
	EnableProxy           bool            `json:"enableProxy,omitempty"`
	MulticastGroupAddress string          `json:"multicastGroupAddress,omitempty"`
	Priority              int             `json:"priority"`
	Ifname                string          `json:"ifname"`
	SecurityZone          VNISecurityZone `json:"securityZone"`
}

/*
*
This method return the VNI ID based on device-id and VNI Name
*/
func (v *Client) GetFmcVNI(ctx context.Context, deviceID string, name string) (*VNIResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces?expanded=true", v.domainBaseURL, deviceID)
	log.Printf("GET FMC VNIs based on deviceid=%s and VNI name=%s", deviceID, name)
	log.Printf("Get FMC VNIs URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for getting VNIs: %s - %s", url, err.Error())
	}

	vni := &VNIsResponse{}
	err = v.DoRequest(req, vni, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while getting VNIs: %s - %s", url, err.Error())
	}
	for _, VNI := range vni.Items {
		log.Printf("VNI.ID=%s  VNI.Name=%s  VNI.Type=%s", VNI.ID, VNI.Name, VNI.Type)

		if VNI.Name == name {
			return &VNIResponse{
				ID:   VNI.ID,
				Name: VNI.Name,
			}, nil
		}
	}
	log.Printf("no interface found with name %s", name)
	return nil, fmt.Errorf("no interface found with name %s", name)
}

/*
*
This method return the all VNI details based on VNI ID
*/
func (v *Client) GetFmcVNIDetails(ctx context.Context, deviceID string, vnid string) (*VNIResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces/%s", v.domainBaseURL, deviceID, vnid)

	log.Printf("GET FMC VNI details based on deviceid=%s and VNI ID=%s", deviceID, vnid)
	log.Printf("Get FMC VNI details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for getting VNI details: %s - %s", url, err.Error())
	}

	vni := &VNIResponse{}
	err = v.DoRequest(req, vni, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while getting VNI by id: %s - %s", url, err.Error())
	}

	return &VNIResponse{
		ID:                    vni.ID,
		Name:                  vni.Name,
		Ifname:                vni.Ifname,
		Type:                  vni.Type,
		Description:           vni.Description,
		EnableProxy:           vni.EnableProxy,
		MulticastGroupAddress: vni.MulticastGroupAddress,
		Priority:              vni.Priority,
		SecurityZone:          vni.SecurityZone,
		SegmentId:             vni.SegmentId,
		VNIID:                 vni.VNIID,
		VtepID:				   vni.VtepID,
	}, nil

}

/*
*
This method creates the VNI
*/

func (v *Client) CreateFmcVNI(ctx context.Context, deviceID string, object *VNIRequest) (*VNIResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces", v.domainBaseURL, deviceID)
	log.Printf("Create VNI URL=%s", url)
	body, err := json.Marshal(&object)
	log.Printf("Create VNI Request=%s", body)
	
	
	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("error while creating context for VNI: %s - %s", url, err.Error())
	}
	vni := &VNIResponse{}
	err = v.DoRequest(req, vni, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("error while cretaing VNI: %s - %s", url, err.Error())
	}
	log.Printf("VNI interface created, response=%s", vni.ID)

	return vni, nil
}

/*
*
This method update the VNI details
*/

func (v *Client) UpdateFmcVNI(ctx context.Context, deviceID string, id string, object *VNIRequest) (*VNIResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces/%s", v.domainBaseURL, deviceID, id)

	log.Printf("Update VNI URL=%s", url)

	body, err := json.Marshal(&object)

	log.Printf("Update VNI Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))

	if err != nil {
		return nil, fmt.Errorf("error while Updating context for VNI: %s - %s", url, err.Error())
	}

	vni := &VNIResponse{}
	err = v.DoRequest(req, vni, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Updating VNI: %s - %s", url, err.Error())
	}

	log.Printf("VNI interface Update, response=%s", vni.ID)

	return vni, nil
}

/*
*
This method delete the VNI based in VNI ID
*/
func (v *Client) DeleteFmcVNIDetails(ctx context.Context, deviceID string, vnid string) (*VNIResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vniinterfaces/%s", v.domainBaseURL, deviceID, vnid)

	log.Printf("Delete FMC VNI details based on deviceid=%s and VNI ID=%s", deviceID, vnid)
	log.Printf("Delete FMC VNI details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for Delete VNI details: %s - %s", url, err.Error())
	}

	vni := &VNIResponse{}
	err = v.DoRequest(req, vni, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Deleting VNI by id: %s - %s", url, err.Error())
	}

	log.Printf("VNI interface Delete, response=%s", vni.ID)

	return vni, nil
}