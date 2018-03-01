package zenefits

import (
	"fmt"
	"net/http"
)

type EmployeeBanksService service

type EmployeeBanks struct {
	RoutingNumber string `json:"routing_number"`
	BankName      string `json:"bank_name"`
	AccountType   string `json:"account_type"`
	Url           string `json:"url"`
	Object        string `json:"object"`
	Person        Ref    `json:"person"`
	AccountNumber string `json:"account_number"`
	IsVerified    bool   `json:"is_verified"`
	Id            string `json:"id"`
}

type EmployeeBanksFilters struct {
	Person int `url:"person,omitempty"`
}

type EmployeeBanksQueryParams struct {
	EmployeeBanksFilters
	Expansion
}

// TODO: http://api.zenefits.com/core/banks/{:bank_id}
// TODO: http://api.zenefits.com/core/banks

func (s *EmployeeBanksService) List(personId int, opt *EmployeeBanksQueryParams) ([]*EmployeeBanks, *http.Response, error) {
	u := fmt.Sprintf("core/people/%d/banks", personId)
	u, err := AddOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	var employeeBanks []*EmployeeBanks

	resp, err := s.client.Do(req, &employeeBanks)

	if err != nil {
		return nil, resp, err
	}

	return employeeBanks, resp, nil
}
