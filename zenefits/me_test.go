package zenefits

import (
	"context"
	"testing"

	"golang.org/x/oauth2"
)

func TestMeService_Get(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	me, resp, err := c.Me.Get(ctx, nil)

	if resp.StatusCode != 200 {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if err != nil {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if me == nil {
		t.Errorf("MeService me is %v, want %v", me, "NOT NIL")
	}

	if got, want := me.Company.RefObject, "/core/companies"; got != want {
		t.Errorf("MeService list is %v, want %v", got, want)
	}
}

func TestMeService_Get_expand(t *testing.T) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: accessToken})
	tc := oauth2.NewClient(ctx, ts)
	c := NewClient(tc)

	qs := &MeQueryParams{Includes: []string{"company"}}
	me, resp, err := c.Me.Get(ctx, qs)

	if resp.StatusCode != 200 {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if err != nil {
		t.Errorf("MeService list is %v, want %v", me, err)
	}

	if me == nil {
		t.Errorf("MeService me is %v, want %v", me, "NOT NIL")
	}

	if got, want := me.Company.RefObject, ""; got != want {
		t.Errorf("MeService list is %v, want %v", got, want)
	}
}
