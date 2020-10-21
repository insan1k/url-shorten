package handlers

import (
	"github.com/apex/log"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"github.com/insan1k/one-qr-dot-me/internal/shorted"
	"net/http"
)

// shortURLPost represents the post request that PostShortURL receives
type shortURLPost struct {
	Target string `json:"target"`
}

// PostShortURL is the endpoint that creates a shorted.ShortURL
func (e Endpoints) PostShortURL(w http.ResponseWriter, r *http.Request) {
	var p shortURLPost
	err := e.DecodeJSON(r.Body, &p)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
	s, err := shorted.NewURL(p.Target)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
	ms := s.ToModel()
	err = ms.PersistCache()
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
	err = ms.PersistDB()
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
}

//GetRedirectShortURL gets a redirect and performs Redirect on http response
func (e Endpoints) GetRedirectShortURL(w http.ResponseWriter, r *http.Request) {
	s, err := shorted.NewShortURLFromAPI(r.URL.String())
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
	hit, err := shorted.HitFromAPI(s, r.RemoteAddr)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
	}
	ms := s.ToModel()

	//todo isolate this into a function that persists stuff into cache after reading from database
	mm, err := model.ShortURLFromCache(ms.ID)
	if err != nil {
		mm, err = model.ShortURLFromDB(ms.ID)
		if err != nil {
			e.Response(w, InternalError)
			log.Errorf("request error %v", err)
		} else {
			hit.Ended(mm.Original, false)
		}
	} else {
		hit.Ended(mm.Original, true)
	}
	http.Redirect(w, r, mm.Original, e.responses[Redirect].Code)
}
