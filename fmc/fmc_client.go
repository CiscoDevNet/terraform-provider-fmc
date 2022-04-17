package fmc

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/juju/ratelimit"
)

// Mutex lock to disable parallelism on create/update/delete APIs
var nonReadMutex = &sync.Mutex{}

// Unlimiting GET requests did not seem to help with time, rather worsened the situation. So, back to only 1 for all requests.
var callSemaphore = make(semaphore, 1)

// Rate Limit at 100 requests per minute using token bucket that fills at a minute's interval.
var rateLimiterBucket = ratelimit.NewBucketWithQuantum(time.Minute, 100, 100)

type Client struct {
	user              string
	password          string
	host              string
	domainBaseURL     string
	accessToken       string
	domainUUID        string
	client            *http.Client
	ratelimiterBucket *ratelimit.Bucket
	nonReadMutex      *sync.Mutex
	callSemaphore     semaphore
}

type ErrorResponse struct {
	Error struct {
		Category string `json:"category"`
		Messages []struct {
			Description string `json:"description"`
		} `json:"messages"`
		Severity string `json:"severity"`
	} `json:"error"`
}

func NewClient(user, password, host string, insecureSkipVerify bool) *Client {
	return &Client{
		user:     user,
		password: password,
		host:     host,
		client: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
			Proxy: http.ProxyFromEnvironment,
		}},
		ratelimiterBucket: rateLimiterBucket,
		nonReadMutex:      nonReadMutex,
		callSemaphore:     callSemaphore,
	}
}

func (v *Client) Login() error {

	req, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/fmc_platform/v1/auth/generatetoken", v.host), nil)
	if err != nil {
		return (err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(v.user, v.password)

	res, err := v.client.Do(req)
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
	v.domainUUID = res.Header.Get("DOMAIN_UUID")
	v.domainBaseURL = fmt.Sprintf("https://%s/api/fmc_config/v1/domain/%s", v.host, v.domainUUID)
	return nil
}

func (v *Client) DoRequest(req *http.Request, item interface{}, status int) error {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Auth-Access-Token", v.accessToken)

	v.ratelimiterBucket.Wait(1) // This is a blocking call. Honors the rate limit by taking 1 token for this request.

	var r *http.Response
	var err error

	v.callSemaphore.Lock()
	if req.Method == "GET" {
		r, err = v.client.Do(req)
	} else {
		v.nonReadMutex.Lock()
		r, err = v.client.Do(req)
		v.nonReadMutex.Unlock()
	}
	v.callSemaphore.Unlock()

	if err != nil {
		return err
	}

	if status == 0 {
		status = http.StatusOK
	}

	// Handle 401 by logging in again
	if r.StatusCode == http.StatusUnauthorized {
		v.Login()
		return v.DoRequest(req, item, status)
	}

	// Handle 429 by sending it again, will go through the same token rate limiter
	if r.StatusCode == http.StatusTooManyRequests {
		return v.DoRequest(req, item, status)
	}

	if r.StatusCode != status {
		defer r.Body.Close()

		errorRes := ErrorResponse{}
		err = json.NewDecoder(r.Body).Decode(&errorRes)
		if err != nil {
			if body, err := ioutil.ReadAll(r.Body); err != nil {
				return fmt.Errorf("wrong status code: %d, could not read error body as error json and string, headers: %+v", r.StatusCode, r.Header)
			} else {
				return fmt.Errorf("wrong status code: %d, could not read error body as error json, body: %s, headers: %+v", r.StatusCode, body, r.Header)
			}
		}
		return fmt.Errorf("wrong status code: %d, error category: %s, error severity: %s, error messages: %v", r.StatusCode, errorRes.Error.Category, errorRes.Error.Severity, errorRes.Error.Messages)
	}
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
