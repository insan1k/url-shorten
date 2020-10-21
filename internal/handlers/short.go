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
		return
	}
	s, err := shorted.NewURL(p.Target)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	ms := s.ToModel()
	err = ms.PersistCache()
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	err = ms.PersistDB()
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
}

//GetRedirectShortURL gets a redirect and performs Redirect on http response
func (e Endpoints) GetRedirectShortURL(w http.ResponseWriter, r *http.Request) {
	s, err := shorted.NewPartialShortURLFromAPI(r.URL.String())
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	hit, err := shorted.HitFromAPI(s, r.RemoteAddr)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	ms := s.ToModel()
	mm,cached, err := model.FindShortURL(ms.ID)
	if err!=nil{
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	hit.Ended(mm.Original,cached)
	mh,err:=hit.ToModel()
	if err!=nil{
		// if model.Hit fails we should still redirect the client
		http.Redirect(w, r, mm.Original, e.responses[Redirect].Code)
		log.Errorf("error in hit: %v", err)
		return
	}
	err=mh.PersistDB()
	if err!=nil{
		// if model.Hit fails we should still redirect the client
		http.Redirect(w, r, mm.Original, e.responses[Redirect].Code)
		log.Errorf("error persisting hit: %v", err)
		return
	}
	http.Redirect(w, r, mm.Original, e.responses[Redirect].Code)
}
