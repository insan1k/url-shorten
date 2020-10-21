package shorted

import (
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"time"
)

//Hit is the struct that represents the event of a redirect
type Hit struct {
	ID        ID
	ShortID   ID
	From      urlShorten
	To        urlShorten
	Address   Address
	WasCached bool
	Took      time.Duration
	Timestamp time.Time
}

//Ended is called once we finished processing the redirect from a hit
func (h *Hit) Ended(original string, cached bool) {
	h.To, _ = newOriginal(original)
	h.WasCached = cached
	h.Took = time.Now().Sub(h.Timestamp)
	return
}

//HitFromAPI creates Hit from the API endpoint
func HitFromAPI(url ShortURL, remoteAddr string) (h Hit, err error) {
	h.ID, err = NewID()
	if err != nil {
		return
	}
	h.From = url.Short
	h.ShortID = url.ID
	err = h.Address.parse(remoteAddr)
	if err != nil {
		return
	}
	h.Timestamp = time.Now()
	return
}

// HitFromModel creates Hit based on model.Hit
func HitFromModel(m model.Hit) (h Hit, err error) {
	err = h.ID.decodeID(m.ID)
	if err != nil {
		return
	}
	err = h.ShortID.decodeID(m.ShortID)
	if err != nil {
		return
	}
	from, err := parseShort(m.From)
	if err != nil {
		return
	}
	h.From = from
	to, err := newOriginal(m.To)
	if err != nil {
		return
	}
	h.To = to
	err = h.Address.parse(m.Address)
	h.WasCached = m.WasCached
	h.Took = m.Took
	h.Timestamp = m.Timestamp
	return
}

// ToModel converts Hit to model.Hit
func (h Hit) ToModel() (m model.Hit, err error) {
	m.ID = h.ID.string()
	m.ShortID = h.ShortID.string()
	m.From = h.From.string()
	m.To = h.To.string()
	m.Address = h.Address.string()
	m.WasCached = h.WasCached
	m.Took = h.Took
	m.Timestamp = h.Timestamp
	return
}
