package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestPeopleService_List(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	people, resp, err := c.People.List(ctx, companyId, nil)

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

func TestPeopleService_List_paginationLimit(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &PeopleQueryParams{Limit: 100}
	people, resp, err := c.People.List(ctx, companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if got := resp.NextPage; got == 0 {
		t.Errorf("PeopleService Response NextPage is %v, want %v", got, "not 0")
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) != 100 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}
}

func TestPeopleService_List_specificPeople(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	queryparams := &PeopleQueryParams{FirstName: "John"}
	people, resp, err := c.People.List(ctx, companyId, queryparams)

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
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &PeopleQueryParams{Includes: []string{"location"}}
	people, resp, err := c.People.List(ctx, companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if got, want := people[0].Location.RefObject, ""; got != want {
		t.Errorf("PeopleService list is %v, want %v", got, want)
	}
}

func TestPeopleService_List_expandMultiple(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &PeopleQueryParams{
		FirstName: "lucas",
		Includes:  []string{"employments", "manager.department"}}
	people, resp, err := c.People.List(ctx, companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if err != nil {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if len(people) == 0 {
		t.Errorf("PeopleService list is %v, want %v", len(people), err)
	}

	if got, want := people[0].Location.RefObject, "/core/locations"; got != want {
		t.Errorf("PeopleService list is %v, want %v", got, want)
	}

	if got, want := people[0].Employments.RefObject, ""; got != want {
		t.Errorf("people[0].EmploymentsRef.RefObject is %v, want %v", got, want)
	}

	if got, want := people[0].Manager.Department.RefObject, ""; got != want {
		t.Errorf("people[0].Manager.Department is %v, want %v", got, want)
	}
}
