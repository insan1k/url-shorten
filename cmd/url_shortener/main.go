package main

import(
	"github.com/insan1k/one-qr-dot-me/internal/configuration"
	"github.com/insan1k/one-qr-dot-me/internal/database"
	"github.com/insan1k/one-qr-dot-me/internal/cache"
	"github.com/insan1k/one-qr-dot-me/internal/logger"
	"github.com/insan1k/one-qr-dot-me/internal/api"
	)


func main() {
	configuration.C.Load()
	logger.LoadLogger()
	err:=database.LoadDB()
	if err!=nil{
		logger.Logger.Errorf("fatal %v",err)
		return
	}
	defer func(){
		err = database.StopDB()
		logger.Logger.Errorf("error stopping db %v",err)
	}()
	err = cache.LoadCache()
	if err!=nil{
		logger.Logger.Errorf("fatal %v",err)
		return
	}
	err=api.Load()
	if err!=nil {
		logger.Logger.Errorf("fatal %v",err)
		return
	}
}
