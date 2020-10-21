package api

import (
	"github.com/gorilla/mux"
	"github.com/insan1k/one-qr-dot-me/internal/handlers"
)

var A API

// API holds mux router
type API struct {
	Router *mux.Router
}

func Load()(err error){
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
	e:=handlers.New()
	a.Router.HandleFunc("/short", e.PostShortURL)
	a.Router.HandleFunc("/hits", e.GetHitsOverTimePeriod)
	a.Router.PathPrefix("/").HandlerFunc(e.GetRedirectShortURL)
}
