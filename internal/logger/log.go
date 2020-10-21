package logger

import (
	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"os"
)

// L is our logger singleton
var L log.Logger

//Load logger singleton
func Load() {
	L.Handler = logfmt.New(os.Stderr)
}

//todo: configure log levels
