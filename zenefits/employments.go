package zenefits

import (
	"fmt"
)

type EmploymentsService service

type Employments struct {
	AnnualSalary        string `json:"annual_salary"`
	CompType            string `json:"comp_type"`
	EmploymentType      string `json:"employment_type"`
	HireDate            string `json:"hire_date"`
	Id                  string `json:"id"`
	Object              string `json:"object"`
	PayRate             string `json:"pay_rate"` // TODO: not sure if this is null all the time
	Person              People `json:"person"`
	RefObject           string `json:"ref_object"`
	TerminationDate     string `json:"termination_date"`
	TerminationType     string `json:"termination_type"`
	Url                 string `json:"url"`
	WorkingHoursPerWeek string `json:"working_hours_per_week"`
}

type EmploymentQueryParams struct {
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Person        int      `url:"person,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

// TODO: https://api.zenefits.com/core/employments/{:employment_id}

func (s *EmploymentsService) List(personId int, opt *EmploymentQueryParams) ([]*Employments, *Response, error) {
	u := fmt.Sprintf("core/people/%d/employments", personId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var employments []*Employments
	b := addMeta(&employments)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return employments, resp, nil
}
