package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AccessRuleSubConfig struct {
	ID   string `json:"id"`
	Type string `json:"type,omitempty"`
}

type AccessRuleSubConfigs struct {
	Objects []AccessRuleSubConfig `json:"objects"`
}

type AccessRuleDefaultAction struct {
	Intrusionpolicy AccessRuleSubConfig `json:"intrusionPolicy"`
	Syslogconfig    AccessRuleSubConfig `json:"syslogConfig"`
	Type            string              `json:"type"`
	Logbegin        string              `json:"logBegin"`
	Logend          string              `json:"logEnd"`
	Sendeventstofmc string              `json:"sendEventsToFMC"`
	Action          string              `json:"action"`
	// Variableset struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"variableSet"`
	// Snmpconfig struct {
	// 	ID   string `json:"id"`
	// 	Type string `json:"type"`
	// } `json:"snmpConfig"`
}

type AccessRule struct {
	ID                  string               `json:"id,omitempty"`
	Name                string               `json:"name"`
	Type                string               `json:"type"`
	Action              string               `json:"action"`
	Syslogseverity      string               `json:"syslogSeverity,omitempty"`
	Enablesyslog        bool                 `json:"enableSyslog"`
	Enabled             bool                 `json:"enabled"`
	Sendeventstofmc     bool                 `json:"sendEventsToFMC"`
	Logfiles            bool                 `json:"logFiles"`
	Logbegin            bool                 `json:"logBegin"`
	Logend              bool                 `json:"logEnd"`
	Sourcezones         AccessRuleSubConfigs `json:"sourceZones,omitempty"`
	Destinationzones    AccessRuleSubConfigs `json:"destinationZones,omitempty"`
	Sourcenetworks      AccessRuleSubConfigs `json:"sourceNetworks,omitempty"`
	Destinationnetworks AccessRuleSubConfigs `json:"destinationNetworks,omitempty"`
	Sourceports         AccessRuleSubConfigs `json:"sourcePorts,omitempty"`
	Destinationports    AccessRuleSubConfigs `json:"destinationPorts,omitempty"`
	Urls                AccessRuleSubConfigs `json:"urls,omitempty"`
	Ipspolicy           *AccessRuleSubConfig `json:"ipsPolicy,omitempty"`
	Filepolicy          *AccessRuleSubConfig `json:"filePolicy,omitempty"`
	Syslogconfig        *AccessRuleSubConfig `json:"syslogConfig,omitempty"`
	Newcomments         []string             `json:"newComments,omitempty"`
}

type AccessRuleUpdate AccessRule

type AccessRuleResponseObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Type        string `json:"type"`
	Version     string `json:"version"`
}

type AccessRuleResponse struct {
	Sourcenetworks struct {
		Objects  []AccessRuleResponseObject `json:"objects"`
		Literals []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"literals"`
	} `json:"sourceNetworks"`
	Syslogseverity string `json:"syslogSeverity"`
	Sourcezones    struct {
		Objects []AccessRuleResponseObject `json:"objects"`
	} `json:"sourceZones"`
	Destinationzones struct {
		Objects []AccessRuleResponseObject `json:"objects"`
	} `json:"destinationZones"`
	Description            string `json:"description"`
	Originalsourcenetworks struct {
		Objects  []AccessRuleResponseObject `json:"objects"`
		Literals []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"literals"`
	} `json:"originalSourceNetworks"`
	Type         string `json:"type"`
	Enablesyslog bool   `json:"enableSyslog"`
	Enabled      bool   `json:"enabled"`
	Urls         struct {
		Objects                     []AccessRuleResponseObject `json:"objects"`
		Urlcategorieswithreputation []struct {
			Reputation string                   `json:"reputation"`
			Type       string                   `json:"type"`
			Category   AccessRuleResponseObject `json:"category"`
		} `json:"urlCategoriesWithReputation"`
		Literals []struct {
			Type string `json:"type"`
			URL  string `json:"url"`
		} `json:"literals"`
	} `json:"urls"`
	Syslogconfig        AccessRuleResponseObject `json:"syslogConfig"`
	Destinationnetworks struct {
		Objects []AccessRuleResponseObject `json:"objects"`
	} `json:"destinationNetworks"`
	Action           string `json:"action"`
	ID               string `json:"id"`
	Logend           bool   `json:"logEnd"`
	Logbegin         bool   `json:"logBegin"`
	Sendeventstofmc  bool   `json:"sendEventsToFMC"`
	Destinationports struct {
		Objects []AccessRuleResponseObject `json:"objects"`
	} `json:"destinationPorts"`
	Sourceports struct {
		Objects []AccessRuleResponseObject `json:"objects"`
	} `json:"sourcePorts"`
	Version     string                   `json:"version"`
	Variableset AccessRuleResponseObject `json:"variableSet"`
	Logfiles    bool                     `json:"logFiles"`
	Filepolicy  AccessRuleResponseObject `json:"filePolicy"`
	Ipspolicy   AccessRuleResponseObject `json:"ipsPolicy"`
	Name        string                   `json:"name"`
}

// /fmc_config/v1/domain/DomainUUID/policy/accesspolicies/{containerUUID}/accessrules?bulk=true ( Bulk POST operation on access rules. )

func (v *Client) CreateFmcAccessRule(ctx context.Context, acpId, section, insertBefore, insertAfter, category string, accessPolicy *AccessRule) (*AccessRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/accessrules", v.domainBaseURL, acpId)
	initialSet := false
	if section != "" {
		url = fmt.Sprintf("%s?section=%s", url, section)
		initialSet = true
	}
	if category != "" {
		url = fmt.Sprintf("%s?category=%s", url, category)
		initialSet = true
	}
	if insertBefore != "" {
		if initialSet {
			url = fmt.Sprintf("%s&insertBefore=%s", url, insertBefore)
		} else {
			url = fmt.Sprintf("%s?insertBefore=%s", url, insertBefore)
			initialSet = true
		}
	}
	if insertAfter != "" {
		if initialSet {
			url = fmt.Sprintf("%s&insertAfter=%s", url, insertAfter)
		} else {
			url = fmt.Sprintf("%s?insertAfter=%s", url, insertAfter)
			initialSet = true
		}
	}
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s", url, err.Error())
	}
	item := &AccessRuleResponse{}
	err = v.DoRequest(req, item, http.StatusCreated)
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s, %s", url, err.Error(), body)
	}
	return item, nil
}

func (v *Client) GetFmcAccessRule(ctx context.Context, acpId string, id string) (*AccessRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/accessrules/%s", v.domainBaseURL, acpId, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting access rules: %s - %s", url, err.Error())
	}
	item := &AccessRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting access rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) UpdateFmcAccessRule(ctx context.Context, acpId, id string, accessPolicy *AccessRuleUpdate) (*AccessRuleResponse, error) {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/accessrules/%s", v.domainBaseURL, acpId, id)
	body, err := json.Marshal(&accessPolicy)
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s", url, err.Error())
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s", url, err.Error())
	}
	item := &AccessRuleResponse{}
	err = v.DoRequest(req, item, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("creating access rules: %s - %s", url, err.Error())
	}
	return item, nil
}

func (v *Client) DeleteFmcAccessRule(ctx context.Context, acpId string, id string) error {
	url := fmt.Sprintf("%s/policy/accesspolicies/%s/accessrules/%s", v.domainBaseURL, acpId, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("deleting access rules: %s - %s", url, err.Error())
	}
	err = v.DoRequest(req, nil, http.StatusOK)
	return err
}
