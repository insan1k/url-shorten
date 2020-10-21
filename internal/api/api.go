package api

import (
	"github.com/gorilla/mux"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/insan1k/one-qr-dot-me/internal/handlers"
)

// A is our API singleton
var A API

// API holds mux router
type API struct {
	Router *mux.Router
}

// Load our API singleton
func Load() (err error) {
	return A.loadAPI()
}

func (a *API) loadAPI() (err error) {
	a.Router = mux.NewRouter()
	err = a.newHandler()
	if err != nil {
		return err
	}
	a.registerRoutes()
	return
}

func (a *API) registerRoutes() {
	e := handlers.New()
	a.Router.HandleFunc("/short", e.PostShortURL)
	a.Router.HandleFunc("/hits", e.GetHitsOverTimePeriod)
	a.Router.PathPrefix(configuration.URLShortenerPath).HandlerFunc(e.GetRedirectShortURL)
}
