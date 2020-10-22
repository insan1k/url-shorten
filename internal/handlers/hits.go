package handlers

import (
	"errors"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
	"github.com/insan1k/one-qr-dot-me/internal/model"
	"github.com/insan1k/one-qr-dot-me/internal/shorted"
	"net/http"
	"time"
)

// hitsOverTime represents a request that gets the model.Hits over a time period
type hitsOverTime struct {
	ShortID    string `schema:"short_id"`
	TimePeriod string `schema:"period"`
}

func (h hitsOverTime) validate() error {
	_, err := shorted.NewIDFromAPI(h.ShortID)
	if err != nil {
		return err
	}
	var predefinedTimePeriods = []string{"week", "day", "all"}
	for _, p := range predefinedTimePeriods {
		if p == h.TimePeriod {
			return nil
		}
	}
	return errors.New("invalid time period")
}

func (h hitsOverTime) getTimePeriod() (start time.Time, end time.Time) {
	end = time.Now()
	if h.TimePeriod == "week" {
		start = end.Add(-time.Hour * 24 * 7)
	}
	if h.TimePeriod == "day" {
		start = end.Add(-time.Hour * 24)
	}
	if h.TimePeriod == "all" {
		start = time.Time{}
	}
	return
}

// GetHitsOverTimePeriod gets the hits over one of the supported time periods
func (e Endpoints) GetHitsOverTimePeriod(w http.ResponseWriter, r *http.Request) {
	var h hitsOverTime
	err := e.DecodeQueryParameters(&h, r.URL.Query())
	if err != nil {
		logger.L.Debugf("invalid request %v", err)
		e.Response(w, Bad)
		return
	}
	err = h.validate()
	if err != nil {
		logger.L.Debugf("invalid request %v", err)
		e.Response(w, Bad)
		return

	}
	start, end := h.getTimePeriod()
	hits, err := model.HitsFromDb(h.ShortID, start, end)
	if err != nil {
		logger.L.Debugf("failed to fetch hits from DB %v", err)
		e.Response(w, NotFound)
		return
	}
	json, err := e.EncodeJSON(hits, false)
	if err != nil {
		logger.L.Errorf("failed to encode JSON %v", err)
		e.Response(w, InternalError)
		return
	}
	e.Response(w, Success, json...)
	return
}
