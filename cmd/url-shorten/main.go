package main

import (
	"github.com/insan1k/one-qr-dot-me/internal/api"
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
)

func main() {
	configuration.C.Load()
	logger.Load()
	err := database.Load()
	if err != nil {
		logger.L.Errorf("fatal %v", err)
		return
	}
	defer func() {
		err = database.Stop()
		logger.L.Errorf("error stopping db %v", err)
	}()
	err = cache.LoadCache()
	if err != nil {
		logger.L.Errorf("fatal %v", err)
		return
	}
	err = api.Load()
	if err != nil {
		logger.L.Errorf("fatal %v", err)
		return
	}
}