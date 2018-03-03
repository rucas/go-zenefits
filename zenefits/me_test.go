package zenefits

import (
	"testing"

	"golang.org/x/oauth2"
)

func TestMeService_Get(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	me, resp, err := c.Me.Get(nil)

	if resp.StatusCode != 200 {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if err != nil {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if me == nil {
		t.Errorf("MeService me is %v, want %v", me, "NOT NIL")
	}
}

// TODO: update this test to make sure banks and location is expanded
func TestMeService_Get_expand(t *testing.T) {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(nil, ts)
	c := NewClient(tc)

	qs := &MeQueryParams{
		Expansion{[]string{"company"}},
	}

	me, resp, err := c.Me.Get(qs)

	if resp.StatusCode != 200 {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if err != nil {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if me == nil {
		t.Errorf("MeService me is %v, want %v", me, "NOT NIL")
	}
}
