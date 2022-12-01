package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ManualNatRuleSubConfig struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ManualNatRulePatOptions struct {
	// Blockallocation bool                 `json:"blockAllocation"`
	// Flatportrange   bool                 `json:"flatPortRange"`
	Patpooladdress ManualNatRuleSubConfig `json:"patPoolAddress"`
	Includereserve bool                   `json:"includeReserve"`
	Interfacepat   bool                   `json:"interfacePat"`
	Extendedpat    bool                   `json:"extendedPat"`
	Roundrobin     bool                   `json:"roundRobin"`
}

type ManualNatRule struct {
	ID                             string                   `json:"id,omitempty"`
	Type                           string                   `json:"type"`
	Description                    string                   `json:"description"`
	Enabled                        bool                     `json:"enabled"`
	Nattype                        string                   `json:"natType"`
	Sourceinterface                *ManualNatRuleSubConfig  `json:"sourceInterface,omitempty"`
	Destinationinterface           *ManualNatRuleSubConfig  `json:"destinationInterface,omitempty"`
	Interfaceinoriginaldestination bool                     `json:"interfaceInOriginalDestination"`
	Interfaceintranslatedsource    bool                     `json:"interfaceInTranslatedSource"`
	Originaldestination            *ManualNatRuleSubConfig  `json:"originalDestination,omitempty"`
	Originaldestinationport        *ManualNatRuleSubConfig  `json:"originalDestinationPort,omitempty"`
	Originalsource                 *ManualNatRuleSubConfig  `json:"originalSource,omitempty"`
	Originalsourceport             *ManualNatRuleSubConfig  `json:"originalSourcePort,omitempty"`
	Translateddestination          *ManualNatRuleSubConfig  `json:"translatedDestination,omitempty"`
	Translateddestinationport      *ManualNatRuleSubConfig  `json:"translatedDestinationPort,omitempty"`
	Translatedsource               *ManualNatRuleSubConfig  `json:"translatedSource,omitempty"`
	Translatedsourceport           *ManualNatRuleSubConfig  `json:"translatedSourcePort,omitempty"`
	Unidirectional                 bool                     `json:"unidirectional"`
	Fallthrough                    bool                     `json:"fallThrough"`
	DNS                            bool                     `json:"dns"`
	Noproxyarp                     bool                     `json:"noProxyArp"`
	Routelookup                    bool                     `json:"routeLookup"`
	Nettonet                       bool                     `json:"netToNet"`
	Interfaceipv6                  bool                     `json:"interfaceIpv6"`
	Patoptions                     *ManualNatRulePatOptions `json:"patOptions,omitempty"`
}

type ManualNatRuleResponse struct {
	ID                             string                  `json:"id"`
	Type                           string                  `json:"type"`
	Description                    string                  `json:"description"`
	Enabled                        bool                    `json:"enabled"`
	Nattype                        string                  `json:"natType"`
	Sourceinterface                ManualNatRuleSubConfig  `json:"sourceInterface"`
	Destinationinterface           ManualNatRuleSubConfig  `json:"destinationInterface"`
	Interfaceinoriginaldestination bool                    `json:"interfaceInOriginalDestination"`
	Interfaceintranslatedsource    bool                    `json:"interfaceInTranslatedSource"`
	Originaldestination            ManualNatRuleSubConfig  `json:"originalDestination"`
	Originaldestinationport        ManualNatRuleSubConfig  `json:"originalDestinationPort"`
	Originalsource                 ManualNatRuleSubConfig  `json:"originalSource"`
	Originalsourceport             ManualNatRuleSubConfig  `json:"originalSourcePort"`
	Translateddestination          ManualNatRuleSubConfig  `json:"translatedDestination"`
	Translateddestinationport      ManualNatRuleSubConfig  `json:"translatedDestinationPort"`
	Translatedsource               ManualNatRuleSubConfig  `json:"translatedSource"`
	Translatedsourceport           ManualNatRuleSubConfig  `json:"translatedSourcePort"`
	Unidirectional                 bool                    `json:"unidirectional"`
	Fallthrough                    bool                    `json:"fallThrough"`
	DNS                            bool                    `json:"dns"`
	Noproxyarp                     bool                    `json:"noProxyArp"`
	Routelookup                    bool                    `json:"routeLookup"`
	Nettonet                       bool                    `json:"netToNet"`
	Interfaceipv6                  bool                    `json:"interfaceIpv6"`
	Patoptions                     ManualNatRulePatOptions `json:"patOptions"`
}

// /fmc_config/v1/domain/DomainUUID/policy/ftdnatpolicies/{containerUUID}/manualnatrules?bulk=true ( Bulk POST operation on manual nat rules. )

func (v *Client) CreateFmcManualNatRule(ctx context.Context, natId, section, targetIndex string, manualNatRule *ManualNatRule) (*ManualNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/manualnatrules", v.domainBaseURL, natId)
	initialSet := false
	if section != "" {
		url = fmt.Sprintf("%s?section=%s", url, section)
		initialSet = true
	}
	if targetIndex != "" {
		if initialSet {
			url = fmt.Sprintf("%s&targetIndex=%s", url, targetIndex)
		} else {
			url = fmt.Sprintf("%s?targetIndex=%s", url, targetIndex)
			initialSet = true
		}
	}
	body, err := json.Marshal(&manualNatRule)
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	item := &ManualNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) GetFmcManualNatRule(ctx context.Context, natId string, id string) (*ManualNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/manualnatrules/%s", v.domainBaseURL, natId, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting manual nat rules: %s - %s", url, err.Error())
	}
	item := &ManualNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return item, fmt.Errorf("getting manual nat rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcManualNatRule(ctx context.Context, natId, id string, manualNatRule *ManualNatRule) (*ManualNatRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/manualnatrules/%s", v.domainBaseURL, natId, id)
	body, err := json.Marshal(&manualNatRule)
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	item := &ManualNatRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("creating manual nat rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcManualNatRule(ctx context.Context, natId string, id string) error {
	url := fmt.Sprintf("%s/policy/ftdnatpolicies/%s/manualnatrules/%s", v.domainBaseURL, natId, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting manual nat rules: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
