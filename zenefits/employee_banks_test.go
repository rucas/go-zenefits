package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestEmployeeBanksService_List(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	employeeBanks, resp, err := c.EmployeeBanks.List(1851863, nil)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if len(employeeBanks) == 0 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if got, want := employeeBanks[0].Person.RefObject, "/core/people"; got != want {
		t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
	}
}

func TestEmployeeBanksService_List_specificEmployee(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmployeeBanksQueryParams{1851863, nil}

	employeeBanks, resp, err := c.EmployeeBanks.List(1851863, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if len(employeeBanks) == 0 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}
}

func TestEmployeeBanksService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmployeeBanksQueryParams{Includes: []string{"person"}}
	employeeBanks, resp, err := c.EmployeeBanks.List(1851863, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if err != nil {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if len(employeeBanks) == 0 {
		t.Errorf("EmployeeBankservice list is %v, want %v", len(employeeBanks), err)
	}

	if got, want := employeeBanks[0].Person.RefObject, ""; got != want {
		t.Errorf("EmployeeBankservice list is %v, want %v", got, want)
	}
}
