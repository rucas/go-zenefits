package zenefits

import (
	"net/http"
	"time"
)

type MeService service

type Me struct {
	Scopes      []string  `json:"scopes"`
	Object      string    `json:"object"`
	Url         string    `json:"url"`
	Company     Ref       `json:"company"`
	Expires     time.Time `json:"expires"`
	Person      Ref       `json:"person"`
	Uninstalled bool      `json:"uninstalled"`
}

type MeQueryParams struct {
	Expansion
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
