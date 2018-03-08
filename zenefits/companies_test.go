package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestCompaniesService_List(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	companies, resp, err := c.Companies.List(ctx, nil)

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

	if got, want := companies[0].Departments.RefObject, "/core/departments"; got != want {
		t.Errorf("CompaniesService list is %v, want %v", got, want)
	}
}

func TestCompaniesService_List_specificCompanies(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &CompaniesQueryParams{Name: "rucas industries"}

	companies, resp, err := c.Companies.List(ctx, qs)

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
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	q := &CompaniesQueryParams{Includes: []string{"departments"}}
	companies, resp, err := c.Companies.List(ctx, q)

	if resp.StatusCode != 200 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if err != nil {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if len(companies) == 0 {
		t.Errorf("CompaniesService list is %v, want %v", len(companies), err)
	}

	if got, want := companies[0].Departments.RefObject, ""; got != want {
		t.Errorf("CompaniesService list is %v, want %v", got, want)
	}
}
