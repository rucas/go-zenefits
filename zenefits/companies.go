package zenefits

import "context"

type CompaniesService service

type Companies struct {
	CompanyBanks CompanyBanksRef `json:"company_banks"`
	Departments  DepartmentsRef  `json:"departments"`
	Ein          string          `json:"ein"`
	Id           string          `json:"id"`
	LegalCity    string          `json:"legal_city"`
	LegalName    string          `json:"legal_name"`
	LegalState   string          `json:"legal_state"`
	LegalStreet1 string          `json:"legal_street1"`
	LegalStreet2 string          `json:"legal_street2"`
	LegalZip     string          `json:"legal_zip"`
	Locations    LocationsRef    `json:"locations"`
	LogoUrl      string          `json:"logo_url"`
	Name         string          `json:"name"`
	Object       string          `json:"object"`
	People       MetaRef         `json:"people"`
	RefObject    string          `json:"ref_object"`
	Url          string          `json:"url"`
}

type CompanyBanksRef struct {
	Data []CompanyBanks `json:"data"`
	MetaRef
}

type DepartmentsRef struct {
	Data []Locations `json:"data"`
	MetaRef
}

type LocationsRef struct {
	Data []Departments `json:"data"`
	MetaRef
}

// TODO: looks like it can only be one include
type CompaniesQueryParams struct {
	EndingBefore  int      `url:"ending_before,omitempty"`
	Includes      []string `url:"includes,space,omitempty"`
	Limit         int      `url:"limit,omitempty"`
	Name          string   `url:"name,omitempty"`
	StartingAfter int      `url:"starting_after,omitempty"`
}

// TODO: GET http://api.zenefits.com/core/companies/{:id}

func (s *CompaniesService) List(ctx context.Context, opt *CompaniesQueryParams) ([]*Companies, *Response, error) {
	u, err := addOptions("core/companies", opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var companies []*Companies
	b := addMeta(&companies)
	resp, err := s.client.Do(ctx, req, &b)

	if err != nil {
		return nil, resp, err
	}

	return companies, resp, nil
}
