package zenefits

import (
	"net/http"
	"time"
)

type MeService service

//TODO: RefObject string    `json:"ref_object"`

type Me struct {
	Company     Companies `json:"company"`
	Expires     time.Time `json:"expires"`
	Object      string    `json:"object"`
	Person      People    `json:"person"`
	Scopes      []string  `json:"scopes"`
	Uninstalled bool      `json:"uninstalled"`
	Url         string    `json:"url"`
}

type MeQueryParams struct {
	Includes []string `url:"includes,omitempty"`
}

func (s *MeService) Get(opt *MeQueryParams) (*Me, *http.Response, error) {
	u, err := addOptions("core/me", opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var me *Me

	type response struct {
		Status int         `json:"status"`
		Object string      `json:"object"`
		Data   interface{} `json:"data"`
		Error  string      `json:"error"`
	}

	b := response{Data: &me}
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return me, resp, nil
}
