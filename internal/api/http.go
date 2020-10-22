package api

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"net"
	"net/http"
)

func (a *API) newHandler() (err error) {
	l, err := net.Listen(configuration.C.HTTPBindProtocol, configuration.C.HTTPHostname)
	if err != nil {
		return
	}

	srv := &http.Server{
		ReadTimeout:  configuration.C.HTTPReadTimeout,
		WriteTimeout: configuration.C.HTTPWriteTimeout,
		IdleTimeout:  configuration.C.HTTPIdleTimeout,
		Handler:      a.Router,
	}

	if configuration.C.HTTPTLSKeyPath == "" {
		err = srv.Serve(l)
	} else {
		err = srv.ServeTLS(l, configuration.C.HTTPTLSCertPath, configuration.C.HTTPTLSKeyPath)
	}
	return
}
