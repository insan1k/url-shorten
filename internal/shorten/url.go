package shorten

import (
	"errors"
	"fmt"
	"net/url"
)

const(
	errorInvalidScheme = "short url schema is not supported"
	errorInvalidHostname = "short url hostname is not supported"
	errorInvalidPort = "short url port is not supported"
	errorInvalidRedirectScheme = "original url scheme %v is not supported"
	errorInvalidRedirectHostname = "cannot redirect to this server"
)

// urlShorten defines how we're going to work with shortened urls
type urlShorten url.URL


func newOriginal(o string) (u urlShorten,err error) {
	// - original must be a valid url
	// - original has to be point to a host != than this server
	parsedURL,err:=url.Parse(o)
	if err!=nil{
		return
	}
	if err=u.checkScheme(parsedURL.Scheme);err!=nil{
		return
	}
	if parsedURL.Hostname() == configHostname {
		err = errors.New(errorInvalidRedirectHostname)
		return
	}
	u = urlShorten(*parsedURL)
	return
}

func newShort(s string) (u urlShorten, err error) {
	// - short must be a valid url
	// - short has to point to a host in this server (eg.: configuration changed, pulled crappy data from db, ????)
	parsedURL,err:=url.Parse(s)
	if err!=nil{
		return
	}
	if parsedURL.Scheme != configScheme {
		err=errors.New(errorInvalidScheme)
		return
	}
	if parsedURL.Hostname() != configHostname {
		err=errors.New(errorInvalidHostname)
		return
	}
	if parsedURL.Port() != configPort {
		err=errors.New(errorInvalidPort)
		return
	}
	u = urlShorten(*parsedURL)
	return
}

func (u urlShorten) escapeEndpoint() (escaped string) {
	escaped=u.Path[len(shortenPath):]
	return
}

//todo: move this to config
var configSchemes = []string{"https","http"}

func (u urlShorten) checkScheme(scheme string)(err error){
	//check if it overlaps once
	for _,as:=range configSchemes{
		if as == scheme{
			return
		}
	}
	err = fmt.Errorf(errorInvalidRedirectScheme,scheme)
	return
}

// returns the string form of the url
func (u urlShorten)string()string{
	l := url.URL(u)
	return l.String()
}