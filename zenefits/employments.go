package zenefits

import (
	"fmt"
	"net/http"
)

type EmploymentsService service

type Employments struct {
	Object              string `json:"object"`
	Person              Ref    `json:"person"`
	Url                 string `json:"url"`
	CompType            string `json:"comp_type"`
	EmploymentType      string `json:"employment_type"`
	PayRate             string `json:"pay_rate"` // TODO: not sure if this is null all the time
	TerminationType     string `json:"termination_type"`
	AnnualSalary        string `json:"annual_salary"`
	TerminationDate     string `json:"termination_date"`
	HireDate            string `json:"hire_date"`
	Id                  string `json:"id"`
	WorkingHoursPerWeek string `json:"working_hours_per_week"`
}

type EmploymentsFilters struct {
	Person int `url:"person,omitempty"`
}

type EmploymentQueryParams struct {
	EmploymentsFilters
	Expansion
}

// TODO: https://api.zenefits.com/core/employments/{:employment_id}

func (s *EmploymentsService) List(personId int, opt *EmploymentQueryParams) ([]*Employments, *http.Response, error) {
	u := fmt.Sprintf("core/people/%d/employments", personId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var employments []*Employments
	b := addPaginationBody(&employments)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return employments, resp, nil
}
