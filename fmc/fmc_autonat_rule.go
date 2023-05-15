package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AutoNatRuleSubConfig struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Name string `json:"name"`
}

type AutoNatRulePatOptions struct {
	Patpooladdress *AutoNatRuleSubConfig `json:"patPoolAddress"`
	// Blockallocation bool                 `json:"blockAllocation"`
	// Flatportrange   bool                 `json:"flatPortRange"`
	Includereserve bool `json:"includeReserve"`
	Interfacepat   bool `json:"interfacePat"`
	Extendedpat    bool `json:"extendedPat"`
	Roundrobin     bool `json:"roundRobin"`
}

type AutoNatRule struct {
	ID                           string                 `json:"id,omitempty"`
	Type                         string                 `json:"type"`
	Description                  string                 `json:"description"`
	Nattype                      string                 `json:"natType"`
	Sourceinterface              *AutoNatRuleSubConfig  `json:"sourceInterface,omitempty"`
	Destinationinterface         *AutoNatRuleSubConfig  `json:"destinationInterface,omitempty"`
	Originalnetwork              *AutoNatRuleSubConfig  `json:"originalNetwork,omitempty"`
	Translatednetwork            *AutoNatRuleSubConfig  `json:"translatedNetwork,omitempty"`
	Interfaceintranslatednetwork bool                   `json:"interfaceInTranslatedNetwork,omitempty"`
	Originalport                 int                    `json:"originalPort,omitempty"`
	Serviceprotocol              string                 `json:"serviceProtocol,omitempty"`
	Translatedport               int                    `json:"translatedPort,omitempty"`
	Fallthrough                  bool                   `json:"fallThrough"`
	DNS                          bool                   `json:"dns"`
	Noproxyarp                   bool                   `json:"noProxyArp"`
	Routelookup                  bool                   `json:"routeLookup"`
	Nettonet                     bool                   `json:"netToNet"`
	Interfaceipv6                bool                   `json:"interfaceIpv6"`
	Patoptions                   *AutoNatRulePatOptions `json:"patOptions,omitempty"`
}

type AutoNatRuleResponse struct {
	ID                           string                `json:"id"`
	Description                  string                `json:"description"`
	Serviceprotocol              string                `json:"serviceProtocol"`
	Originalnetwork              AutoNatRuleSubConfig  `json:"originalNetwork"`
	Translatednetwork            AutoNatRuleSubConfig  `json:"translatedNetwork"`
	Sourceinterface              AutoNatRuleSubConfig  `json:"sourceInterface"`
	Destinationinterface         AutoNatRuleSubConfig  `json:"destinationInterface"`
	Interfaceintranslatednetwork bool                  `json:"interfaceInTranslatedNetwork"`
	Translatedport               int                   `json:"translatedPort"`
	Originalport                 int                   `json:"originalPort"`
	Type                         string                `json:"type"`
	Nattype                      string                `json:"natType"`
	Interfaceipv6                bool                  `json:"interfaceIpv6"`
	Fallthrough                  bool                  `json:"fallThrough"`
	DNS                          bool                  `json:"dns"`
	Routelookup                  bool                  `json:"routeLookup"`
	Noproxyarp                   bool                  `json:"noProxyArp"`
	Nettonet                     bool                  `json:"netToNet"`
	Patoptions                   AutoNatRulePatOptions `json:"patOptions"`
}

// /fmc_config/v1/domain/DomainUUID/policy/ftdnatpolicies/{containerUUID}/autonatrules?bulk=true ( Bulk POST operation on auto nat rules. )

func (v *Client) CreateFmcAutoNatRule(ctx context.Context, natId string, autoNatRule *AutoNatRule) (*AutoNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/autonatrules", v.domainBaseURL, natId)
	body, err := json.Marshal(&autoNatRule)
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s", url, err.Error())
	}
	item := &AutoNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s\n%s", url, err.Error(), body)
	}
	return item, nil
}

func (v *Client) GetFmcAutoNatRule(ctx context.Context, natId string, id string) (*AutoNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/autonatrules/%s", v.domainBaseURL, natId, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting auto nat rules: %s - %s", url, err.Error())
	}
	item := &AutoNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting auto nat rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcAutoNatRule(ctx context.Context, natId, id string, autoNatRule *AutoNatRule) (*AutoNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/autonatrules/%s", v.domainBaseURL, natId, id)
	body, err := json.Marshal(&autoNatRule)
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s", url, err.Error())
	}
	item := &AutoNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("creating auto nat rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcAutoNatRule(ctx context.Context, natId string, id string) error {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/autonatrules/%s", v.domainBaseURL, natId, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting auto nat rules: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
