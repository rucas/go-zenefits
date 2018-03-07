package zenefits

import (
	"fmt"
)

type DepartmentsService service

type Departments struct {
	Company   Companies `json:"company"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Object    string    `json:"object"`
	People    MetaRef   `json:"people"`
	RefObject string    `json:"ref_object"`
	Url       string    `json:"url"`
}

type DepartmentQueryParams struct {
	Company       int      `url:"company,omitempty"`
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,space,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Name          string   `url:"name,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

// TODO: GET http://api.zenefits.com/core/departments/{:department_id}

// TODO: ListAll
// Once zenefits changes their api to one token per multiple companies
// http://api.zenefits.com/core/departments

func (s *DepartmentsService) List(companyId int, opt *DepartmentQueryParams) ([]*Departments, *Response, error) {
	u := fmt.Sprintf("core/companies/%d/departments", companyId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	var departments []*Departments
	b := addMeta(&departments)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return departments, resp, nil
}
