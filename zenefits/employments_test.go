package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestEmploymentsService_List(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	employments, resp, err := c.Employments.List(ctx, personId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if err != nil {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if len(employments) == 0 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if got, want := employments[0].Person.RefObject, "/core/people"; got != want {
		t.Errorf("EmploymentService list is %v, want %v", got, want)
	}
}

func TestEmploymentsService_List_specificEmployments(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{Person: personId}
	employments, resp, err := c.Employments.List(ctx, personId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if err != nil {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if len(employments) == 0 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}
}

func TestEmploymentsService_List_expand(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{Includes: []string{"person"}}
	employments, resp, err := c.Employments.List(ctx, personId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if err != nil {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if len(employments) == 0 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if got, want := employments[0].Person.RefObject, ""; got != want {
		t.Errorf("EmploymentService list is %v, want %v", got, want)
	}
}
