package core

import (
	"fmt"
	"net/url"

	"github.com/microcosm-cc/bluemonday"

	"github.com/google/uuid"
)

func GenerateShortUrl(link string) (string, string, error) {
	// 1. Validate the url
	_, validateErr := validateUrl(link)
	if validateErr != nil {
		return "", "", fmt.Errorf("failed to validate url: %w", validateErr)
	}

	// 2. Sanitize the url
	policy := bluemonday.UGCPolicy()
	sanitizedUrl := policy.Sanitize(link)

	// Generate a short url using a randomId of length 6
	id := uuid.New().String()[:8]

	shortUrl := fmt.Sprintf("https://shortit.sh/%s", id)

	return shortUrl, sanitizedUrl, nil
}

func validateUrl(link string) (*url.URL, error) {

	matched, err := url.ParseRequestURI(link)

	return matched, err
}

func sanitizeUrl(url string) (string, error) {
	// Implement URL sanitization logic here
	return "", nil
}
