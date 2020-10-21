package api

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/rs/cors"
	"net"
	"net/http"
)

func (a *API) newHandler() (err error) {
	l, err := net.Listen(configuration.C.HttpProtocol, configuration.C.HttpHostname)
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
		ReadTimeout:  configuration.C.HttpReadTimeout,
		WriteTimeout: configuration.C.HttpWriteTimeout,
		IdleTimeout:  configuration.C.HttpIdleTimeout,
		Handler:      cr.Handler(a.Router),
	}

	if configuration.C.ServerKeyPath == "" {
		err = srv.Serve(l)
	} else {
		err = srv.ServeTLS(l, configuration.C.ServerCertPath, configuration.C.ServerKeyPath)
	}
	return
}
