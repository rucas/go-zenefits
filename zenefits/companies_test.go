package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestCompaniesService_List(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	companies, resp, err := c.Companies.List(nil)

	if resp.StatusCode != 200 {
		t.Errorf("CompaniesService list is %v, want %v",
			len(companies), err)
	}

	if err != nil {
		t.Errorf("CompaniesService list is %v, want %v",
			len(companies), err)
	}

	if len(companies) == 0 {
		t.Errorf("CompaniesService list is %v, want %v",
			len(companies), err)
	}

	if got, want := companies[0].Departments.Object, "/meta/ref/list"; got != want {
		t.Errorf("CompaniesService list is %v, want %v", got, want)
	}
}

func TestCompaniesService_List_specificCompanies(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &CompaniesQueryParams{Name: "lucas"}

	companies, resp, err := c.Companies.List(qs)

	if resp.StatusCode != 200 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if err != nil {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if len(companies) != 0 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}
}

func TestCompaniesService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	q := &CompaniesQueryParams{Includes: []string{"departments"}}
	companies, resp, err := c.Companies.List(q)

	if resp.StatusCode != 200 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if err != nil {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if len(companies) == 0 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if got, want := companies[0].Departments.Object, "/meta/list"; got != want {
		t.Errorf("CompaniesService list is %v, want %v", got, want)
	}
}
