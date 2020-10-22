package api

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/rs/cors"
	"net"
	"net/http"
)

func (a *API) newHandler() (err error) {
	l, err := net.Listen(configuration.C.HTTPBindProtocol, configuration.C.HTTPHostname)
	if err != nil {
		return
	}

	cr := cors.New(
		cors.Options{
			AllowedOrigins:     []string{"*"},
			AllowedMethods:     []string{"GET", "HEAD", "POST", "PUT", "OPTIONS"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
			AllowedHeaders:     []string{"Origin, Accept, Authorization, X-Requested-With, X-Forwarded-For, Content-Type"},
			Debug:              true,
		},
	)

	srv := &http.Server{
		ReadTimeout:  configuration.C.HTTPReadTimeout,
		WriteTimeout: configuration.C.HTTPWriteTimeout,
		IdleTimeout:  configuration.C.HTTPIdleTimeout,
		Handler:      cr.Handler(a.Router),
	}

	if configuration.C.HTTPTLSKeyPath == "" {
		err = srv.Serve(l)
	} else {
		err = srv.ServeTLS(l, configuration.C.HTTPTLSCertPath, configuration.C.HTTPTLSKeyPath)
	}
	return
}
