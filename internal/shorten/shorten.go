package shorten

import (
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/shamaton/msgpack"
)

//todo: change  const to configuration
const (
	shortenPath    = "/r/"
	configHostname = "localhost"
	configScheme   = "http"
	configPort     = ""
)

// ShortURL struct holds all the data required for redirecting a request
type ShortURL struct {
	ID       ID
	Original urlShorten
	Short    urlShorten
}

// NewShortenedUrl retrieves ShortURL from a shortened urlShorten
func NewShortenedUrl(u string) (s ShortURL, err error) {
	if s.Short, err = newShort(u); err != nil {
		return
	}
	err = s.ID.decodeID(s.Short.escapeEndpoint())
	if err != nil {
		return
	}
	//todo: look for ShortURL in cache
	//todo: if not found look for ShortURL in database
	return
}

// NewShortenedUrl creates ShortURL from a original urlShorten
func NewOriginalUrl(u string) (s ShortURL, err error) {
	if s.Original, err = newOriginal(u); err != nil {
		return
	}
	s.ID, err = NewID()
	if err != nil {
		return
	}
	//todo: go func cache this immediately after it's created (it's more likely to be used at creation time)
	err = s.setInCache()
	if err != nil {
		return
	}
	//todo: go func persist to database
	return
}

func (s ShortURL) setInCache()(err error){
	packed,err := s.toMsgPack()
	if err != nil {
		return
	}
	err=cache.C.Set(s.ID.encodeID(),packed)
	return
}

func (s *ShortURL)getFromCache()(err error){
	packed,err:=cache.C.Get(s.ID.encodeID())
	if err != nil {
		return
	}
	err = s.fromMsgPack(packed)
	return
}

type ShortURLPacked struct {
	ID       string `msgpack:"id"`
	Original string `msgpack:"original"`
	Short    string `msgpack:"shortened"`
}

func (s ShortURL) toMsgPack() (packed []byte, err error) {
	toPack := ShortURLPacked{
		ID:       s.ID.encodeID(),
		Original: s.Original.string(),
		Short:    s.Short.string(),
	}
	packed, err = msgpack.Encode(toPack)
	return
}

func (s *ShortURL) fromMsgPack(packed []byte) (err error) {
	unpacked := ShortURLPacked{}
	err = msgpack.Decode(packed, &unpacked)
	if err != nil {
		return
	}
	err=s.ID.decodeID(unpacked.ID)
	if err != nil{
		return
	}
	original, err := newOriginal(unpacked.Original)
	if err != nil {
		return
	}
	s.Original = original
	short, err := newShort(unpacked.Short)
	if err != nil {
		return
	}
	s.Short = short
	return
}
