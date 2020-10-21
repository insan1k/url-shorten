package shorted

import (
	"github.com/insan1k/one-qr-dot-me/internal/model"
)

// ShortURL struct holds all the data required for redirecting a request
type ShortURL struct {
	ID       ID
	Original urlShorten
	Short    urlShorten
}

// NewURL creates ShortURL from an original url
func NewURL(u string) (s ShortURL, err error) {
	s.Original, err = newOriginal(u)
	if err != nil {
		return
	}
	s.ID, err = NewID()
	if err != nil {
		return
	}
	s.Short, err = newShort(s.ID.encodeID())
	return
}

// NewPartialShortURLFromAPI creates ShortURL from an API request
func NewPartialShortURLFromAPI(ss string) (s ShortURL, err error) {
	short, err := parseShort(ss)
	if err != nil {
		return
	}
	err = s.ID.decodeID(short.escapeEndpoint())
	if err != nil {
		return
	}
	s.Short = short
	return
}

// NewShortURLFromModel creates ShortURL from model.ShortURL
func NewShortURLFromModel(m model.ShortURL) (s ShortURL, err error) {
	err = s.ID.decodeID(m.ID)
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
		ID:       s.ID.encodeID(),
		Original: s.Original.string(),
		Short:    s.Short.string(),
	}
}
