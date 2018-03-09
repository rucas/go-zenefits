package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestEmployeeBanksService_List(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	employeeBanks, resp, err := c.EmployeeBanks.List(ctx, personId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	/*
		if len(employeeBanks) == 0 {
			t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
		}

		if got, want := employeeBanks[0].Person.RefObject, "/core/people"; got != want {
			t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
		}
	*/
}

func TestEmployeeBanksService_List_specificEmployee(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &EmployeeBanksQueryParams{Person: personId}
	employeeBanks, resp, err := c.EmployeeBanks.List(ctx, personId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	/*
		if len(employeeBanks) == 0 {
			t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
		}
	*/
}

func TestEmployeeBanksService_List_expand(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &EmployeeBanksQueryParams{Includes: []string{"person"}}
	employeeBanks, resp, err := c.EmployeeBanks.List(1851863, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	/*
		if len(employeeBanks) == 0 {
			t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
		}

		if got, want := employeeBanks[0].Person.RefObject, ""; got != want {
			t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
		}
	*/
}
