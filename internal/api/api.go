package api

import (
	"github.com/gorilla/mux"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
)

// API holds mux router
type API struct {
	Router *mux.Router
}

func (a *API) LoadAPI(c configuration.Configuration) (err error) {
	a.Router = mux.NewRouter()
	err = a.newHandler(c)
	if err != nil {
		return err
	}
	a.registerRoutes()
	return
}

func (a *API) registerRoutes() {
	//todo: register routes here
}
