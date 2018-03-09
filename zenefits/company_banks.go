package zenefits

import (
	"context"
	"fmt"
)

type CompanyBanksService service

type CompanyBanks struct {
	AccountNumber string     `json:"account_number"`
	BankName      string     `json:"bank_name"`
	Company       CompanyRef `json:"company"`
	Id            string     `json:"id"`
	RefObject     string     `json:"ref_object"`
	RoutingNumber string     `json:"routing_number"`
}

type CompanyBanksQueryParams struct {
	Company       int      `url:"company,omitempty"`
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,space,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

// TODO: GET http://api.zenefits.com/core/company_banks/{:bank_id}

// The following endpoint gives all the information for all banks across
// all companies
// (note access tokens are unique per company, so this will be the same as the main endpoint)
// TODO: http://api.zenefits.com/core/company_banks

func (s *CompanyBanksService) List(ctx context.Context, companyId int, opt *CompanyBanksQueryParams) ([]*CompanyBanks, *Response, error) {
	u := fmt.Sprintf("core/companies/%d/company_banks", companyId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var banks []*CompanyBanks
	b := addMeta(&banks)
	resp, err := s.client.Do(ctx, req, &b)

	if err != nil {
		return nil, resp, err
	}
	return banks, resp, nil
}
