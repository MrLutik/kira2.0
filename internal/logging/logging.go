package logging

import (
	"github.com/mrlutik/kira2.0/internal/hooks"
	"github.com/sirupsen/logrus"
)

var (
	Log            *logrus.Logger
	Hooks          = []logrus.Hook{&hooks.ColorHook{}}
	ValidLogLevels = []string{"trace", "debug", "info", "warn", "error", "fatal", "panic"}
)

func init() {
	Log = InitLogger(Hooks, "panic")
}

func InitLogger(h []logrus.Hook, logLevelStr string) *logrus.Logger {
	log := logrus.New()
	for _, hook := range h {
		log.AddHook(hook)
	}
	log.SetLevel(toLogLevel(logLevelStr))

	return log
}

func SetLevel(s string) {
	Log.SetLevel(toLogLevel(s))
}

func toLogLevel(s string) (t logrus.Level) {
	t = map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}[s]

	return
}
