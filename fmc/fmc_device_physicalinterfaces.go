package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PhysicalInterfaceSecurityZone struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type PhysicalInterfaceResponse struct {
	Type         string                        `json:"type"`
	Ifname       string                        `json:"ifname"`
	Description  string                        `json:"description"`
	ID           string                        `json:"id"`
	Name         string                        `json:"name"`
	MTU          int64                         `json:"mtu"`
	Mode         string                        `json:"mode"`
	SecurityZone PhysicalInterfaceSecurityZone `json:"securityZone"`
}

type PhysicalInterfaceRequest struct {
	Type         string                        `json:"type"`
	Ifname       string                        `json:"ifname"`
	Description  string                        `json:"description"`
	ID           string                        `json:"id"`
	Name         string                        `json:"name"`
	MTU          int                           `json:"MTU"`
	Mode         string                        `json:"mode"`
	SecurityZone PhysicalInterfaceSecurityZone `json:"securityZone"`
}

type PhysicalInterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items []PhysicalInterfaceResponse `json:"items"`
}

func (v *Client) GetFmcPhysicalInterface(ctx context.Context, deviceID string, name string) (*PhysicalInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces?expanded=true", v.domainBaseURL, deviceID)
	log.Printf("GET FMC Physical interface details based on deviceid=%s and physical interface name=%s", deviceID, name)
	log.Printf("Get FMC PhysicalInterface URL=%s", url)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)

	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}

	physicalInterfaces := &PhysicalInterfacesResponse{}
	err = v.DoRequest(req, physicalInterfaces, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	for _, PhysicalInterface := range physicalInterfaces.Items {
		log.Printf("PhysicalInterface.ID=%s  PhysicalInterface.Name=%s  PhysicalInterface.Type=%s", PhysicalInterface.ID, PhysicalInterface.Name, PhysicalInterface.Type)
		if PhysicalInterface.Name == name {
			return &PhysicalInterfaceResponse{
				ID:           PhysicalInterface.ID,
				Name:         PhysicalInterface.Name,
				Ifname:       PhysicalInterface.Ifname,
				Type:         PhysicalInterface.Type,
				Description:  PhysicalInterface.Description,
				MTU:          PhysicalInterface.MTU,
				Mode:         PhysicalInterface.Mode,
				SecurityZone: PhysicalInterface.SecurityZone,
			}, nil
		}
	}

	return nil, fmt.Errorf("no Interface found with name %s", name)
}

func (v *Client) GetFmcPhysicalInterfaceByID(ctx context.Context, deviceID string, physicalInterfaceId string) (*PhysicalInterfaceResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces/%s", v.domainBaseURL, deviceID, physicalInterfaceId)
	log.Printf("Get physical interface by id URL=%s", url)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	physicalInterface := &PhysicalInterfaceResponse{}
	err = v.DoRequest(req, physicalInterface, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}

	log.Printf("Get physical interface by name=%s", physicalInterface.Name)

	if physicalInterface != nil {
		return &PhysicalInterfaceResponse{
			ID:           physicalInterface.ID,
			Name:         physicalInterface.Name,
			Ifname:       physicalInterface.Ifname,
			Type:         physicalInterface.Type,
			Description:  physicalInterface.Description,
			MTU:          physicalInterface.MTU,
			Mode:         physicalInterface.Mode,
			SecurityZone: physicalInterface.SecurityZone,
		}, nil
	}

	return nil, fmt.Errorf("no Interface found with physicalInterfaceId %s", physicalInterfaceId)
}

func (v *Client) UpdateFmcPhysicalInterface(ctx context.Context, deviceID string, id string, object *PhysicalInterfaceRequest) (*PhysicalInterfaceResponse, error) {

	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces/%s", v.domainBaseURL, deviceID, id)
	log.Printf("Update physical interface URL=%s", url)
	body, err := json.Marshal(&object)
	log.Printf("Update physical interface Request=%s", body)

	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating physical interfaces: %s - %s", url, err.Error())
	}

	item := &PhysicalInterfaceResponse{}
	err = v.DoRequest(req, item, http.StatusOK)

	if err != nil {
		return nil, fmt.Errorf("getting physical interfaces: %s - %s", url, err.Error())
	}
	log.Printf("Physical interface updated, response=%s", item)

	return item, nil
}
