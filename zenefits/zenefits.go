package zenefits

import (
	"net/http"
	"net/url"
)

const (
	baseURL   = "https://api.zenefits.com"
	userAgent = "go-zenefits"
)

type Client struct {
	client    *http.Client
	BaseURL   *url.URL
	UserAgent string

	// TODO:
	// Add services heres...
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	baseURL, _ := url.Parse(baseURL)
	return &Client{client: httpClient, BaseURL: baseURL, UserAgent: userAgent}
}
