package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestDepartmentsService_List(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	departments, resp, err := c.Departments.List(companyId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if err != nil {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if len(departments) == 0 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}
}

func TestDepartmentsService_List_specificDepartments(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &DepartmentQueryParams{
		DepartmentsFilters{0, "rucas"},
		Expansion{},
	}

	departments, resp, err := c.Departments.List(companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if err != nil {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	// TODO: change this to got, want
	if len(departments) == 0 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}
}

// TODO: update this test to make sure company is expanded
func TestDepartmentsService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &DepartmentQueryParams{
		DepartmentsFilters{},
		Expansion{[]string{"company"}},
	}

	departments, resp, err := c.Departments.List(companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if err != nil {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if len(departments) == 0 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}
}
