package zenefits

import (
	"os"
	"strconv"
	"testing"

	"golang.org/x/oauth2"
)

var (
	accessToken  = os.Getenv("ZENEFITS_API_KEY")
	companyId, _ = strconv.Atoi(os.Getenv("ZENEFITS_COMPANY_ID"))
)

func TestPeopleService_ListAll(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	people, resp, err := c.People.List(companyId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}
}

func TestPeopleService_List_specificPeople(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	f := PeopleFilters{FirstName: "John"}
	queryparams := &PeopleQueryParams{PeopleFilters: f}
	people, resp, err := c.People.List(companyId, queryparams)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}
}

// TODO: update this test to make sure banks and location is expanded
func TestPeopleService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	e := Expansion{[]string{"banks", "location"}}
	queryparams := &PeopleQueryParams{Expansion: e}

	people, resp, err := c.People.List(companyId, queryparams)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}
}

func TestPeopleService_Get(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	people, resp, err := c.People.Get(166216, nil)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", people, err)
	}

	/*if err != nil {
		t.Errorf("PeopleService list is %v, want %v", people, err)
	}*/

	/*if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}*/
}
