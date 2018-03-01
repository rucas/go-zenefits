package zenefits

import (
	"net/http"
)

type CompaniesService service

type Companies struct {
	LegalName    string `json:"legal_name"`
	LegalZip     string `json:"legal_zip"`
	LegalStreet1 string `json:"legal_street1"`
	LegalStreet2 string `json:"legal_street2"`
	People       Ref    `json:"people"`
	CompanyBanks Ref    `json:"company_banks"`
	Object       string `json:"object"`
	Locations    Ref    `json:"locations"`
	Departments  Ref    `json:"departments"`
	LegalCity    string `json:"legal_city"`
	Url          string `json:"url"`
	Ein          string `json:"ein"`
	LogoUrl      string `json:"logo_url"`
	Id           string `json:"id"`
	LegalState   string `json:"legal_state"`
	Name         string `json:"name"`
}

// TODO: does this make more sense to make private and then expose
// companyqueryparams?
type CompaniesFilters struct {
	Name string `url:"name,omitempty"`
}

type CompaniesQueryParams struct {
	CompaniesFilters
	Expansion
}


// TODO: http://api.zenefits.com/core/companies/{:id}
func (s *CompaniesService) List(opt *CompaniesQueryParams) ([]*Companies, *http.Response, error) {
	u, err := AddOptions("core/companies", opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var companies []*Companies
	resp, err := s.client.Do(req, &companies)

	if err != nil {
		return nil, resp, err
	}

	return companies, resp, nil
}
