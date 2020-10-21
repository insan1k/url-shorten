package logger

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"os"
)

var Logger log.Logger

func LoadLogger() {
	Logger.Handler=logfmt.New(os.Stderr)
}
