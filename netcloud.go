package netcloud

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client ...
type Client struct {
	baseURL   string
	cpApiID   string
	cpApiKey  string
	ecmApiID  string
	ecmApiKey string
	http      *http.Client
}

// NewClient ...
func NewClient(apiAuth AuthParams) *Client {
	return &Client{
		baseURL:   "https://www.cradlepointecm.com/api/v2",
		cpApiID:   apiAuth.CpApiID,
		cpApiKey:  apiAuth.CpApiKey,
		ecmApiID:  apiAuth.EcmApiID,
		ecmApiKey: apiAuth.EcmApiKey,
		http: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) generateReq(uri, method string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseURL+uri, body)
	if err != nil {
		return req, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-CP-API-ID", c.cpApiID)
	req.Header.Set("X-CP-API-KEY", c.cpApiKey)
	req.Header.Set("X-ECM-API-ID", c.ecmApiID)
	req.Header.Set("X-ECM-API-KEY", c.ecmApiKey)
	return req, nil
}

func (c *Client) makeRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.http.Do(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// GetRouter ...
func (c *Client) GetRouter(r RouterReqParams) ([]Router, error) {
	var rtrReq RouterReq
	endpoint := "/routers?limit=349"
	if r.ID != "" {
		endpoint += fmt.Sprintf("&id=%s", r.ID)
	}
	if r.Ipv4Addr != "" {
		endpoint += fmt.Sprintf("&id=%s", r.Ipv4Addr)
	}
	if r.MacAddr != "" {
		endpoint += fmt.Sprintf("&id=%s", r.MacAddr)
	}
	if r.Name != "" {
		endpoint += fmt.Sprintf("&id=%s", r.Name)
	}
	req, err := c.generateReq(endpoint, "GET", nil)
	if err != nil {
		return rtrReq.Data, err
	}
	res, err := c.makeRequest(req)
	if err != nil {
		return rtrReq.Data, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&rtrReq); err != nil {
		return rtrReq.Data, err
	}
	return rtrReq.Data, nil
}

// func (c *Client) createReqBody(v interface{}) (*bytes.Reader, error) {
// 	payload, err := json.Marshal(&v)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bytes.NewReader(payload), nil
// }
