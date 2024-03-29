package core

import (
	"os"
	"sync"

	"github.com/op/go-logging"
)

type LogErrorLevel uint8

const (
	LogFatal LogErrorLevel = iota
	LogPanic
	LogCrit
	LogErr
	LogWarn
	LogNotice
	LogInfo
	LogDebug
)

type log struct {
	logger    *logging.Logger
	formatter logging.Formatter
	backend   logging.Backend
}

var logInstance *log
var logOnce sync.Once

func getInstanceLog() *log {
	prefix := "MODEX "
	logOnce.Do(func() {
		logInstance = &log{
			logger: logging.MustGetLogger("modex"),
			formatter: logging.MustStringFormatter(
				`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
			),
			backend: logging.NewLogBackend(os.Stderr, prefix, 0),
		}

		logInstance.logger.ExtraCalldepth = 1
		logging.SetLevel(logging.DEBUG, prefix)
		logging.SetBackend(logInstance.backend)
		logging.SetFormatter(logInstance.formatter)
	})
	return logInstance
}

func Log(level LogErrorLevel, args ...interface{}) {
	log := getInstanceLog()
	switch level {
	case LogPanic:
		log.logger.Panic(args)
	case LogErr:
		log.logger.Error(args)
	case LogNotice:
		log.logger.Notice(args)
	case LogDebug:
		log.logger.Debug(args)
	}
}
