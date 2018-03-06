package zenefits

import (
	"fmt"
	"net/http"
)

type PeopleService service

type People struct {
	Object               string         `json:"object"`
	RefObject            string         `json:"ref_object"`
	Lastname             string         `json:"last_name"`
	PreferredName        string         `json:"preferred_name"`
	Manager              *People        `json:"manager"`
	PostalCode           string         `json:"postal_code"`
	Id                   string         `json:"id"`
	City                 string         `json:"city"`
	FirstName            string         `json:"first_name"`
	MiddleName           string         `json:"middle_name"`
	Title                string         `json:"title"`
	WorkPhone            string         `json:"work_phone"`
	PersonalEmail        string         `json:"personal_email"`
	State                string         `json:"state"`
	DateOfBirth          string         `json:"date_of_birth"`
	Location             Locations      `json:"location"`
	Subordinates         MetaRef        `json:"subordinates"`
	Department           Departments    `json:"department"`
	Employments          EmploymentsRef `json:"employments"`
	Type                 string         `json:"type"`
	Company              Companies      `json:"company"`
	Status               string         `json:"status"`
	Street1              string         `json:"street1"`
	Street2              string         `json:"street2"`
	PersonalPhone        string         `json:"personal_phone"`
	SocialSecurityNumber string         `json:"social_security_number"`
	FederalFilingStatus  string         `json:"federal_filing_status"`
	WorkEmail            string         `json:"work_email"`
	Url                  string         `json:"url"`
	Country              string         `json:"country"`
	Gender               string         `json:"gender"`
	Banks                BanksRef       `json:"banks"`
}

type EmploymentsRef struct {
	Data []Employments `json:"data"`
	MetaRef
}

type BanksRef struct {
	Data []EmployeeBanks `json:"data"`
	MetaRef
}

type PeopleFilters struct {
	Company    int    `url:"company,omitempty"`
	Department int    `url:"department,omitempty"`
	FirstName  string `url:"first_name,omitempty"`
	LastName   string `url:"lastname,omitempty"`
	Status     string `url:"status,omitempty"`
	Location   int    `url:"location,omitempty"`
	Manager    int    `url:"manager,omitempty"`
}

type PeopleQueryParams struct {
	Company    int      `url:"company,omitempty"`
	Department int      `url:"department,omitempty"`
	FirstName  string   `url:"first_name,omitempty"`
	LastName   string   `url:"lastname,omitempty"`
	Status     string   `url:"status,omitempty"`
	Location   int      `url:"location,omitempty"`
	Manager    int      `url:"manager,omitempty"`
	Includes   []string `url:"includes,omitempty"`
}

// TODO: GET http://api.zenefits.com/core/people/{:id}

// TODO: ListAll
// Once zenefits changes their api to one token per multiple companies
// http://api.zenefits.com/core/people

func (s *PeopleService) List(companyId int, opt *PeopleQueryParams) ([]*People, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/people", companyId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var people []*People
	b := addPaginationBody(&people)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return people, resp, nil
}

/*func (s *PeopleService) Get(personId int, opt *PeopleQueryParams) (*People, *http.Response, error) {
	u := fmt.Sprintf("core/people/%d", personId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var people *People
	resp, err := s.client.Do(req, &people)

	if err != nil {
		return nil, resp, err
	}

	return people, resp, nil
}*/
