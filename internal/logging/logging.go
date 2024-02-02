package logging

import (
	"errors"
	"fmt"
	"strings"

	"github.com/mrlutik/kira2.0/internal/hooks"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var ErrInvalidLoggingLevel = errors.New("invalid logging level")

func InitLogger(hooks []logrus.Hook, logLevelStr string) (*Logger, error) {
	log := &Logger{logrus.New()}
	for _, hook := range hooks {
		log.AddHook(hook)
	}

	err := log.SetLoggingLevel(logLevelStr)
	if err != nil {
		return nil, fmt.Errorf("setting logging level error: %w", err)
	}

	return log, nil
}

func (l *Logger) SetLoggingLevel(s string) error {
	level, err := toLogLevel(s)
	if err != nil {
		return fmt.Errorf("can't set the logging level '%s', error: %w", s, err)
	}

	l.SetLevel(level)

	return nil
}

func toLogLevel(s string) (logrus.Level, error) {
	levelsMap := map[string]logrus.Level{
		"trace": logrus.TraceLevel,
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	level, ok := levelsMap[strings.ToLower(s)]
	if !ok {
		return 0, ErrInvalidLoggingLevel
	}

	return level, nil
}

func GetValidLogLevels() []string {
	return []string{"trace", "debug", "info", "warn", "error", "fatal", "panic"}
}

func GetHooks() []logrus.Hook {
	return []logrus.Hook{&hooks.ColorHook{}}
}

func IsValidLogLevel(levelValue string) bool {
	for _, validLevel := range GetValidLogLevels() {
		if levelValue == validLevel {
			return true
		}
	}
	return false
}
