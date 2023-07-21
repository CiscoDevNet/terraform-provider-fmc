package fmc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type SmartLicenseResponse struct {
	Links struct {
		Parent string `json:"parent,omitempty"`
		Self   string `json:"self"`
	} `json:"links"`
	Paging struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Count  int `json:"count"`
		Pages  int `json:"pages"`
	} `json:"paging"`
	Items []struct {
		RegStatus string `json:"regStatus"`
		Metadata  struct {
			LastSynchronizedTime string `json:"lastSynchronizedTime,omitempty"`
			EvalExpiresInDays    int    `json:"evalExpiresInDays,omitempty"`
			LastUser             struct {
				Name  string `json:"name,omitempty"`
				Links struct {
					Parent string `json:"parent,omitempty"`
					Self   string `json:"self"`
				} `json:"links,omitempty"`
				Id   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"lastUser,omitempty"`
			LastRenewedTime string `json:"lastRenewedTime,omitempty"`
			ExportControl   bool   `json:"exportControl,omitempty"`
			AuthStatus      string `json:"authStatus,omitempty"`
			Domain          struct {
				Name  string `json:"name,omitempty"`
				Links struct {
					Parent string `json:"parent,omitempty"`
					Self   string `json:"self"`
				} `json:"links,omitempty"`
				Id   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"domain,omitempty"`
			ReadOnly struct {
				Reason string `json:"reason,omitempty"`
				State  bool   `json:"state,omitempty"`
			} `json:"readOnly,omitempty"`
			VirtualAccount string   `json:"virtualAccount,omitempty"`
			Matches        []string `json:"matches,omitempty"`
			EvalUsed       bool     `json:"evalUsed,omitempty"`
			Timestamp      int      `json:"timestamp,omitempty"`
		} `json:"metadata,omitempty"`
		RegistrationType string `json:"registrationType,omitempty"`
		Name             string `json:"name,omitempty"`
		Description      string `json:"description,omitempty"`
		Links            struct {
			Parent string `json:"parent,omitempty"`
			Self   string `json:"self"`
		} `json:"links,omitempty"`
		Id    string `json:"id,omitempty"`
		Type  string `json:"type,omitempty"`
		Token string `json:"token,omitempty"`
	}
}

type SmartLicenseRequest struct {
	Type             string `json:"type"`
	RegistrationType string `json:"registrationType"`
	Token            string `json:"token,omitempty"`
}

type SmartLicense struct {
	Type               string `json:"type"`
	RegistrationType   string `json:"registrationType"`
	Token              string `json:"token,omitempty"`
	RegistrationStatus string `json:"regStatus,omitempty"`
	EvalUsed           bool   `json:"evalUsed,omitempty"`
	EvalExpiresInDays  int    `json:"evalExpiresInDays,omitempty"`
	VirtualAccount     string `json:"virtualAccount,omitempty"`
}

func (v *Client) CreateFmcSmartLicense(ctx context.Context, license *SmartLicenseRequest) (*SmartLicense, error) {
	if license.RegistrationType == "REGISTER" && license.Token == "" {
		return nil, fmt.Errorf("token required when registration_type = \"REGISTER\"")
	}

	url := fmt.Sprintf("%s/license/smartlicenses", v.fmcPlatformBaseURL)

	body, err := json.Marshal(&license)
	if err != nil {
		return nil, fmt.Errorf("creating smart license: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("creating smart license: %s - %s", url, err.Error())
	}

	item := &SmartLicense{}
	err = v.DoRequest(req, item, http.StatusCreated)

	return item, err
}

func (v *Client) GetFmcSmartLicense(ctx context.Context) (*SmartLicense, error) {
	url := fmt.Sprintf("%s/license/smartlicenses", v.fmcPlatformBaseURL)

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("getting smart license: %s - %s", url, err.Error())
	}

	smartLicenseRes := &SmartLicenseResponse{}
	err = v.DoRequest(req, smartLicenseRes, http.StatusOK)
	if err != nil {
		return nil, fmt.Errorf("getting smart license: %s - %s", url, err.Error())
	}

	smartLicense := &SmartLicense{}
	smartLicense.RegistrationStatus = smartLicenseRes.Items[0].RegStatus
	smartLicense.EvalUsed = smartLicenseRes.Items[0].Metadata.EvalUsed
	smartLicense.EvalExpiresInDays = smartLicenseRes.Items[0].Metadata.EvalExpiresInDays
	smartLicense.VirtualAccount = smartLicenseRes.Items[0].Metadata.VirtualAccount

	return smartLicense, nil
}

func (v *Client) DeleteFmcSmartLicense(ctx context.Context) error {
	url := fmt.Sprintf("%s/license/smartlicenses", v.fmcPlatformBaseURL)

	smartLicenseReq := &SmartLicenseRequest{
		RegistrationType: "DEREGISTER",
	}

	body, err := json.Marshal(&smartLicenseReq)
	if err != nil {
		return fmt.Errorf("creating smart license: %s - %s", url, err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("creating smart license: %s - %s", url, err.Error())
	}

	err = v.DoRequest(req, nil, http.StatusCreated)
	return err
}
