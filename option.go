package glog

import (
	"fmt"
)

type Opt func(*loggingT)

func LogDir(dir string) Opt {
	return func(lg *loggingT) {
		logDir = dir
	}
}

func LogToStderr(flag bool) Opt {
	return func(lg *loggingT) {
		lg.toStderr = flag
	}
}

func AlsoLogToStderr(flag bool) Opt {
	return func(lg *loggingT) {
		lg.alsoToStderr = flag
	}
}

func Verbosity(v int) Opt {
	return func(lg *loggingT) {
		lg.verbosity.Set(fmt.Sprint(v))
	}
}

func StderrThreshold(value string) Opt {
	return func(lg *loggingT) {
		lg.stderrThreshold.Set(value)
	}
}

func VModule(value string) Opt {
	return func(lg *loggingT) {
		lg.vmodule.Set(value)
	}
}

func TraceLocation(value string) Opt {
	return func(lg *loggingT) {
		lg.traceLocation.Set(value)
	}
}

func Open(opts ...Opt) {
	for _, opt := range opts {
		opt(&logging)
	}
}

func Close() {
	logging.lockAndFlushAll()
}
