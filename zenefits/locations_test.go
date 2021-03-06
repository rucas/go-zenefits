package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestLocationService_List(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	locations, resp, err := c.Locations.List(ctx, companyId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if err != nil {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	// TODO: change this to got, want
	if len(locations) == 0 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if got, want := locations[0].Company.RefObject, "/core/companies"; got != want {
		t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
	}
}

func TestLocationService_List_specific(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &LocationQueryParams{Company: 4567}
	locations, resp, err := c.Locations.List(ctx, companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if err != nil {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if len(locations) == 0 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}
}

func TestLocationService_List_expand(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &LocationQueryParams{Includes: []string{"company"}}
	locations, resp, err := c.Locations.List(ctx, companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if err != nil {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if len(locations) == 0 {
		t.Errorf("LocationService list is %v, want %v", len(locations), err)
	}

	if got, want := locations[0].Company.RefObject, ""; got != want {
		t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
	}
}
