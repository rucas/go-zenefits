package zenefits

import (
	"context"
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
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,space,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

func (s *MeService) Get(ctx context.Context, opt *MeQueryParams) (*Me, *Response, error) {
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
	resp, err := s.client.Do(ctx, req, &b)

	if err != nil {
		return nil, resp, err
	}

	return me, resp, nil
}
