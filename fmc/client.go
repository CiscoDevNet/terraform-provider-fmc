package fmc

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Client struct {
	InsecureSkipVerify bool
	user               string
	password           string
	baseURL            string
	accessToken        string
	DomainUUID         string
}

func NewClient(user, password, baseURL string) *Client {
	return &Client{
		InsecureSkipVerify: true,
		user:               user,
		password:           password,
		baseURL:            baseURL,
	}
}

func (v *Client) Login() error {

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/api/fmc_platform/v1/auth/generatetoken", v.baseURL), nil)
	if err != nil {
		return (err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(v.user, v.password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	res, err := c.Do(req)
	if err != nil {
		return (err)
	}
	defer res.Body.Close()
	if res.StatusCode == 401 {
		return (fmt.Errorf("wrong username or password %d %v", res.StatusCode, req.URL))
	} else if res.StatusCode != http.StatusNoContent {
		return (fmt.Errorf("cannot login unknown error, status code: %d %v", res.StatusCode, req.URL))
	}

	v.accessToken = res.Header.Get("X-Auth-Access-Token")
	v.DomainUUID = res.Header.Get("DOMAIN_UUID")
	return nil
}

// /fmc_config/v1/domain/DomainUUID/object/networks?bulk=true ( Bulk POST operation on network objects. )

func (v *Client) DeleteNetworkObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/networks/%s", v.baseURL, v.DomainUUID, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("deleting network objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("deleting network objects: %s - %s", url, err.Error()))
	}
	if r.StatusCode == http.StatusOK {
		return nil
	}
	// if r.StatusCode == http.StatusTooManyRequests handle retries

	return errors.New(fmt.Sprintf("deleting network objects: %s - %s", url, r.Status))
}

func (v *Client) GetNetworkObject(ctx context.Context, id string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/networks/%s", v.baseURL, v.DomainUUID, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting network objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting network objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &NetworkObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting network objects: %s - %s", url, err.Error()))
	}
	return item, nil
}

// func (v *Client) CreateNetworkObject(ctx context.Context, object *NetworkObject) (map[string]interface{}, error) {
func (v *Client) CreateNetworkObject(ctx context.Context, object *NetworkObject) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/networks", v.baseURL, v.DomainUUID)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating network objects: %s - %s", url, err.Error()))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating network objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating network objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &NetworkObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating network objects: %s - %s", url, err.Error()))
	}
	return item, nil
}

func (v *Client) UpdateNetworkObject(ctx context.Context, id string, object *NetworkObjectUpdateInput) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/networks/%s", v.baseURL, v.DomainUUID, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating network objects: %s - %s", url, err.Error()))
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating network objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating network objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &NetworkObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating network objects: %s - %s", url, err.Error()))
	}
	return item, nil
}

func (v *Client) DeleteURLObject(ctx context.Context, id string) error {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, err.Error()))
	}
	if r.StatusCode == http.StatusOK {
		return nil
	}
	// if r.StatusCode == http.StatusTooManyRequests handle retries

	return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, r.Status))
}

func (v *Client) GetNetworkObject(ctx context.Context, id string) (*NetworkObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &URLObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
	}
	return item, nil
}

func (v *Client) CreateURLObject(ctx context.Context, object *URLObject) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls", v.baseURL, v.DomainUUID)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
	}
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &URLObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
	}
	return item, nil
}


func (v *Client) UpdateURLObject(ctx context.Context, id string, object *URLObjectUpdateInput) (*URLObjectResponse, error) {
	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
	body, err := json.Marshal(&object)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
	}
	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	c := &http.Client{Transport: tr}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
	}

	defer r.Body.Close()

	item := &URLObjectResponse{}
	err = json.NewDecoder(r.Body).Decode(item)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
	}
	return item, nil
}