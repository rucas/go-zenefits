package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestPeopleService_List(t *testing.T) {
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

	// TODO: RefObject is not null on non expansions
	if got, want := people[0].Location.RefObject, "/core/locations"; got != want {
		t.Errorf("PeopleService list is %v, want %v", got, want)
	}
}

func TestPeopleService_List_specificPeople(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	queryparams := &PeopleQueryParams{FirstName: "John"}
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

func TestPeopleService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	queryparams := &PeopleQueryParams{Includes: []string{"location"}}

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

	// TODO: Should perform this check for all expands
	if got, want := people[0].Location.RefObject, ""; got != want {
		t.Errorf("PeopleService list is %v, want %v", got, want)
	}
}
