package main

import (
	"github.com/insan1k/one-qr-dot-me/internal/api"
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
)

func main() {
	logger.Load()
	configuration.C.Load()
	err := database.Load()
	if err != nil {
		logger.L.Errorf("could not connect to the database %v", err)
		return
	}
	defer func() {
		if deferErr := database.Stop(); deferErr != nil {
			logger.L.Errorf("session close error %v", deferErr)
		}
	}()
	err = cache.Load()
	if err != nil {
		logger.L.Errorf("could not load cache %v", err)
		return
	}
	err = api.Load()
	if err != nil {
		logger.L.Errorf("could not start API %v", err)
		return
	}
}
