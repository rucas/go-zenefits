package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestEmploymentsService_List(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	employments, resp, err := c.Employments.List(1851863, nil)

	if resp.StatusCode != 200 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if err != nil {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if len(employments) == 0 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if got, want := employments[0].Person.Object, "/meta/ref/detail"; got != want {
		t.Errorf("EmploymentService list is %v, want %v", got, want)
	}
}

func TestEmploymentsService_List_specificEmployments(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{1851863, nil}

	employments, resp, err := c.Employments.List(1851863, qs)

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
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{Includes: []string{"person"}}

	employments, resp, err := c.Employments.List(1851863, qs)

	if resp.StatusCode != 200 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if err != nil {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if len(employments) == 0 {
		t.Errorf("EmploymentService list is %v, want %v", len(employments), err)
	}

	if got, want := employments[0].Person.Object, "/core/people"; got != want {
		t.Errorf("EmploymentService list is %v, want %v", got, want)
	}
}
