package fmc

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	insecureSkipVerify bool
	user               string
	password           string
	host               string
	domainBaseURL      string
	accessToken        string
	DomainUUID         string
}

func NewClient(user, password, host string, insecureSkipVerify bool) *Client {
	return &Client{
		insecureSkipVerify: insecureSkipVerify,
		user:               user,
		password:           password,
		host:               host,
		domainBaseURL:      "",
	}
}

func (v *Client) Login() error {

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/fmc_platform/v1/auth/generatetoken", v.host), nil)
	if err != nil {
		return (err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(v.user, v.password)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: v.insecureSkipVerify},
	}
	c := &http.Client{Transport: tr}

	res, err := c.Do(req)
	if err != nil {
		return (err)
	}
	defer res.Body.Close()
	if res.StatusCode == 401 {
		return fmt.Errorf("wrong username or password %d %v", res.StatusCode, req.URL)
	} else if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("cannot login unknown error, status code: %d %v", res.StatusCode, req.URL)
	}

	v.accessToken = res.Header.Get("X-Auth-Access-Token")
	v.DomainUUID = res.Header.Get("DOMAIN_UUID")
	v.domainBaseURL = fmt.Sprintf("https://%s/api/fmc_config/v1/domain/%s", v.host, v.DomainUUID)
	return nil
}

func (v *Client) DoRequest(req *http.Request, item interface{}, status int) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: v.insecureSkipVerify},
	}
	c := &http.Client{Transport: tr}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	r, err := c.Do(req)

	if err != nil {
		return err
	}
	if status == 0 {
		status = http.StatusOK
	}
	if r.StatusCode != status {
		return fmt.Errorf("wrong status code: %d", r.StatusCode)
	}
	// Handle 429 if any
	log.Printf("Status code: %d", r.StatusCode)
	if item != nil {
		defer r.Body.Close()
		err = json.NewDecoder(r.Body).Decode(item)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (v *Client) DeleteURLObject(ctx context.Context, id string) error {
// 	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
// 	req, err := http.NewRequestWithContext(ctx, "DELETE", url, nil)
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, err.Error()))
// 	}
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	c := &http.Client{Transport: tr}

// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Set("X-Auth-Access-Token", v.accessToken)

// 	r, err := c.Do(req)
// 	if err != nil {
// 		return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, err.Error()))
// 	}
// 	if r.StatusCode == http.StatusOK {
// 		return nil
// 	}
// 	// if r.StatusCode == http.StatusTooManyRequests handle retries

// 	return errors.New(fmt.Sprintf("deleting url objects: %s - %s", url, r.Status))
// }

// func (v *Client) GetURLObject(ctx context.Context, id string) (*URLObjectResponse, error) {
// 	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
// 	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
// 	}
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	c := &http.Client{Transport: tr}

// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Set("X-Auth-Access-Token", v.accessToken)

// 	r, err := c.Do(req)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
// 	}

// 	defer r.Body.Close()

// 	item := &URLObjectResponse{}
// 	err = json.NewDecoder(r.Body).Decode(item)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("getting url objects: %s - %s", url, err.Error()))
// 	}
// 	return item, nil
// }

// func (v *Client) CreateURLObject(ctx context.Context, object *URLObject) (*URLObjectResponse, error) {
// 	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls?bulk=true", v.baseURL, v.DomainUUID)
// 	body, err := json.Marshal(&object)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
// 	}
// 	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
// 	}
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	c := &http.Client{Transport: tr}

// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Set("X-Auth-Access-Token", v.accessToken)

// 	r, err := c.Do(req)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
// 	}

// 	defer r.Body.Close()

// 	item := &URLObjectResponse{}
// 	err = json.NewDecoder(r.Body).Decode(item)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("creating url objects: %s - %s", url, err.Error()))
// 	}
// 	return item, nil
// }

// func (v *Client) UpdateURLObject(ctx context.Context, id string, object *URLObjectUpdateInput) (*URLObjectResponse, error) {
// 	url := fmt.Sprintf("%s/api/fmc_config/v1/domain/%s/object/urls/%s", v.baseURL, v.DomainUUID, id)
// 	body, err := json.Marshal(&object)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
// 	}
// 	req, err := http.NewRequestWithContext(ctx, "PUT", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
// 	}
// 	tr := &http.Transport{
// 		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
// 	}
// 	c := &http.Client{Transport: tr}

// 	req.Header.Set("Content-Type", "application/json; charset=utf-8")
// 	req.Header.Set("X-Auth-Access-Token", v.accessToken)

// 	r, err := c.Do(req)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
// 	}

// 	defer r.Body.Close()

// 	item := &URLObjectResponse{}
// 	err = json.NewDecoder(r.Body).Decode(item)
// 	if err != nil {
// 		return nil, errors.New(fmt.Sprintf("updating url objects: %s - %s", url, err.Error()))
// 	}
// 	return item, nil
// }
