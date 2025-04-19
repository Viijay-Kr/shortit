package core

import (
	"regexp"
	"testing"
)

// TestGenerateShortUrl tests the GenerateShortUrl function from core/core.go
func TestGenerateShortUrl(t *testing.T) {
	url := "https://google.com?<script>alert('xss')</script>query=1&res=2"
	expected := "https://shortit.sh"

	short_url, _, err := GenerateShortUrl(url)

	want := regexp.MustCompile(expected)

	if !want.MatchString(short_url) || err != nil {
		t.Errorf(`GenerateShortUrl(%q) = %q,%v, want match for %#q,nil`, url, expected, err, want)
	}
}
