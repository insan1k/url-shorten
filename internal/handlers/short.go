package handlers

import (
	"github.com/insan1k/one-qr-dot-me/internal/logger"
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
		logger.L.Debugf("invalid request %v", err)
		e.Response(w, Bad)
		return
	}
	s, err := shorted.NewURL(p.Target)
	if err != nil {
		e.Response(w, Bad)
		logger.L.Debugf("request error %v", err)
		return
	}
	ms := s.ToModel()
	err = ms.PersistCache()
	if err != nil {
		// this is not a problem for the user
		logger.L.Errorf("failed persisting short url in cache %v", err)
		return
	}
	err = ms.PersistDB()
	if err != nil {
		logger.L.Errorf("failed persisting short url in database %v", err)
		e.Response(w, InternalError)
		return
	}
	json, err := e.EncodeJSON(ms, false)
	if err != nil {
		logger.L.Errorf("failed to encode JSON %v", err)
		e.Response(w, InternalError)
		return
	}
	e.Response(w, Success, json...)
	return
}

//RedirectShortURL gets a redirect and performs Redirect on http response
func (e Endpoints) RedirectShortURL(w http.ResponseWriter, r *http.Request) {
	s, err := shorted.NewPartialShortURLFromAPI(r.URL.Path)
	if err != nil {
		logger.L.Debugf("invalid short url %v", err)
		e.Response(w, Bad)
		return
	}
	var hit shorted.Hit
	hit, err = shorted.HitFromAPI(r.RemoteAddr)
	if err != nil {
		// this is not a problem for the user
		logger.L.Errorf("failed to create a hit for short url %v", err)
		return
	}
	ms := s.ToModel()
	var cached bool
	var mm model.ShortURL
	mm, cached, err = model.FindShortURL(ms.ShortID)
	if err != nil {
		logger.L.Errorf("failed to find short url %v", err)
		e.Response(w, NotFound)
		return
	}
	hit.Ended(mm,cached)
	var mh model.Hit
	mh, err = hit.ToModel()
	if err != nil {
		// if model.Hit fails we should still redirect the client
		logger.L.Errorf("error in hit: %v", err)
		e.Redirect(w, mm.Original)
		return
	}
	err = mh.Persist()
	if err != nil {
		// if model.Hit.Persist() fails we should still redirect the client
		logger.L.Errorf("error persisting hit: %v", err)
		e.Redirect(w, mm.Original)
		return
	}
	e.Redirect(w, mm.Original)
	return
}
