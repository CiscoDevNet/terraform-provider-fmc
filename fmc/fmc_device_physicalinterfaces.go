package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type PhysicalinterfacesResponse struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Metadata struct {
		Timestamp int `json:"timestamp"`
		Domain    struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		}
		Commnet        string `json:"_comment"`
		SupportedSpeed string `json:"supportedSpeed"`
	} `json:"metadata"`
	Type              string   `json:"type"`
	Mode              string   `json:"mode"`
	HardWare          HardWare `json:"hardware"`
	LLDP              LLDP     `json:"LLDP"`
	PowerOverEthernet struct {
		ConsumptionWattage int  `json:"consumptionWattage"`
		State              bool `json:"state"`
	} `json:"powerOverEthernet"`
	Enabled        bool   `json:"enabled"`
	Description    string `json:"description"`
	MTU            int    `json:"MTU"`
	Priority       int    `json:"priority"`
	PathMonitoring struct {
		Enabled     bool   `json:"enabled"`
		Type        string `json:"type"`
		MonitoredIp string `json:"monitoredIp"`
	} `json:"pathMonitoring"`
	ManagementOnly bool `json:"managementOnly"`
	NveOnly        bool `json:"nveOnly"`
	SecurityZone   struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"securityZone"`
	Ifname                         string `json:"ifname"`
	EnableAntiSpoofing             bool   `json:"enableAntiSpoofing"`
	EnableSGTPropagate             bool   `json:"enableSGTPropagate"`
	ManagementAccess               bool   `json:"managementAccess"`
	ID                             string `json:"id"`
	Name                           string `json:"name"`
	OverrideDefaultFragmentSetting struct {
	} `json:"overrideDefaultFragmentSetting"`
	SwitchPortConfig struct {
		PortMode                string `json:"portMode"`
		TrunkModeAllowedVlanIds string `json:"trunkModeAllowedVlanIds"`
		TrunkModeNativeVlanId   int    `json:"trunkModeNativeVlanId"`
		ProtectedEnabled        bool   `json:"protectedEnabled"`
	} `json:"switchPortConfig"`
	IPv4 struct {
		DHCP struct {
			EnableDefaultRouteDHCP string `json:"enableDefaultRouteDHCP"`
			DhcpRouteMetric        int    `json:"dhcpRouteMetric"`
		} `json:"dhcp"`
	} `json:"ipv4"`
	IPv6 struct {
		EnableRA   string `json:"enableRA"`
		RaLifeTime string `json:"raLifeTime"`
		RaInterval string `json:"raInterval"`
	} `json:"ipv6"`
}

type PhysicalinterfacesRes struct {
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Items  []PhysicalinterfacesResponse `json:"items"`
	Paging struct {
		OffsetMTU int `json:"offset"`
		Limit     int `json:"limit"`
		Count     int `json:"count"`
		Pages     int `json:"pages"`
	} `json:"paging"`
}
type HardWare struct {
	Duplex          string `json:"duplex"`
	Speed           string `json:"speed"`
	AutoNegState    string `json:"autoNegState"`
	FecMode         string `json:"fecMode"`
	FlowControlSend string `json:"flowControlSend"`
}
type LLDP struct {
	Transmit bool `json:"transmit"`
	Receive  bool `json:"receive"`
}
type IPv6 struct {
	EnableIPV6 bool `json:"enableIPV6"`
}
type Physicalinterfaces struct {
	Type               string   `json:"type"`
	Enabled            bool     `json:"enabled"`
	MTU                int      `json:"MTU"`
	Name               string   `json:"name"`
	ID                 string   `json:"id"`
	Mode               string   `json:"mode"`
	HardWare           HardWare `json:"hardware"`
	LLDP               LLDP     `json:"LLDP"`
	ManagementOnly     bool     `json:"managementOnly"`
	NveOnly            bool     `json:"nveOnly"`
	EnableSGTPropagate bool     `json:"enableSGTPropagate"`
	IPv6               IPv6     `json:"ipv6"`
}

// api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/physicalinterfaces/{objectId}
func (v *Client) GetPhysicalInterfacesByName(ctx context.Context, id, name string) (*PhysicalinterfacesResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces", v.domainBaseURL, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting physicalinterfaces by name: %s - %s", url, err.Error())
	}
	resp := &PhysicalinterfacesRes{}
	err = v.DoRequest(req, resp, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting physicalinterfaces by name: %s - %s", url, err.Error())
	}
	switch l := len(resp.Items); {
	case l == 1:
		return v.GetPhysicalInterfaces(ctx, resp.Items[0].ID)
	case l > 1:
		for _, item := range resp.Items {
			if item.Name == name {
				return v.GetPhysicalInterfaces(ctx, item.ID)
			}
		}
		return nil, fmt.Errorf("duplicates found, no exact match, length of response is: %d, expected 1, please search using a unique id, name or value", l)
	case l == 0:
		return nil, fmt.Errorf("no access policies found, length of response is: %d, expected 1, please check your filter", l)
	}
	return nil, fmt.Errorf("no device found with name %s", name)
}

// api/fmc_config/v1/domain/{domainUUID}/devices/devicerecords/{containerUUID}/physicalinterfaces
func (v *Client) GetPhysicalInterfaces(ctx context.Context, acpId string) (*PhysicalinterfacesResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces", v.domainBaseURL, acpId)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting physicalinterfaces by name: %s - %s", url, err.Error())
	}
	physicalinterfacesResponse := &PhysicalinterfacesResponse{}
	err = v.DoRequest(req, physicalinterfacesResponse, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting physicalinterfaces by name: %s - %s", url, err.Error())
	}
	return physicalinterfacesResponse, fmt.Errorf("no device found with name %s", acpId)
}

// ​/api​/fmc_config​/v1​/domain​/{domainUUID}​/devices​/devicerecords​/{containerUUID}​/physicalinterfaces​/{objectId}
func (v *Client) UpdatePhysicalInterfaces(ctx context.Context, acpId, name string, physicalinterfaces *Physicalinterfaces) (*PhysicalinterfacesResponse, error) {
	url := fmt.Sprintf("%s/devices/devicerecords/%s/physicalinterfaces/%s", v.domainBaseURL, acpId, name)
	body, err := json.Marshal(&physicalinterfaces)
	if err != nil {
		return nil, fmt.Errorf("updating physicalinterfaces: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("updating physicalinterfaces: %s - %s", url, err.Error())
	}
	item := &PhysicalinterfacesResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("updating physicalinterfaces: %s - %s", url, err.Error())
	}
	return item, nil
}