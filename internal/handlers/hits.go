package handlers

import (
	"errors"
	"github.com/apex/log"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"net/http"
	"time"
)

// HitsOverTime represents a request that gets the model.Hits over a time period
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

// GetHitsOverTimePeriod gets the hits over one of the supported time periods
func (e Endpoints) GetHitsOverTimePeriod(w http.ResponseWriter, r *http.Request) {
	var h HitsOverTime
	err := e.DecodeQueryParameters(&h, r.URL.Query())
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	err = h.validate()
	if err != nil {
		e.Response(w, NotAllowed)
		log.Errorf("request error %v", err)
		return

	}
	start, end := h.getTimePeriod()
	hits, err := model.HitsFromDb(h.shortID, start, end)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	json, err := e.EncodeJSON(hits, false)
	if err != nil {
		e.Response(w, InternalError)
		log.Errorf("request error %v", err)
		return
	}
	e.Response(w, Success, json...)
}
