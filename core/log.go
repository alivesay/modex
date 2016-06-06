package core

import (
	"github.com/op/go-logging"
	"os"
	"sync"
)

type LogErrorLevel uint8

const (
	LOG_FATAL = iota
	LOG_PANIC
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
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

		logging.SetLevel(logging.DEBUG, prefix)
		logging.SetBackend(logInstance.backend)
		logging.SetFormatter(logInstance.formatter)
	})
	return logInstance
}

func Log(level LogErrorLevel, args ...interface{}) {
	log := getInstanceLog()
	switch level {
	case LOG_PANIC:
		log.logger.Panic(args)
		break
	case LOG_ERR:
		log.logger.Error(args)
		break
	case LOG_NOTICE:
		log.logger.Notice(args)
		break
	case LOG_DEBUG:
		log.logger.Debug(args)
		break
	}
}
