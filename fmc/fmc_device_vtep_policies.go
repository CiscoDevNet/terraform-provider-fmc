package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type SourceInterface struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type Literal struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}
type Object struct {
	Overridable bool   `json:"overridable,omitempty"`
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
}
type NveNeighborAddress struct {
	Literal Literal `json:"literal,omitempty"`
	Object  Object  `json:"object,omitempty"`
}


type VTEPEntry struct {
	SourceInterface      SourceInterface    `json:"sourceInterface"`
	NveNeighborAddress   NveNeighborAddress `json:"nveNeighborAddress,omitempty"`
	NveVtepId            int                `json:"nveVtepId"`
	NveDestinationPort   int                `json:"nveDestinationPort"`
	NveEncapsulationType string             `json:"nveEncapsulationType,omitempty"`
	NveNeighborDiscoveryType string         `json:"nveNeighborDiscoveryType,omitempty"`
}
type VTEPPolicyResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	ID          string      `json:"id"`
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

type VTEPPoliciesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []VTEPPolicyResponse `json:"items"`
}

type VTEPPolicyUpdate struct {
	ID          string      `json:"id"`
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

type VTEPPolicy struct {
	NveEnable   bool        `json:"nveEnable"`
	VTEPEntries []VTEPEntry `json:"vtepEntries"`
	Type        string      `json:"type"`
}

func (v *Client) GetVTEPPolicies(ctx context.Context, deviceID string) (*VTEPPolicyResponse, error) {
	log.Printf("GetVTEPPolicies")

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, deviceID)

	log.Printf("GET FMC VTEP based on deviceid=%s", deviceID)
	log.Printf("Get FMC VTEP URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies: %s - %s", url, err.Error())
	}
	resp := &VTEPPoliciesResponse{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting vteppolicies: %s - %s", url, err.Error())
	}
	for _, vtepPolicy := range resp.Items {
		log.Printf("VTEP.ID=%v   VTEP.Type=%v", vtepPolicy.ID, vtepPolicy.Type)

		return &VTEPPolicyResponse{
			ID:   vtepPolicy.ID,
			Type: vtepPolicy.Type,
		}, nil
	}
	return nil, fmt.Errorf("no vteppolicies found for the device-ID%s", deviceID)
}

/*
*
This method return the all VTEP details based on VTEP ID
*/
func (v *Client) GetFmcVTEPDetails(ctx context.Context, deviceID string, vtepid string) (*VTEPPolicyResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies/%s", v.domainBaseURL, deviceID, vtepid)

	log.Printf("GET FMC VTEP details based on deviceid=%s and VTEP ID=%s", deviceID, vtepid)
	log.Printf("Get FMC VTEP details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for getting VTEP details: %s - %s", url, err.Error())
	}

	vtep := &VTEPPolicyResponse{}
	err = v.DoRequest(req, vtep, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while getting VTEP by id: %s - %s", url, err.Error())
	}

	return &VTEPPolicyResponse{
		ID:          vtep.ID,
		Type:        vtep.Type,
		NveEnable:   vtep.NveEnable,
		VTEPEntries: vtep.VTEPEntries,
	}, nil

}

/*
*
This method creates the VTEP
*/

func (v *Client) CreateFmcVTEP(ctx context.Context, deviceID string, object *VTEPPolicy) (*VTEPPolicyResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies", v.domainBaseURL, deviceID)

	log.Printf("Create VTEP URL=%s", url)

	body, err := json.Marshal(&object)

	log.Printf("Create VTEP Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}

	s2 := strings.ReplaceAll(string(body),`"literal":{},`,"" )
	s2 = strings.ReplaceAll(string(s2),`"nveNeighborAddress":{"object":{}},`,"" )
	jsonValue := []byte(s2)
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonValue))
	
	if err != nil {
		return nil, fmt.Errorf("error while creating context for VTEP: %s - %s", url, err.Error())
	}

	vtep := &VTEPPolicyResponse{}
	err = v.DoRequest(req, vtep, http.StatusCreated)

	if err != nil {
		return nil, fmt.Errorf("error while cretaing VTEP: %s - %s", url, err.Error())
	}

	log.Printf("VTEP interface created, response=%s", vtep.ID)

	return vtep, nil
}

/*
*
This method update the vtep details
*/

func (v *Client) UpdateFmcVTEP(ctx context.Context, deviceID string, id string, object *VTEPPolicyUpdate) (*VTEPPolicyResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies/%s", v.domainBaseURL, deviceID, id)

	log.Printf("Update VTEP URL=%s", url)

	body, err := json.Marshal(&object)

	log.Printf("Update VTEP Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("invalid request json : %s - %s", url, err.Error())
	}
	s2 := strings.ReplaceAll(string(body),`"literal":{},`,"" )
	s2 = strings.ReplaceAll(string(s2),`"nveNeighborAddress":{"object":{}},`,"" )
	jsonValue := []byte(s2)
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(jsonValue))

	if err != nil {
		return nil, fmt.Errorf("error while Updating context for VTEP: %s - %s", url, err.Error())
	}

	vtep := &VTEPPolicyResponse{}
	err = v.DoRequest(req, vtep, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Updating VTEP: %s - %s", url, err.Error())
	}

	log.Printf("VTEP Update, response=%s", vtep.ID)

	return vtep, nil
}

/*
*
This method delete the VTEP based in VTEP ID
*/
func (v *Client) DeleteFmcVTEPDetails(ctx context.Context, deviceID string, id string) (*VTEPPolicyResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/vteppolicies/%s", v.domainBaseURL, deviceID, id)

	log.Printf("Delete FMC VTEP details based on deviceid=%s and VTEP ID=%s", deviceID, id)
	log.Printf("Delete FMC VTEP details URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)

	if err != nil {
		return nil, fmt.Errorf("error while creating context for Delete VTEP details: %s - %s", url, err.Error())
	}

	vtepResp := &VTEPPolicyResponse{}
	err = v.DoRequest(req, vtepResp, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("error while Deleting VTEP by id: %s - %s", url, err.Error())
	}

	log.Printf("VTEP interface Delete, response=%s", vtepResp.ID)

	return vtepResp, nil
}
