package handlers

import (
	"github.com/apex/log"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"github.com/insan1k/one-qr-dot-me/internal/shorted"
	"net/http"
)

type shortUrlPost struct {
	Target string `json:"target"`
}

func (e Endpoints) PostShortURL(w http.ResponseWriter, r *http.Request) {
	var p shortUrlPost
	err := e.DecodeJson(r.Body, &p)
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	s, err := shorted.NewUrl(p.Target)
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	ms := s.ToModel()
	err = ms.PersistCache()
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	err = ms.PersistDB()
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
}

func (e Endpoints) GetRedirectShortURL(w http.ResponseWriter, r *http.Request) {
	s, err := shorted.NewShortURLFromApi(r.URL.String())
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	hit, err := shorted.HitFromAPI(s, r.RemoteAddr)
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	ms := s.ToModel()
	mm, err := model.ShortURLFromCache(ms.ID)
	if err != nil {
		mm, err = model.ShortURLFromDB(ms.ID)
		if err != nil {
			e.Response(w, InternalError, nil)
			log.Errorf("request error %v", err)
		} else {
			hit.Ended(mm.Original, false)
		}
	} else {
		hit.Ended(mm.Original, true)
	}
	http.Redirect(w, r, mm.Original, e.responses[Redirect].Code)
}
