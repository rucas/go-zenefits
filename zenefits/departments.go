package zenefits

import (
	"fmt"
	"net/http"
)

type DepartmentsService service

type Departments struct {
	Name    string `json:"name"`
	People  Ref    `json:"people"`
	Url     string `json:"url"`
	Company Ref    `json:"company"`
	Object  string `json:"object"`
	Id      string `json:"id"`
}

type DepartmentsFilters struct {
	Company int    `url:"company,omitempty"`
	name    string `url:"name,omitempty"`
}

type DepartmentQueryParams struct {
	DepartmentsFilters
	Expansion
}

func (s *DepartmentsService) List(companyId int, opt *DepartmentQueryParams) ([]*Departments, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/departments", companyId)
	u, err := AddOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	var departments []*Departments

	resp, err := s.client.Do(req, &departments)

	if err != nil {
		return nil, resp, err
	}

	return departments, resp, nil
}
