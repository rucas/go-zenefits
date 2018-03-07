package zenefits

type Pagination struct {
	EndingBefore  int `url:"ending_before,omitempty"`
	Limit         int `url:"limit,omitempty"`
	StartingAfter int `url:"starting_after,omitempty"`
}

type MetaRef struct {
	Object    string `json:"object"`
	RefObject string `json:"ref_object"`
	Url       string `json:"url"`
}

type MetaList struct {
	Data    interface{} `json:"data"`
	NextUrl string      `json:"next_url"`
	Object  string      `json:"object"`
	Url     string      `json:"url"`
}

type MetaResponse struct {
	Object string   `json:"object"`
	Page   MetaList `json:"data"`
	Status int      `json:"status"`
	// TODO: Error
}

type CompanyRef struct {
	Companies
	MetaRef
}

func addMeta(v interface{}) MetaResponse {
	return MetaResponse{
		Page: MetaList{Data: v},
	}
}
