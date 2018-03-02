// Package  provides ...
package zenefits

import (
	"io"
	"reflect"
	"testing"
)

func TestZenefits_NewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), baseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}

func TestZenefits_NewRequest(t *testing.T) {
	c := NewClient(nil)
	r, _ := c.NewRequest("GET", "/foo", nil)

	if got, want := r.Method, "GET"; got != want {
		t.Errorf("NewRequest path is %v, want %v", got, want)
	}

	if got, want := r.URL.String(), baseURL+"/foo"; got != want {
		t.Errorf("NewRequest path is %v, want %v", got, want)
	}

	if got, want := r.Header.Get("Content-Type"), "application/json"; got != want {
		t.Errorf("NewRequest header for Content-Type is %v, want %v", got, want)
	}

	if got, want := r.Header.Get("User-Agent"), "go-zenefits"; got != want {
		t.Errorf("NewRequest header for User-Agent is %v, want %v", got, want)
	}

	var buf io.Reader
	if got, want := r.Body, buf; got != want {
		t.Errorf("NewRequest body is %v, want %v", got, want)
	}
}

func TestZenefits_addPaginationBody(t *testing.T) {
	var people []*People
	b := addPaginationBody(people)
	if !reflect.DeepEqual(b.Page.Data, people) {
		t.Errorf("PaginationBody page data body = %v, want = %v", b.Page.Data, people)
	}
}

func TestZenefits_Do(t *testing.T) {
	c := NewClient(nil)
	r, _ := c.NewRequest("GET", "/core/me", nil)
	type foo struct {
		foobar int
	}
	body := &foo{}
	c.Do(r, body)
	want := &foo{}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want = %v", body, want)
	}
}

func TestZenefits_addOptions(t *testing.T) {
	type queryparams struct {
		Real bool   `url:"real,omitempty"`
		Id   int    `url:"id,omitempty"`
		Help string `url:"help,omitempty"`
	}
	q := queryparams{true, 34, "me"}

	got, err := addOptions("https://thelucas.blog", q)
	want := "https://thelucas.blog?help=me&id=34&real=true"

	if err != nil {
		t.Errorf("addOptions got error %v", err)
	}

	if got != want {
		t.Errorf("addOptions is %v, want %v", got, want)
	}
}
