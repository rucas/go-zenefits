package zenefits

import (
	"fmt"
	"net/http"
)

// NOTE: These will need to be parsed differently since they arent paginated
// TODO: http://api.zenefits.com/core/locations/{:location_id}
// TODO: http://api.zenefits.com/core/locations

type LocationsService service

type Locations struct {
	City    string `json:"city"`
	Name    string `json:"name"`
	People  Ref    `json:"people"`
	Url     string `json:"url"`
	Street1 string `json:"street1"`
	Street2 string `json:"street2"`
	Object  string `json:"object"`
	Id      string `json:"id"`
	Phone   string `json:"phone"`
	State   string `json:"state"`
	Country string `json:"country"`
	Company Ref    `json:"company"`
}

type LocationsFilters struct {
	Company int    `url:"company,omitempty"`
	Name    string `url:"name,omitempty"`
}

type LocationQueryParams struct {
	LocationsFilters
	Expansion
}

func (s *LocationsService) List(companyId int, opt *LocationQueryParams) ([]*Locations, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/locations", companyId)
	u, err := AddOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}
	var locations []*Locations
	resp, err := s.client.Do(req, &locations)

	if err != nil {
		return nil, resp, err
	}

	return locations, resp, nil
}
