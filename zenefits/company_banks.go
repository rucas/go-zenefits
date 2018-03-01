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

// TODO: http://api.zenefits.com/core/company_banks/{:bank_id}
// TODO: http://api.zenefits.com/core/company_banks

func (s *CompanyBanksService) List(companyId int, opt *CompanyBanksQueryParams) ([]*CompanyBanks, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/company_banks", companyId)
	u, err := AddOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}
	var companyBanks []*CompanyBanks

	resp, err := s.client.Do(req, &companyBanks)

	if err != nil {
		return nil, resp, err
	}
	return companyBanks, resp, nil
}