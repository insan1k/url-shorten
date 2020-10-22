package shorted

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"strings"
	"time"
)

// ShortURL struct holds all the data required for redirecting a request
type ShortURL struct {
	ShortID   ID
	Original  urlShorten
	Short     urlShorten
	timestamp time.Time
}

// NewURL creates ShortURL from an original url
func NewURL(u string) (s ShortURL, err error) {
	s.Original, err = newOriginal(u)
	if err != nil {
		return
	}
	s.ShortID, err = NewID()
	if err != nil {
		return
	}
	s.Short, err = newShort(s.ShortID.encodeID())
	s.timestamp = time.Now()
	return
}

// NewPartialShortURLFromAPI creates ShortURL from an API request
func NewPartialShortURLFromAPI(ss string) (s ShortURL, err error) {
	err = s.ShortID.decodeID(strings.ReplaceAll(ss, configuration.URLShortenerPath, ""))
	return
}

// NewShortURLFromModel creates ShortURL from model.ShortURL
func NewShortURLFromModel(m model.ShortURL) (s ShortURL, err error) {
	err = s.ShortID.decodeID(m.ShortID)
	if err != nil {
		return
	}
	original, err := newOriginal(m.Original)
	if err != nil {
		return
	}
	s.Original = original
	short, err := parseShort(m.Short)
	if err != nil {
		return
	}
	s.Short = short
	return
}

// ToModel converts ShortURL to model.ShortURL
func (s ShortURL) ToModel() model.ShortURL {
	return model.ShortURL{
		ShortID:   s.ShortID.encodeID(),
		Original:  s.Original.string(),
		Short:     s.Short.string(),
		Timestamp: s.timestamp,
	}
}
