package zenefits

import (
	"fmt"
	"net/http"
)

type CompanyBanksService service

type CompanyBanks struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	Company       Ref    `json:"company"`
	Id            string `json:"id"`
	RoutingNumber string `json:"routing_number"`
}

type CompanyBanksFilters struct {
	Company int `url:"company,omitempty"`
}

type CompanyBanksQueryParams struct {
	CompanyBanksFilters
	Expansion
}

// TODO: GET http://api.zenefits.com/core/company_banks/{:bank_id}

// The following endpoint gives all the information for all banks across
// all companies
// (note access tokens are unique per company, so this will be the same as the main endpoint)
// TODO: http://api.zenefits.com/core/company_banks

func (s *CompanyBanksService) List(companyId int, opt *CompanyBanksQueryParams) ([]*CompanyBanks, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/company_banks", companyId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var banks []*CompanyBanks
	b := addPaginationBody(&banks)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}
	return banks, resp, nil
}
