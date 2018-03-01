package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestEmploymentsService_ListAll(t *testing.T) {
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
}

func TestEmploymentsService_List_specificEmployments(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{
		EmploymentsFilters{1851863},
		Expansion{},
	}

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

// TODO: update this test to make sure company is expanded
func TestEmploymentsService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &EmploymentQueryParams{
		EmploymentsFilters{},
		Expansion{[]string{"person"}},
	}

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
