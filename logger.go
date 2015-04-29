package zerver

import (
	"github.com/cosiner/gohper/log"
)

type Logger interface {
	Flush()
	Close()

	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{}) // panic goroutine
	Fatal(...interface{}) // exit process

	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})

	Debugln(...interface{})
	Infoln(...interface{})
	Warnln(...interface{})
	Errorln(...interface{})
	Fatalln(...interface{})

	DebugDepth(int, ...interface{})
	InfoDepth(int, ...interface{})
	WarnDepth(int, ...interface{})
	ErrorDepth(int, ...interface{})
	FatalDepth(int, ...interface{})
}

func DefaultLogger() Logger {
	l := log.New(&log.LoggerOption{
		Level: log.LEVEL_DEBUG,
	})
	l.AddWriter(new(log.ConsoleWriter), nil)
	return l
}