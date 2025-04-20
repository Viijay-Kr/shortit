package core

import (
	"fmt"
	"net/url"

	"github.com/Viijay-Kr/shortit/config"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
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

	// Generate a short url using a randomId of length 8
	id := uuid.New().String()[:8]

	cfg := config.GetConfig()
	shortUrl := fmt.Sprintf("%s/%s", cfg.ShortitRedirectHost, id)

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
