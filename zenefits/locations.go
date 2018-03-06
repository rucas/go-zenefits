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
	City      string    `json:"city"`
	Company   Companies `json:"company"`
	Country   string    `json:"country"`
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Object    string    `json:"object"`
	People    MetaRef   `json:"people"`
	Phone     string    `json:"phone"`
	RefObject string    `json:"ref_object"`
	State     string    `json:"state"`
	Street1   string    `json:"street1"`
	Street2   string    `json:"street2"`
	Url       string    `json:"url"`
}

type LocationQueryParams struct {
	Company  int      `url:"company,omitempty"`
	Includes []string `url:"includes,omitempty"`
	Name     string   `url:"name,omitempty"`
}

// TODO: http://api.zenefits.com/core/locations/{:location_id}
// TODO: http://api.zenefits.com/core/locations

func (s *LocationsService) List(companyId int, opt *LocationQueryParams) ([]*Locations, *http.Response, error) {
	u := fmt.Sprintf("core/companies/%d/locations", companyId)
	u, err := addOptions(u, opt)
	req, err := s.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, nil, err
	}

	var locations []*Locations
	b := addPaginationBody(&locations)
	resp, err := s.client.Do(req, &b)

	if err != nil {
		return nil, resp, err
	}

	return locations, resp, nil
}
