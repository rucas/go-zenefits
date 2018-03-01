package zenefits

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	baseURL   = "https://api.zenefits.com"
	userAgent = "go-zenefits"
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	common        service // Reuse a single struct instead of allocating one for each service on the heap.
	People        *PeopleService
	Companies     *CompaniesService
	Departments   *DepartmentsService
	Employments   *EmploymentsService
	EmployeeBanks *EmployeeBanksService
	CompanyBanks  *CompanyBanksService
	Locations     *LocationsService
}

type service struct {
	client *Client
}

type Expansion struct {
	Includes []string `url:"includes,omitempty"`
}

type Ref struct {
	Url       string `json:"url"`
	Object    string `json:"object"`
	RefObject string `json:"ref_object"`
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURL)
	c := &Client{
		client:    httpClient,
		BaseURL:   baseURL,
		UserAgent: userAgent,
	}
	c.common.client = c

	// (*Point)(p) // p is converted to *Point
	// service is converted to people service and allocated
	c.People = (*PeopleService)(&c.common)
	c.Companies = (*CompaniesService)(&c.common)
	c.Departments = (*DepartmentsService)(&c.common)
	c.Employments = (*EmploymentsService)(&c.common)
	c.EmployeeBanks = (*EmployeeBanksService)(&c.common)
	c.CompanyBanks = (*CompanyBanksService)(&c.common)
	c.Locations = (*LocationsService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	url, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	// TODO: Need to work in the body for POST, PUT, ...etc requests
	var buf io.Reader
	req, err := http.NewRequest(method, url.String(), buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	type Page struct {
		Url     string      `json:"url"`
		NextUrl string      `json:"next_url"`
		Object  string      `json:"object"`
		Data    interface{} `json:"data"`
	}

	type MetaResponse struct {
		Status int    `json:"status"`
		Object string `json:"object"`
		Page   Page   `json:"data"`
	}

	valium := MetaResponse{Page: Page{Data: v}}
	err = json.NewDecoder(resp.Body).Decode(&valium)
	return resp, err
}

func AddOptions(s string, opt interface{}) (string, error) {
	v, err := query.Values(opt)

	// TODO: Check if it is a pointer...
	if err != nil {
		return s, err
	}

	u, err := url.Parse(s)

	if err != nil {
		return s, err
	}

	u.RawQuery = v.Encode()
	return u.String(), err
}
