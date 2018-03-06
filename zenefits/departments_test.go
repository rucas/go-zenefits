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

	if got, want := departments[0].Company.Object, "/meta/ref/detail"; got != want {
		t.Errorf("DepartmentsService list is %v, want %v", got, want)
	}
}

func TestDepartmentsService_List_specificDepartments(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &DepartmentQueryParams{Name: "foobarbuzz"}

	departments, resp, err := c.Departments.List(companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if err != nil {
		t.Errorf("DepartmentService list is %v, want %v", len(departments), err)
	}

	if got, want := len(departments), 0; got != want {
		t.Errorf("DepartmentService list is %v, want %v", got, "> 0")
	}
}

func TestDepartmentsService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &DepartmentQueryParams{Includes: []string{"company"}}
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

	if got, want := departments[0].Company.Object, "/core/companies"; got != want {
		t.Errorf("DepartmentsService list is %v, want %v", got, want)
	}
}
