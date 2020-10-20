package api

import (
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/rs/cors"
	"net"
	"net/http"
)

func (a *API) newHandler(c configuration.Configuration) (err error) {
	l, err := net.Listen(c.HttpProtocol, c.HttpHostname)
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
		ReadTimeout:  c.HttpReadTimeout,
		WriteTimeout: c.HttpWriteTimeout,
		IdleTimeout:  c.HttpIdleTimeout,
		Handler:      cr.Handler(a.Router),
	}

	if c.ServerKeyPath == "" {
		err = srv.Serve(l)
	} else {
		err = srv.ServeTLS(l, c.ServerCertPath, c.ServerKeyPath)
	}
	return
}
