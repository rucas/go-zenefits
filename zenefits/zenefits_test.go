// Package  provides ...
package zenefits

import "testing"

func TestNewClient(t *testing.T) {
	c := NewClient(nil)

	if got, want := c.BaseURL.String(), baseURL; got != want {
		t.Errorf("NewClient BaseURL is %v, want %v", got, want)
	}

	if got, want := c.UserAgent, userAgent; got != want {
		t.Errorf("NewClient UserAgent is %v, want %v", got, want)
	}
}
