package zenefits

import (
	"context"
	"fmt"
)

type EmployeeBanksService service

type EmployeeBanks struct {
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	BankName      string `json:"bank_name"`
	Id            string `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	Object        string `json:"object"`
	Person        People `json:"person"`
	RefObject     string `json:"ref_object"`
	RoutingNumber string `json:"routing_number"`
	Url           string `json:"url"`
}

type EmployeeBanksQueryParams struct {
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,space,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Person        int      `url:"person,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

// TODO: GET http://api.zenefits.com/core/banks/{:bank_id}

// The following endpoint gives all the information for all banks across all people
// TODO: ListALL http://api.zenefits.com/core/banks

func (s *EmployeeBanksService) List(ctx context.Context, personId int, opt *EmployeeBanksQueryParams) ([]*EmployeeBanks, *Response, error) {
	u := fmt.Sprintf("core/people/%d/banks", personId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var banks []*EmployeeBanks
	b := addMeta(&banks)
	resp, err := s.client.Do(ctx, req, &b)

	if err != nil {
		return nil, resp, err
	}

	return banks, resp, nil
}
