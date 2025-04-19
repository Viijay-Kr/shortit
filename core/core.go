package core

import (
	"fmt"
	"net/url"

	"github.com/microcosm-cc/bluemonday"

	"github.com/google/uuid"
)

type ShortUrl struct {
	ID        string
	Sanitized string
	Shortened string
}

func GenerateShortUrl(link string) (ShortUrl, error) {
	// 1. Validate the url
	_, validateErr := validateUrl(link)
	if validateErr != nil {
		return ShortUrl{}, fmt.Errorf("failed to validate url: %w", validateErr)
	}

	// 2. Sanitize the url
	policy := bluemonday.UGCPolicy()
	sanitizedUrl := policy.Sanitize(link)

	// Generate a short url using a randomId of length 6
	id := uuid.New().String()[:8]

	shortUrl := fmt.Sprintf("https://shortit.sh/%s", id)

	return ShortUrl{
		ID:        id,
		Sanitized: sanitizedUrl,
		Shortened: shortUrl,
	}, nil
}

func validateUrl(link string) (*url.URL, error) {

	matched, err := url.ParseRequestURI(link)

	return matched, err
}

func sanitizeUrl(url string) (string, error) {
	// Implement URL sanitization logic here
	return "", nil
}
