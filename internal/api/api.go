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
	h      handlers.Endpoints
}

func (a *API) loadAPI() (err error) {
	a.Router = mux.NewRouter()
	a.registerRoutes()
	err = a.newHandler()
	return
}

// Load our API singleton
func Load() (err error) {
	return A.loadAPI()
}

func (a *API) registerRoutes() {
	a.h = handlers.New()
	a.Router.HandleFunc("/short", a.h.PostShortURL).Methods("POST")
	a.Router.HandleFunc("/hits", a.h.GetHitsOverTimePeriod).Methods("GET")
	a.Router.PathPrefix(configuration.URLShortenerPath).HandlerFunc(a.h.RedirectShortURL).Methods("GET")
}
