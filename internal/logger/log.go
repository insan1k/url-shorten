package logger

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"os"
)

var Logger log.Logger

func LoadLogger() {
	log.SetHandler(logfmt.New(os.Stderr))
}
