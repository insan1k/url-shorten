package handlers

import (
	"errors"
	"github.com/apex/log"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"net/http"
	"time"
)

type HitsOverTime struct {
	shortID    string `schema:"id"`
	timePeriod string `schema:"period"`
}

func (h HitsOverTime) validate() error {
	var predefinedTimePeriods = []string{"week", "day", "all"}
	for _, p := range predefinedTimePeriods {
		if p == h.timePeriod {
			return nil
		}
	}
	return errors.New("invalid time period")
}

func (h HitsOverTime) getTimePeriod() (start time.Time, end time.Time) {
	end = time.Now()
	if h.timePeriod == "week" {
		start = end.Add(-time.Hour * 24 * 7)
	}
	if h.timePeriod == "day" {
		start = end.Add(-time.Hour * 24)
	}
	if h.timePeriod == "all" {
		start = time.Time{}
	}
	return
}

func (e Endpoints) GetHitsOverTimePeriod(w http.ResponseWriter, r *http.Request) {
	var h HitsOverTime
	err := e.DecodeQueryParameters(&h, r.URL.Query())
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	err = h.validate()
	if err != nil {
		e.Response(w, NotAllowed, nil)
		log.Errorf("request error %v", err)
	}
	start, end := h.getTimePeriod()
	hits, err := model.HitsFromDb(h.shortID, start, end)
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	json, err := e.EncodeJson(hits, false)
	if err != nil {
		e.Response(w, InternalError, nil)
		log.Errorf("request error %v", err)
	}
	e.Response(w, Success, json)
	return
}
