package shorted

import (
	"errors"
	"fmt"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"net/url"
)

const (
	errorInvalidScheme           = "short url schema is not supported"
	errorInvalidHostname         = "short url hostname is not supported"
	errorInvalidPort             = "short url port is not supported"
	errorInvalidRedirectScheme   = "original url scheme %v is not supported"
	errorInvalidRedirectHostname = "cannot redirect to this server"
)

// urlShorten defines how we're going to work with shortened urls
type urlShorten url.URL

func newOriginal(o string) (u urlShorten, err error) {
	// - original must be a valid url
	// - original has to be point to a host != than this server
	parsedURL, err := url.Parse(o)
	if err != nil {
		return
	}
	if err = u.validateScheme(parsedURL.Scheme); err != nil {
		return
	}
	if parsedURL.Hostname() == configuration.C.ShortenDomain {
		err = errors.New(errorInvalidRedirectHostname)
		return
	}
	u = urlShorten(*parsedURL)
	return
}

func newShort(id string) (u urlShorten, err error) {
	url := configuration.C.ShortenScheme + "://" + configuration.C.ShortenDomain
	if configuration.C.ShortenPort != "" {
		url += ":" + configuration.C.ShortenPort
	}
	url += configuration.URLShortenerPath + id
	return parseShort(url)
}

func parseShort(s string) (u urlShorten, err error) {
	// - short must be a valid url
	// - short has to point to a host in this server (eg.: configuration changed, pulled crappy data from db, ????)
	parsedURL, err := url.Parse(s)
	if err != nil {
		return
	}
	if parsedURL.Scheme != configuration.C.ShortenScheme {
		err = errors.New(errorInvalidScheme)
		return
	}
	if parsedURL.Hostname() != configuration.C.ShortenDomain {
		err = errors.New(errorInvalidHostname)
		return
	}
	if parsedURL.Port() != configuration.C.ShortenPort {
		err = errors.New(errorInvalidPort)
		return
	}
	u = urlShorten(*parsedURL)
	return
}

func (u urlShorten) escapeEndpoint() (escaped string) {
	escaped = u.Path[len(configuration.URLShortenerPath):]
	return
}

func (u urlShorten) validateScheme(scheme string) (err error) {
	//check if it overlaps once
	var configSchemes = []string{"https", "http"}
	for _, as := range configSchemes {
		if as == scheme {
			return
		}
	}
	err = fmt.Errorf(errorInvalidRedirectScheme, scheme)
	return
}

// returns the string form of the url
func (u urlShorten) string() string {
	l := url.URL(u)
	return l.String()
}
