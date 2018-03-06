package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestCompanyBanksService_List(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	companyBanks, resp, err := c.CompanyBanks.List(companyId, nil)

	if resp.StatusCode != 200 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if err != nil {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if len(companyBanks) != 0 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}
}

func TestCompanyBanksService_List_specificCompany(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &CompanyBanksQueryParams{Company: 1234}

	companyBanks, resp, err := c.CompanyBanks.List(companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if err != nil {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if len(companyBanks) != 0 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}
}

func TestCompanyBanksService_List_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &CompanyBanksQueryParams{
		Includes: []string{"company"},
	}

	companyBanks, resp, err := c.CompanyBanks.List(companyId, qs)

	if resp.StatusCode != 200 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if err != nil {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}

	if len(companyBanks) != 0 {
		t.Errorf("CompanyBankservice list is %v, want %v", len(companyBanks), err)
	}
}
