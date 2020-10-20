package shorted

import (
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"time"
)

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

func (h *Hit) Ended(original string, cached bool) {
	h.To, _ = newOriginal(original)
	h.WasCached = cached
	h.Took = time.Now().Sub(h.Timestamp)
	return
}

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
