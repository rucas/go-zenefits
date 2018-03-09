package zenefits

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

const (
	baseURL   = "https://api.zenefits.com"
	userAgent = "go-zenefits"
)

type Client struct {
	BaseURL   *url.URL
	UserAgent string

	client *http.Client

	// Reuse a single struct (service)
	// instead of allocating one for each
	// service on the heap.
	common        service
	People        *PeopleService
	Companies     *CompaniesService
	Departments   *DepartmentsService
	Employments   *EmploymentsService
	EmployeeBanks *EmployeeBanksService
	CompanyBanks  *CompanyBanksService
	Locations     *LocationsService
	Me            *MeService
}

type service struct {
	client *Client
}

// Response is a Zenefits response.
type Response struct {
	*http.Response

	NextPage int
	PrevPage int
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
	c.Companies = (*CompaniesService)(&c.common)
	c.CompanyBanks = (*CompanyBanksService)(&c.common)
	c.Departments = (*DepartmentsService)(&c.common)
	c.EmployeeBanks = (*EmployeeBanksService)(&c.common)
	c.Employments = (*EmploymentsService)(&c.common)
	c.Locations = (*LocationsService)(&c.common)
	c.Me = (*MeService)(&c.common)
	c.People = (*PeopleService)(&c.common)
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

func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	req = req.WithContext(ctx)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	nresp := newResponse(resp, v)
	return nresp, err
}

func newResponse(r *http.Response, v interface{}) *Response {
	resp := &Response{Response: r}
	if meta, ok := v.(*MetaResponse); ok && meta.Page.NextUrl != "" {
		u, _ := url.Parse(meta.Page.NextUrl)
		s := u.Query().Get("starting_after")
		b := u.Query().Get("ending_before")
		resp.NextPage, _ = strconv.Atoi(s)
		resp.PrevPage, _ = strconv.Atoi(b)
	}
	return resp
}

func addOptions(s string, opt interface{}) (string, error) {
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
