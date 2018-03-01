package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestLocationService_ListAll(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	locations, resp, err := c.Locations.List(companyId, nil)

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
}

func TestLocationService_List_specific(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &LocationQueryParams{
		LocationsFilters{Company: 4567},
		Expansion{},
	}

	locations, resp, err := c.Locations.List(companyId, qs)
	//fmt.Printf("sup %#v", locations)

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
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &LocationQueryParams{
		LocationsFilters{},
		Expansion{[]string{"company"}},
	}

	locations, resp, err := c.Locations.List(companyId, qs)

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
